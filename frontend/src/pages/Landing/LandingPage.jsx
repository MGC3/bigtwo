import React from 'react';
import { Landing } from '../../components/Landing';

export const LandingPage = (props) => {
  return (
    <div className="app-container">
      <Landing {...props} />
    </div>
  );
};
