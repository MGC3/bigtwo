import React from "react";
import styled from "styled-components";
import { GiDiamonds, GiHearts, GiSpades, GiClubs } from "react-icons/gi";

export const Card = ({ data, selected, handleCardSelect, hidden }) => {
  if (hidden) {
    return <CardContainer hidden />;
  }

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

  const icon = iconMaker(data.suit);
  const rank = data.rank;

  return (
    <CardContainer isSelected={selected} onClick={() => handleCardSelect(data)}>
      <Upper>
        <div>{rank}</div>
        <div>{icon}</div>
      </Upper>
      <Lower>
        <div>{rank}</div>
        <div>{icon}</div>
      </Lower>
    </CardContainer>
  );
};

const CardContainer = styled.div`
  height: ${(props) => (props.hidden ? "140px" : "196px")};
  width: ${(props) => (props.hidden ? "100px" : "140px")};
  border: solid;
  padding: 6px 0 0 8px;
  box-sizing: border-box;
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  margin-right: ${(props) => (props.hidden ? "-50px" : "-65px")};
  transition-duration: 0.15s;
  box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.2);
  border-radius: 5px;
  background: ${(props) => (props.isSelected ? "#fffae6" : "#ffffff")};
  transform: ${(props) => (props.isSelected ? "translateY(-32px)" : "")};
  color: black;

  &:hover {
    cursor: ${(props) => (props.isSelected ? "grab" : "pointer")};
  }
`;

const Upper = styled.div``;
const Lower = styled.div`
  align-self: flex-end;
  margin-top: auto;
`;
