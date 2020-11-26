import React, { useState, useEffect } from "react";
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
  const [playerId, setPlayerId] = useState(null);
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
          setPlayerId(data.client_id);
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
    <PageWrapper>
      {loading ? (
        "Loading"
      ) : (
        <div>
          <h1>Room id is: {roomId}</h1>
          <Button onClick={handleStartGameClick} text="Start Game" />
          <h1>Player id is: {playerId}</h1>
          <h1>You are: {players[playerId]}</h1>
          <h1>Players list:</h1>
          <ul>
            {players.map((name) => (
              <li>{name}</li>
            ))}
          </ul>
        </div>
      )}
    </PageWrapper>
  );
};
