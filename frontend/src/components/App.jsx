import React, { Component } from "react";
import { Route, Switch } from "react-router-dom";
import Landing from "./Landing";
import "../css/app.css";

class App extends Component {
  render() {
    return (
      <div className="app-container">
        <Switch>
          <Route exact path="/" render={() => <Landing />}></Route>
        </Switch>
      </div>
    );
  }
}

export default App;
