import React, { useState, useEffect } from "react";

export const LobbyPage = ({
  match: {
    params: { roomId },
  },
}) => {
  const [roomCode, setRoomCode] = useState("");

  useEffect(() => {
    roomId ? setRoomCode(roomId) : setRoomCode("error");
  }, []);

  return (
    <div className="app-container">
      <div>
        <h1>Room id is: {roomCode}</h1>
      </div>
    </div>
  );
};
