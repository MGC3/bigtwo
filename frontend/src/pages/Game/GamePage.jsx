import React, { useState, useEffect } from "react";
import styled from "styled-components";
import { PageWrapper } from "../../components/PageWrapper";
import { Button } from "../../components/Button";
import { Hand } from "../../components/Hand";

export const GamePage = ({ socket }) => {
  const [loading, setLoading] = useState(true);
  const [userHand, setUserHand] = useState([]);
  const [lastPlayedHand, setLastPlayedHand] = useState([]);
  const [currentUserTurn, setCurrentUserTurn] = useState("");
  const [selectedCards, setSelectedCards] = useState([]);
  const [player1, setPlayer1] = useState(null);
  const [player2, setPlayer2] = useState(null);
  const [player3, setPlayer3] = useState(null);
  const [player4, setPlayer4] = useState(null);

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
          setPlayer1(data.all_player_hands[data.client_id]);

          data.all_player_hands.splice(data.client_id, 1);
          let otherPlayers = data.all_player_hands;

          setPlayer2(otherPlayers[0]);

          if (otherPlayers[1]) {
            setPlayer3(otherPlayers[1]);
          }

          if (otherPlayers[2]) {
            setPlayer4(otherPlayers[2]);
          }

          setUserHand(data.user_hand);
          setLastPlayedHand(data.last_played_hand);
          setCurrentUserTurn(data.current_user_turn);
          setLoading(false);

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
  };

  const handlePassButtonClick = () => {
    socket.send(
      JSON.stringify({
        type: "pass_move",
      })
    );
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
                disabled={player1.name !== currentUserTurn || !lastPlayedHand} // uncomment when game_state WS response implemented
              />
              <Button
                text="Play"
                onClick={handlePlayButtonClick}
                disabled={player1.name !== currentUserTurn} // uncomment when game_state WS response implemented
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
  border: 1px solid red;
`;

const DroppableArea = styled.div`
  width: 600px;
  height: 300px;
  border: 1px dashed pink;
`;

const ButtonGroup = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
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
