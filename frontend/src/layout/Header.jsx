import React from "react";
import styled from "styled-components";

export const Header = () => {
  return (
    <Container>
      <Logo>Big Two - insert logo here...</Logo>
    </Container>
  );
};

const Container = styled.div`
  height: 64px;
  width: 100%;
  margin-bottom: 64px;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  background: #ffffff;
  border-bottom: solid 1px #424542;

  // FF7 styles below attrib to: https://codepen.io/Kaizzo/pen/aGWwMM
  box-shadow: 1px 1px #e7dfe7, -1px -1px #e7dfe7, 1px -1px #e7dfe7,
    -1px 1px #e7dfe7, 0 -2px #9c9a9c, -2px 0 #7b757b, 0 2px #424542;
  background: #04009d;
  background: -moz-linear-gradient(top, #04009d 0%, #06004d 100%);
`;

const Logo = styled.h1`
  padding: 6px 64px;
`;
