import React, { useState } from "react";
import styled from "styled-components";
import { Card } from "../Card/Card";

export const Hand = ({ cards }) => {
  const [selectedCards, setSelectedCards] = useState([]);

  const handleCardSelect = (card) => {
    // deselect the card if it's already been selected
    if (selectedCards.includes(card)) {
      let updatedCards = [...selectedCards].filter((d) => d !== card);
      setSelectedCards(updatedCards);
    } else {
      // don't allow the user to select more than 5 cards at a time
      if (selectedCards.length > 4) {
        return;
      }
      // add the card to the selectedCards array
      let updatedCards = [...selectedCards];
      updatedCards.push(card);
      setSelectedCards(updatedCards);
    }
  };

  const isCardSelected = (card) => {
    return selectedCards.includes(card) ? true : false;
  };

  return (
    <>
      {/* TODO: debugging code, remove later vvvvvvvvv*/}
      <h1 style={{ marginBottom: "64px" }}>
        SelectedCards state is:
        {selectedCards.length > 0 &&
          selectedCards.map((card) => (
            <span>
              {card.rank} of {card.suit},{" "}
            </span>
          ))}
        {selectedCards.length === 0 && "None Selected"}
      </h1>
      {/* TODO: debugging code, remove later ^^^^^^^^^*/}
      <HandContainer>
        {cards.map((card, i) => {
          return (
            <Card
              key={card.rank + card.suit}
              data={card}
              handleCardSelect={handleCardSelect}
              selected={isCardSelected(card)}
            />
          );
        })}
      </HandContainer>
    </>
  );
};

const HandContainer = styled.div`
  display: flex;
`;
