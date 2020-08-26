import React, { Component } from "react";
import { Card } from "./Card";
import "../css/hand.css";

export class Hand extends Component {
  constructor(props) {
    super(props);

    this.state = {
      cards: this.props.cards,
      selectedCards: [],
    };
  }

  handleCardSelect = (card) => {
    if (this.state.selectedCards.includes(card)) {
      let updatedCards = [...this.state.selectedCards].filter(
        (d) => d !== card
      );
      this.setState({
        selectedCards: updatedCards,
      });
    } else {
      let updatedCards = [...this.state.selectedCards];
      updatedCards.push(card);
      this.setState({
        selectedCards: updatedCards,
      });
    }
  };

  isCardSelected = (card) => {
    return this.state.selectedCards.includes(card) ? true : false;
  };

  render() {
    return (
      <>
        <h1>
          SelectedCard state is:
          {this.state.selectedCards.length > 0 &&
            this.state.selectedCards.map((card) => (
              <span>
                {card.rank} of {card.suit}
              </span>
            ))}
          {this.state.selectedCards.length === 0 && "None Selected"}
        </h1>
        <div className="hand-container">
          {this.state.cards.map((card, i) => {
            return (
              <Card
                key={i}
                data={card}
                zdex={i}
                handleCardSelect={this.handleCardSelect}
                selected={this.isCardSelected(card)}
              />
            );
          })}
        </div>
      </>
    );
  }
}
