import React, { useState, useEffect } from "react";
import { PageWrapper } from "../../components/PageWrapper";

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
    <PageWrapper>
      <div>
        <h1>Room id is: {roomCode}</h1>
      </div>
    </PageWrapper>
  );
};
