import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter } from "react-router-dom";
import { Routes } from "./routes/router";
import { Layout } from "./layout";
import "./shared/styles/global.css";

let endPoint =
  process.env.NODE_ENV === "development"
    ? "ws://localhost:8000"
    : "wss://bigtwo-backend.herokuapp.com";

let socket = new WebSocket(endPoint);

const App = () => {
  return (
    <React.StrictMode>
      <BrowserRouter>
        <Layout>
          <Routes socket={socket} />
        </Layout>
      </BrowserRouter>
    </React.StrictMode>
  );
};

ReactDOM.render(<App />, document.getElementById("root"));
