import React, { useState } from "react";
import { createGame } from "../../api/game";
import { PageWrapper } from "../../components/PageWrapper";
import { Hand } from "../../components/Hand";
import { Button } from "../../components/Button/Button";
import TextField from "@material-ui/core/TextField";
import { mockCardsData } from "../../api/testData";

export const LandingPage = ({ history, socket }) => {
  const [roomId, setRoomId] = useState("");
  const [socketTest, setSocketTest] = useState(socket);

  const handleCreateGame = (e) => {
    console.log("Creating game");
    socketTest.send(
      JSON.stringify({
        type: "create_room",
      })
    );
    /* 
    createGame()
      .then((res) => {
        let roomId = res.data.RoomId;
        handleJoinRoom(e, roomId);
      })
      .catch((err) => console.log("Error creating game: ", err));
    */
  };

  const handleJoinRoom = (e, roomId) => {
    e.preventDefault();
    // TODO: verify if the room id is valid. If not, show error and don't send the user to the lobby page
    //history.push("/room/" + roomId);
    socketTest.send(
      JSON.stringify({
        type: "join_room",
        data: {
          room: "ABCD",
          name: "testplayer",
        },
      })
    );
  };

  const handleTextFieldChange = (e) => {
    setRoomId(e.target.value);
  };

  return (
    <PageWrapper>
      <Button onClick={(e) => handleCreateGame(e)} text="Create New Game" />
      <div>OR </div>
      <form>
        <TextField
          id="outlined-basic"
          label="Enter room code"
          value={roomId}
          onChange={handleTextFieldChange}
          fullWidth={true}
        />
        <Button
          onClick={(e) => handleJoinRoom(e, roomId)}
          text="Join Existing Game"
        />
      </form>
      <Hand cards={mockCardsData} />
    </PageWrapper>
  );
};
