import React, { useState, useEffect } from "react";
import styled from "styled-components";
import { PageWrapper } from "../../components/PageWrapper";
import { Button } from "../../components/Button";

export const LobbyPage = ({
  match: {
    params: { roomId },
  },
  socket,
  history,
}) => {
  const [loading, setLoading] = useState(true);
  const [players, setPlayers] = useState([]);

  useEffect(() => {
    socket.send(
      JSON.stringify({
        type: "request_room_state",
      })
    );
  }, []);

  useEffect(() => {
    socket.onmessage = function (event) {
      const message = JSON.parse(event.data);
      const { type, data } = message;

      switch (type) {
        case "room_state":
          setPlayers(data.players);
          setLoading(false);
          break;
        case "game_started":
          history.push(`/room/${roomId}/game`);
          break;
        default:
          console.warn("received unknown WS type");
      }
    };
  });

  const handleStartGameClick = () => {
    socket.send(
      JSON.stringify({
        type: "start_game",
      })
    );
  };

  return (
    <PageWrapper
      customStyles={{
        padding: "16px",
        alignItems: "flex-start",
        minWidth: "300px",
      }}
    >
      {loading ? (
        "Loading"
      ) : (
        <>
          <StyledParagraph>Room Code: {roomId.toUpperCase()}</StyledParagraph>
          {players.length > 1 ? (
            <Button
              onClick={handleStartGameClick}
              text="Start Game"
              style={{ alignSelf: "center" }}
            />
          ) : (
            <Button
              onClick={handleStartGameClick}
              text="Waiting for Players"
              disabled
              style={{ alignSelf: "center" }}
            />
          )}
          <StyledParagraph>Players</StyledParagraph>
          <StyledList>
            {players.map((name) => (
              <li>"{name}"</li>
            ))}
          </StyledList>
        </>
      )}
    </PageWrapper>
  );
};

const StyledParagraph = styled.p`
  text-align: left;
  margin-bottom: 0;
  margin-top: 0;
`;

const StyledList = styled.ul`
  text-align: left;
  list-style-type: none;
  padding: 0;
  margin: 0;
`;
