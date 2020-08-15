import React, { Component } from "react";
import { createGame } from "../api/game";
import TextField from "@material-ui/core/TextField";
import "../css/landing.css";

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      test: null,
    };
  }

  ws = new WebSocket("ws://localhost:8000/");

  componentDidMount() {
    this.ws.onopen = () => {
      // on connecting, do nothing but log it to the console
      console.log("connected");
    };

    this.ws.onmessage = (evt) => {
      // listen to data sent from the websocket server
      const message = JSON.parse(evt.data);
      this.setState({ dataFromServer: message });
      console.log(message);
    };

    this.ws.onclose = () => {
      console.log("disconnected");
      // automatically try to reconnect on connection loss
    };
  }

  handleCreateGame = () => {
    // create game
    createGame();
    // if success, join game
  };

  handleJoinRoom = () => {
    alert("Join Room Clicked");
  };

  render() {
    return (
      <div className="app-container">
        <h1 className="title">Big Two</h1>
        <button className="primary-cta-button" onClick={this.handleCreateGame}>
          <span className="text">New Game</span>
        </button>
        <div>OR</div>
        <form className="landing-form">
          <button className="primary-cta-button" onClick={this.handleJoinRoom}>
            <span className="text">Join Game</span>
          </button>
          <TextField
            id="outlined-basic"
            label="Enter room code"
            variant="outlined"
          />
        </form>
      </div>
    );
  }
}

export default App;
