import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter } from "react-router-dom";
import { Routes } from "./routes/router";
import { Layout } from "./layout";

let socket = new WebSocket("ws://localhost:8000");
socket.onmessage = function(event) {
    console.log("Socket got, ", event)
}

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
