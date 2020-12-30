import React from "react";
import styled from "styled-components";

export const PageWrapper = ({ children, customStyles }) => {
  return (
    <PageWrapperContainer style={customStyles}>{children}</PageWrapperContainer>
  );
};

const PageWrapperContainer = styled.div`
  display: flex;
  flex-flow: column wrap;
  align-items: center;
  text-align: center;
  background: white;
  padding: 16px 32px;
  border-radius: 8px;

  // FF7 styles below attrib to: https://codepen.io/Kaizzo/pen/aGWwMM
  border: solid 1px #424542;
  box-shadow: 1px 1px #e7dfe7, -1px -1px #e7dfe7, 1px -1px #e7dfe7,
    -1px 1px #e7dfe7, 0 -2px #9c9a9c, -2px 0 #7b757b, 0 2px #424542;
  background: #04009d;
  background: -moz-linear-gradient(top, #04009d 0%, #06004d 100%);
  background: -webkit-gradient(
    linear,
    left top,
    left bottom,
    color-stop(0%, #04009d),
    color-stop(100%, #06004d)
  );
`;
