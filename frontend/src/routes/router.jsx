import React from "react";
import { Switch, Route } from "react-router-dom";
import { LandingPage } from "../pages/Landing";
import { LobbyPage } from "../pages/Lobby";

export const Routes = ({ socket }) => {
  return (
    <Switch>
      <Route
        path="/"
        exact
        render={(props) => <LandingPage {...props} socket={socket} />}
      />
      <Route
        path="/room/:roomId"
        exact
        render={(props) => <LobbyPage {...props} socket={socket} />}
      />
    </Switch>
  );
};
