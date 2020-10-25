import React, { Component } from "react";
import { createGame } from "../api/game";
import { Hand } from "./Hand";
import TextField from "@material-ui/core/TextField";
import "../css/app.css";
import "../css/landing.css";
import { mockCardsData } from "../api/testData";

export class Landing extends Component {
  constructor(props) {
    super(props);

    this.state = {
      roomId: "",
    };
  }

  handleCreateGame = (e) => {
    // create game
    createGame()
      .then((res) => {
        let roomId = res.data.RoomId;
        this.handleJoinRoom(e, roomId);
      })
      .catch((err) => console.log("Error creating game: ", err));
  };

  handleJoinRoom = (e, roomId) => {
    e.preventDefault();
    // TODO: verify if the room id is valid. If not, show error and don't send the user to the lobby page
    this.props.history.push("/room/" + roomId);
  };

  handleTextFieldChange = (e) => {
    this.setState({
      roomId: e.target.value,
    });
  };

  render() {
    return (
      <>
        <h1 className="title">Big Two</h1>
        <button
          className="primary-cta-button"
          onClick={(e) => this.handleCreateGame(e)}
        >
          <span className="text">Create New Game</span>
        </button>
        <div>OR</div>
        <form className="landing-form">
          <TextField
            id="outlined-basic"
            label="Enter room code"
            value={this.state.roomId}
            onChange={this.handleTextFieldChange}
            fullWidth={true}
          />
          <button
            className="primary-cta-button"
            onClick={(e) => this.handleJoinRoom(e, this.state.roomId)}
          >
            Join Existing Game
          </button>
        </form>
        <Hand cards={mockCardsData} />
      </>
    );
  }
}
