import React, { Component } from "react";
import { Card } from "./Card";
import "../css/hand.css";

export class Hand extends Component {
  constructor(props) {
    super(props);

    this.state = {
      cards: this.props.cards,
    };
  }

  render() {
    return (
      <div className="hand-container">
        {this.state.cards.map((card, i) => {
          return <Card key={i} data={card} zdex={i} />;
        })}
      </div>
    );
  }
}
