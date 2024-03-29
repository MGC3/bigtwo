import React from "react";
import styled from "styled-components";
import { Card } from "../Card/";

export const Hand = ({
  cards,
  isPlayer,
  count,
  selectedCards,
  setSelectedCards,
}) => {
  const handleCardSelect = (card) => {
    // return if user is trying to click a card in last played hand area
    if (!isPlayer) {
      return false;
    }

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
      {/* <h1 style={{ marginBottom: "64px" }}>
        SelectedCards state is:
        {selectedCards.length > 0 &&
          selectedCards.map((card) => (
            <span>
              {card.rank} of {card.suit},{" "}
            </span>
          ))}
        {selectedCards.length === 0 && "None Selected"}
      </h1> */}
      {/* TODO: debugging code, remove later ^^^^^^^^^*/}
      <HandContainer>
        {cards &&
          cards.map((card, i) => {
            return (
              <Card
                key={card.rank + card.suit}
                data={card}
                handleCardSelect={handleCardSelect}
                selected={isCardSelected(card)}
              />
            );
          })}
        {count &&
          !isPlayer &&
          [...Array(count)].map((card, i) => {
            return <Card key={i + 1} hidden />;
          })}
      </HandContainer>
    </>
  );
};

const HandContainer = styled.div`
  display: flex;
`;
