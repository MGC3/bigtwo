import React from "react";
import styled from "styled-components";

export const Button = ({ onClick, text, classes, icon }) => {
  return (
    <ButtonWrapper className={classes} onClick={onClick}>
      {icon}
      {text}
    </ButtonWrapper>
  );
};

const ButtonWrapper = styled.button`
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 64px;
  padding: 0 32px;
  background-color: #fff;
  border-radius: 8px;
  font-weight: 700;
  font-size: 20px;
  line-height: 30px;
  box-sizing: border-box;
  border: 2px solid black;
  box-shadow: 4px 4px 0 0 black;
  transform: translate(0, 0);
  transition: all 0.2s ease;
  margin: 16px 0;
  &:hover {
    cursor: pointer;
  }

  &:active {
    transform: translate(4px, 4px);
    box-shadow: 0 0 0 0 black;
  }

  &:focus {
    outline: none;
  }
`;
