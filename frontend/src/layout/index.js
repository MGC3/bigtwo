import React from "react";
import styled from "styled-components";
import { Header } from "./Header";

export const Layout = ({ children }) => {
  return (
    <AppContainer>
      <Header />
      {children}
    </AppContainer>
  );
};

const AppContainer = styled.div`
  display: flex;
  flex-flow: column wrap;
  align-items: center;
  text-align: center;
  width: 100%;
  height: 100%;
  min-height: 100vh;
`;
