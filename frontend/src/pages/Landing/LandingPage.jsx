import React, { useState } from "react";
import { createGame } from "../../api/game";
import { Hand } from "../../components/Hand";
import TextField from "@material-ui/core/TextField";
import "../../css/app.css";
import "../../css/landing.css";
import { mockCardsData } from "../../api/testData";

export const LandingPage = ({ history, socket }) => {
  const [roomId, setRoomId] = useState("");
  const [socketTest, setSocketTest] = useState(socket);

  const handleCreateGame = (e) => {
    createGame()
      .then((res) => {
        let roomId = res.data.RoomId;
        handleJoinRoom(e, roomId);
      })
      .catch((err) => console.log("Error creating game: ", err));
  };

  const handleJoinRoom = (e, roomId) => {
    e.preventDefault();
    // TODO: verify if the room id is valid. If not, show error and don't send the user to the lobby page
    history.push("/room/" + roomId);
  };

  const handleTextFieldChange = (e) => {
    setRoomId(e.target.value);
  };

  return (
    <div className="app-container">
      <h1 className="title">Big Two</h1>
      <button
        className="primary-cta-button"
        onClick={(e) => handleCreateGame(e)}
      >
        <span className="text">Create New Game</span>
      </button>
      <div>OR</div>
      <form className="landing-form">
        <TextField
          id="outlined-basic"
          label="Enter room code"
          value={roomId}
          onChange={handleTextFieldChange}
          fullWidth={true}
        />
        <button
          className="primary-cta-button"
          onClick={(e) => handleJoinRoom(e, roomId)}
        >
          Join Existing Game
        </button>
      </form>
      <Hand cards={mockCardsData} />
    </div>
  );
};
