import React, { Component } from "react";
import { withRouter } from "react-router-dom";
import { createGame } from "../api/game";
import { Hand } from "./Hand";
import TextField from "@material-ui/core/TextField";
import "../css/app.css";
import "../css/landing.css";
import { mockCardsData } from "../api/testData";

class Landing extends Component {
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

    // TODO: remove this code after figuring out new WS approach (connecting on page
    // load, instead of on join room request)
    // this.ws = new WebSocket(`ws://localhost:8000/rooms/${roomId}`);
    // console.log("Attempting to join", roomId);
    // this.ws.onopen = () => {
    //   // on connecting, do nothing but log it to the console
    //   console.log("connected");
    // };

    // this.ws.onmessage = (evt) => {
    //   // listen to data sent from the websocket server
    //   const message = JSON.parse(evt.data);
    //   this.setState({ dataFromServer: message });
    //   console.log(message);
    // };

    // this.ws.onclose = () => {
    //   console.log("disconnected");
    //   // automatically try to reconnect on connection loss
    // };
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
            fullWidth="true"
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

export default withRouter(Landing);
