import React, { Component } from "react";
import { Route, Switch } from "react-router-dom";
import Landing from "./Landing";
import Lobby from "./Lobby";
import "../css/app.css";

class App extends Component {
  render() {
    return (
      <div className="app-container">
        <Switch>
          <Route exact path="/" render={() => <Landing />}></Route>
          <Route
            exact
            path="/room/:roomId"
            render={(props) => <Lobby {...props} />}
          ></Route>
        </Switch>
      </div>
    );
  }
}

export default App;
