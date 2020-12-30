import React, { useState, useEffect } from "react";
import styled from "styled-components";
import finger from "../../assets/images/finger.png";

export const Button = ({
  onClick,
  text,
  classes,
  customIcon,
  disabled,
  ...props
}) => {
  const [isHovering, setIsHovering] = useState(false);

  // hack: fixes bug where hovering remains true after a users turn
  useEffect(() => {
    setIsHovering(false);
  }, [disabled]);

  return (
    <ButtonWrapper
      className={classes}
      onClick={onClick}
      customIcon={customIcon}
      disabled={disabled}
      onMouseEnter={() => setIsHovering(true)}
      onMouseLeave={() => setIsHovering(false)}
      {...props}
    >
      {customIcon ? (
        customIcon
      ) : (
        <FingerIcon src={finger} isHovering={isHovering && !disabled} />
      )}
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
  padding: ${(props) => (props.customIcon ? "24px" : "0 40px 0 0")};
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
    cursor: ${(props) => (props.disabled ? "not-allowed" : "pointer")};
  }

  &:active {
    transform: translate(4px, 4px);
    box-shadow: 0 0 0 0 black;
  }

  &:focus {
    outline: none;
  }
`;

const FingerIcon = styled.img`
  width: 32px;
  margin-right: 8px;
  visibility: ${(props) => (props.isHovering ? "visible" : "hidden")};
`;
