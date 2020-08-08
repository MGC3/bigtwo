import React from "react";
import { createGame } from "./api/game";

function App() {
  const handleCreateGame = () => {
    // create game
    createGame();
    // if success, join game
  };

  return (
    <div>
      <button onClick={handleCreateGame}>Create Game</button>
    </div>
  );
}

export default App;
