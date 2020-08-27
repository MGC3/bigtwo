import React, { Component } from "react";
import "../css/card.css";
import { GiDiamonds, GiHearts, GiSpades, GiClubs } from "react-icons/gi";

export class Card extends Component {
  constructor(props) {
    super(props);

    this.state = {
      data: this.props.data,
    };
  }

  render() {
    const iconMaker = (suit) => {
      switch (suit) {
        case "D":
          return <GiDiamonds color={"red"} />;
        case "H":
          return <GiHearts color={"red"} />;
        case "S":
          return <GiSpades color={"black"} />;
        case "C":
          return <GiClubs color={"black"} />;
        default:
          return "something bad happened";
      }
    };

    const icon = iconMaker(this.state.data.suit);

    return (
      <div
        className={"card-container " + (this.props.selected && "card-selected")}
        onClick={() => this.props.handleCardSelect(this.state.data)}
      >
        <div className="upper">
          <div className="card-rank">{this.state.data.rank}</div>
          <div className="card-suit">{icon}</div>
        </div>
        <div className="lower">
          <div className="card-rank">{this.state.data.rank}</div>
          <div className="card-suit">{icon}</div>
        </div>
      </div>
    );
  }
}
