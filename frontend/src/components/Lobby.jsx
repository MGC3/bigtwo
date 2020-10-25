import React, { Component } from "react";

class Lobby extends Component {
  constructor(props) {
    super(props);
    this.state = {
      roomId: this.props.match.params.roomId,
    };
  }

  componentDidMount = () => {};
  render() {
    return (
      <div>
        <h1>Room id is: {this.state.roomId}</h1>
      </div>
    );
  }
}

export default Lobby;
