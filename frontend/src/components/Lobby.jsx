import React, { Component } from "react";

class Lobby extends Component {
  render() {
    const { roomId } = this.props.match.params;

    return (
      <div>
        <h1>Room id is: {roomId}</h1>
      </div>
    );
  }
}

export default Lobby;
