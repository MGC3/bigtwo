import React from 'react';
import Lobby from '../../components/Lobby';

export const LobbyPage = (props) => {
  return (
    <div className="app-container">
      <Lobby {...props} />
    </div>
  );
};
