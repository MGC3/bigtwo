import React, { useState, useEffect } from "react";
import styled from "styled-components";
import { PageWrapper } from "../../components/PageWrapper";
import { Button } from "../../components/Button";
import { Hand } from "../../components/Hand";
import { Dialog, DialogActions, DialogTitle } from "@material-ui/core/";

export const GamePage = ({
  match: {
    params: { roomId },
  },
  socket,
  history,
}) => {
  const [loading, setLoading] = useState(true);
  const [userHand, setUserHand] = useState([]);
  const [lastPlayedHand, setLastPlayedHand] = useState([]);
  const [currentUserTurn, setCurrentUserTurn] = useState("");
  const [selectedCards, setSelectedCards] = useState([]);
  const [player1, setPlayer1] = useState(null);
  const [player2, setPlayer2] = useState(null);
  const [player3, setPlayer3] = useState(null);
  const [player4, setPlayer4] = useState(null);
  const [gameOver, setGameOver] = useState(false);
  const [winner, setWinner] = useState("");

  useEffect(() => {
    socket.send(
      JSON.stringify({
        type: "request_game_state",
      })
    );
  }, []);

  useEffect(() => {
    socket.onmessage = function (event) {
      const message = JSON.parse(event.data);
      const { type, data } = message;

      switch (type) {
        case "game_state":
          // set users data
          setPlayer1(data.all_player_hands[data.client_id]);

          // set all other players data
          const otherPlayers = data.all_player_hands.filter(
            (player, idx) => idx !== data.client_id
          );
          setPlayer2(otherPlayers[0]);
          otherPlayers[1] ? setPlayer3(otherPlayers[1]) : setPlayer3(null);
          otherPlayers[2] ? setPlayer4(otherPlayers[2]) : setPlayer4(null);

          setUserHand(data.user_hand);
          setLastPlayedHand(data.last_played_hand);
          setCurrentUserTurn(data.current_user_turn);
          setLoading(false);

          // check win condition
          if (data.game_over) {
            setWinner(findWinner(data));
            setGameOver(true);
          } else {
            setGameOver(false);
          }
          break;
        default:
          console.warn("received unknown WS type");
      }
    };
  });

  const handlePlayButtonClick = () => {
    socket.send(
      JSON.stringify({
        type: "play_move",
        data: {
          cards: selectedCards,
        },
      })
    );
    setSelectedCards([]);
  };

  const handlePassButtonClick = () => {
    socket.send(
      JSON.stringify({
        type: "pass_move",
      })
    );
  };

  const findWinner = (data) => {
    return (
      data.all_player_hands.find((player) => player.count === 0)?.name ||
      "Error detecting winner"
    );
  };

  const handleBackToLobbyClick = () => {
    history.push(`/room/${roomId}`);
  };

  return (
    <PageWrapper>
      <GameContainer>
        <PlayerSlot>
          {player2 && (
            <>
              <div>{player2.name}</div>
              <Hand count={player2.count} />
            </>
          )}
        </PlayerSlot>
        <CenterRow>
          <LeftPlayerSlot rotate>
            {player3 && (
              <>
                <div>{player3.name}</div>
                <Hand count={player3.count} />
              </>
            )}
          </LeftPlayerSlot>
          <DroppableArea>
            {lastPlayedHand && (
              <Hand
                cards={lastPlayedHand}
                selectedCards={selectedCards}
                setSelectedCards={setSelectedCards}
              />
            )}
          </DroppableArea>
          <RightPlayerSlot rotate>
            {player4 && (
              <>
                <div>{player4.name}</div>
                <Hand count={player4.count} />
              </>
            )}
          </RightPlayerSlot>
        </CenterRow>
        {player1 && (
          <>
            <ButtonGroup>
              <Button
                text="Pass"
                onClick={handlePassButtonClick}
                disabled={player1.name !== currentUserTurn || !lastPlayedHand}
                style={{ marginRight: "24px" }}
              />
              <Button
                text="Play"
                onClick={handlePlayButtonClick}
                disabled={player1.name !== currentUserTurn}
              />
            </ButtonGroup>
            <PlayerSlot>
              <Hand
                cards={userHand}
                isPlayer
                selectedCards={selectedCards}
                setSelectedCards={setSelectedCards}
              />
              <div>{player1.name}</div>
            </PlayerSlot>
          </>
        )}
      </GameContainer>
      <Dialog
        open={gameOver}
        aria-labelledby="form-dialog-title"
        maxWidth="xs"
        fullWidth
        disableBackdropClick
        disableEscapeKeyDown
      >
        <DialogTitle id="form-dialog-title">{winner} won!</DialogTitle>
        <DialogActions>
          <Button onClick={handleBackToLobbyClick} text="Back to Room" />
        </DialogActions>
      </Dialog>
    </PageWrapper>
  );
};

const GameContainer = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  max-width: 1400px;
  width: 100%;
`;

const DroppableArea = styled.div`
  width: 600px;
  height: 300px;
`;

const ButtonGroup = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 16px;
`;

const CenterRow = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
`;

const PlayerSlot = styled.div``;

const LeftPlayerSlot = styled.div`
  width: 300px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  transform: ${(props) => (props.rotate ? "rotate(-90deg)" : "none")};
`;

const RightPlayerSlot = styled.div`
  width: 300px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  transform: ${(props) => (props.rotate ? "rotate(90deg)" : "none")};
`;
