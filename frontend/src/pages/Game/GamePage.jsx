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
  last_played_hand: [],
  current_user_turn: "Michael",
  user_player_number: 0,
};

export const GamePage = ({ socket }) => {
  const [loading, setLoading] = useState(true);
  const [userHand, setUserHand] = useState([]);
  const [allPlayerHands, setAllPlayerHands] = useState([]);
  const [lastPlayedHand, setLastPlayedHand] = useState([]);
  const [currentUserTurn, setCurrentUserTurn] = useState("");
  const [userPlayerNumber, setUserPlayerNumber] = useState(0);

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
          setUserHand(data.user_hand);
          setAllPlayerHands(data.all_player_hands);
          setLastPlayedHand(data.last_played_hand);
          setCurrentUserTurn(data.current_user_turn);
          setUserPlayerNumber(data.user_player_number);
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
        <div>
          <div>"Player 3"</div>
          <Hand cards={mockCardsData} />
        </div>
        <CenterRow>
          <div>
            <div> "Player 2"</div>
            <Hand cards={mockCardsData} rotate />
          </div>

          <DroppableArea />
          <div>
            <div> "Player 4"</div>
            <Hand cards={mockCardsData} rotate />
          </div>
        </CenterRow>
        <ButtonGroup>
          <Button text="Pass" />
          <Button text="Play" />
        </ButtonGroup>
        <div>
          <Hand cards={mockCardsData} isPlayer />
          <div> "Michael"</div>
        </div>
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
  justify-content: center;
  align-items: center;
`;
