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
  border-bottom: 3px solid black;
  width: 100%;
  margin-bottom: 64px;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  background: #ffffff;
`;

const Logo = styled.h1`
  padding: 6px 64px;
`;
