import React, { Component } from "react";
import { createGame } from "../api/game";
import { Hand } from "./Hand";
import TextField from "@material-ui/core/TextField";
import "../css/landing.css";

let fakeData = [
  {
    rank: "K",
    suit: "H",
  },
  {
    rank: "2",
    suit: "D",
  },
  {
    rank: "3",
    suit: "C",
  },
  {
    rank: "K",
    suit: "S",
  },
  {
    rank: "2",
    suit: "H",
  },
  {
    rank: "3",
    suit: "D",
  },
  {
    rank: "K",
    suit: "C",
  },
  {
    rank: "5",
    suit: "D",
  },
  {
    rank: "6",
    suit: "C",
  },
  {
    rank: "7",
    suit: "H",
  },
  {
    rank: "8",
    suit: "D",
  },
  {
    rank: "J",
    suit: "C",
  },
  {
    rank: "J",
    suit: "S",
  },
];

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      roomId: "",
    };
  }

  componentDidMount() {
    console.log("Component Mounted");
  }

  handleCreateGame = () => {
    // create game
    createGame();
    // if success, join game
  };

  handleJoinRoom = (e, roomId) => {
    e.preventDefault();
    this.ws = new WebSocket(`ws://localhost:8000/rooms/${roomId}`);
    console.log("Attempting to join", roomId);
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
  };

  handleTextFieldChange = (e) => {
    this.setState({
      roomId: e.target.value,
    });
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
          <button
            className="primary-cta-button"
            onClick={(e) => this.handleJoinRoom(e, this.state.roomId)}
          >
            <span className="text">Join Game</span>
          </button>
          <TextField
            id="outlined-basic"
            label="Enter room code"
            variant="outlined"
            value={this.state.roomId}
            onChange={this.handleTextFieldChange}
          />
        </form>
        <Hand cards={fakeData} />
      </div>
    );
  }
}

export default App;
