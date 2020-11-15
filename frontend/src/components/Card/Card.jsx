import React, { Component } from "react";
import styled from "styled-components";
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
      <CardContainer
        isSelected={this.props.selected}
        onClick={() => this.props.handleCardSelect(this.state.data)}
      >
        <Upper>
          <div>{this.state.data.rank}</div>
          <div>{icon}</div>
        </Upper>
        <Lower>
          <div>{this.state.data.rank}</div>
          <div>{icon}</div>
        </Lower>
      </CardContainer>
    );
  }
}

const CardContainer = styled.div`
  height: 210px;
  width: 150px;
  border: solid;
  padding: 6px 0 0 8px;
  box-sizing: border-box;
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  margin-right: -65px;
  transition-duration: 0.15s;
  box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.2);
  border-radius: 5px;
  background: ${(props) => (props.isSelected ? "#fffae6" : "#ffffff")};
  transform: ${(props) => (props.isSelected ? "translateY(-50px)" : "")};

  &:hover {
    cursor: ${(props) => (props.isSelected ? "grab" : "pointer")};
  }
`;

const Upper = styled.div``;
const Lower = styled.div`
  align-self: flex-end;
  margin-top: auto;
`;
