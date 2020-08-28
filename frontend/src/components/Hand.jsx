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
    // deselect the card if it's already been selected
    if (this.state.selectedCards.includes(card)) {
      let updatedCards = [...this.state.selectedCards].filter(
        (d) => d !== card
      );
      this.setState({
        selectedCards: updatedCards,
      });
    } else {
      // don't allow the user to select more than 5 cards at a time
      if (this.state.selectedCards.length > 4) {
        return;
      }
      // add the card to the selectedCards array
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
        {/* TODO: debugging code, remove later vvvvvvvvv*/}
        <h1 style={{ marginBottom: "64px" }}>
          SelectedCards state is:
          {this.state.selectedCards.length > 0 &&
            this.state.selectedCards.map((card) => (
              <span>
                {card.rank} of {card.suit},{" "}
              </span>
            ))}
          {this.state.selectedCards.length === 0 && "None Selected"}
        </h1>
        {/* TODO: debugging code, remove later ^^^^^^^^^*/}
        <div className="hand-container">
          {this.state.cards.map((card, i) => {
            return (
              <Card
                key={card.rank + card.suit}
                data={card}
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
