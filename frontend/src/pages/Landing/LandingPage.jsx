import React, { useState, useEffect } from "react";
import { PageWrapper } from "../../components/PageWrapper";
import { Button } from "../../components/Button";
import { Speech } from "../../components/Speech";
import {
  TextField,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  withStyles,
} from "@material-ui/core/";

const CustomTextField = withStyles({
  root: {
    "& .MuiInputLabel-root": {
      color: "white",
    },
    "& label.Mui-focused": {
      color: "white",
    },
    "& .MuiInputBase-root": {
      color: "white",
    },
    "& .MuiInput-underline:after": {
      borderBottomColor: "white",
    },
    "& .MuiInput-underline:before": {
      borderBottomColor: "white",
    },
  },
})(TextField);

export const LandingPage = ({ history, socket }) => {
  const [roomId, setRoomId] = useState("");
  const [username, setUsername] = useState("");
  const [open, setOpen] = useState(false);

  useEffect(() => {
    socket.onmessage = function (event) {
      const message = JSON.parse(event.data);
      const { type, data } = message;

      switch (type) {
        case "room_created":
          setRoomId(data.room_id);
          setOpen(true);
          break;
        case "room_joined":
          history.push(`/room/${roomId}`);
          break;
        default:
          console.warn("received unknown WS type");
      }
    };
  });

  const handleCreateGameClick = (e) => {
    e.preventDefault();
    socket.send(
      JSON.stringify({
        type: "create_room",
      })
    );
  };

  const handleJoinRoomClick = (e, roomId) => {
    e.preventDefault();

    if (roomId) {
      // TODO: verify if the room id is valid before opening the username dialog
      setOpen(true);
    } else {
      console.log("roomId state is blank");
    }
  };

  const joinRoom = () => {
    socket.send(
      JSON.stringify({
        type: "join_room",
        data: {
          room: roomId,
          name: username,
        },
      })
    );
  };

  // TODO: refactor these two handleChange functions into one state object
  const handleTextFieldChange = (e) => {
    setRoomId(e.target.value);
  };

  const handleUsernameInputChange = (e) => {
    setUsername(e.target.value);
  };

  const handleCloseDialog = () => {
    setOpen(false);
  };

  return (
    <PageWrapper>
      <form
        style={{ maxWidth: "280px", display: "flex", flexDirection: "column" }}
      >
        <Button
          onClick={(e) => handleCreateGameClick(e)}
          text="Create New Game"
        />
        <div>OR </div>
        <CustomTextField
          id="outlined-basic"
          label="Enter room code"
          value={roomId}
          onChange={handleTextFieldChange}
          fullWidth={true}
          inputProps={{
            style: {
              fontSize: "40px",
              textAlign: "center",
            },
            maxLength: 4,
          }}
        />
        <Button
          onClick={(e) => handleJoinRoomClick(e, roomId)}
          text="Join Existing Game"
        />
      </form>
      <Dialog
        open={open}
        onClose={handleCloseDialog}
        aria-labelledby="form-dialog-title"
        maxWidth="xs"
        fullWidth
      >
        <DialogTitle id="form-dialog-title">Choose a Nickname</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            margin="dense"
            id="name"
            label="Nickname"
            type="text"
            fullWidth
            value={username}
            onChange={handleUsernameInputChange}
            inputProps={{
              style: { fontSize: "48px" },
            }}
          />
        </DialogContent>
        <DialogActions>
          <Speech callbackFn={setUsername} />
          <Button onClick={joinRoom} text="Submit" />
        </DialogActions>
      </Dialog>
    </PageWrapper>
  );
};
