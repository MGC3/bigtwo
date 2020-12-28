import React, { useState, useEffect } from "react";
import styled from "styled-components";
import { PageWrapper } from "../../components/PageWrapper";
import { Button } from "../../components/Button";
import { Hand } from "../../components/Hand";
import { mockCardsData } from "../../api/testData";

const mockGameState = {
  user_hand: mockCardsData,
  all_player_hands: [
    { name: "Michael", count: 13 },
    { name: "Player 2", count: 6 },
    { name: "Player 3", count: 8 },
    { name: "Player 4", count: 4 },
  ],
  last_played_hand: [
    {
      rank: "3",
      suit: "C",
    },
  ],
  current_user_turn: "Michael",
  client_id: 0,
};

export const GamePage = ({ socket }) => {
  const [loading, setLoading] = useState(true);
  const [userHand, setUserHand] = useState([]);
  // const [allPlayerHands, setAllPlayerHands] = useState([]); // might not need this
  const [lastPlayedHand, setLastPlayedHand] = useState([]);
  const [currentUserTurn, setCurrentUserTurn] = useState("");
  const [userPlayerNumber, setUserPlayerNumber] = useState(0);

  let isTwoPlayerGame = false;
  let isThreePlayerGame = false;
  let player1 = null;
  let player2 = null;
  let player3 = null;
  let player4 = null;

  // comment out until backened WS is ready or mocked
  // useEffect(() => {
  //   socket.send(
  //     JSON.stringify({
  //       type: "request_game_state",
  //     })
  //   );
  // }, []);

  useEffect(() => {
    socket.onmessage = function (event) {
      const message = JSON.parse(event.data);
      const { type, data } = message;

      switch (type) {
        case "game_state":
          if (data.all_player_hands.length === 2) {
            isTwoPlayerGame = true;
          } else if (data.all_player_hands.length === 3) {
            isThreePlayerGame = true;
          }

          player1 = data.all_player_hands[data.client_id];

          let otherPlayers = data.all_player_hands.splice(data.client_id, 1);

          player2 = otherPlayers[0];

          if (otherPlayers[1]) {
            player3 = otherPlayers[1];
          }

          if (otherPlayers[2]) {
            player4 = otherPlayers[2];
          }

          setUserHand(data.user_hand);
          // setAllPlayerHands(data.all_player_hands); // might not need this
          setLastPlayedHand(data.last_played_hand);
          setCurrentUserTurn(data.current_user_turn);
          setUserPlayerNumber(data.client_id);
          setLoading(false);

          break;
        default:
          console.warn("received unknown WS type");
      }
    };
  });

  return (
    <PageWrapper>
      <GameContainer>
        <PlayerSlot>
          {/* <div>{player2.name}</div>
          <Hand count={player2.count} /> */}
          <div>"Player 2"</div>
          <Hand count={13} />
        </PlayerSlot>
        <CenterRow>
          <LeftPlayerSlot rotate>
            {/* <div>{player3.name}</div>
            <Hand count={player3.count} /> */}
            <div> "Player 3"</div>
            <Hand count={6} rotate />
          </LeftPlayerSlot>
          <DroppableArea />
          <RightPlayerSlot rotate>
            {/* <div>{player4.name}</div>
          <Hand count={player4.count} /> */}
            <div> "Player 4"</div>
            <Hand count={6} />
          </RightPlayerSlot>
        </CenterRow>
        <ButtonGroup>
          <Button text="Pass" />
          <Button text="Play" />
        </ButtonGroup>
        <PlayerSlot>
          <Hand cards={mockCardsData} isPlayer />
          {/* <Hand cards={userHand} isPlayer /> */}
          <div> "Player 1"</div>
          {/* <div>{player1.name}</div> */}
        </PlayerSlot>
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
