import React, { Component } from "react";
import "../css/card.css";

export class Card extends Component {
  constructor(props) {
    super(props);

    this.state = {
      data: this.props.data,
      zdex: this.props.zdex + 1 * 100,
    };
  }

  render() {
    const cardStyle = {
      zIndex: `${this.state.zdex}`,
      position: "relative",
    };

    return (
      <div className="card-container" style={cardStyle}>
        <div className="upper">
          <div className="card-rank">{this.state.data.rank}</div>
          <div className="card-suit">{this.state.data.suit}</div>
        </div>
        <div className="lower">
          <div className="card-rank">{this.state.data.rank}</div>
          <div className="card-suit">{this.state.data.suit}</div>
        </div>
      </div>
    );
  }
}
