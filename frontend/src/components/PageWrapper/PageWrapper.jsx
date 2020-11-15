import React from "react";
import styled from "styled-components";

export const PageWrapper = ({ children }) => {
  return <PageWrapperContainer>{children}</PageWrapperContainer>;
};

const PageWrapperContainer = styled.div`
  display: flex;
  flex-flow: column wrap;
  align-items: center;
  text-align: center;
`;
