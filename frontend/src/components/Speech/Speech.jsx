import React, { useState } from "react";
import { Button } from "../Button";
import { AiOutlineAudio, AiFillAudio } from "react-icons/ai";

export const Speech = ({ callbackFn }) => {
  const [isListening, setIsListening] = useState(false);

  // instantiate SpeechRecognition object and setup config
  // this probably only works in Chrome
  const SpeechRecognition =
    window.SpeechRecognition || window.webkitSpeechRecognition;
  const recognition = new SpeechRecognition();
  recognition.interimResults = true;
  recognition.lang = "en-US";
  recognition.onspeechend = function () {
    recognition.stop();
    setIsListening(false);
  };

  const voiceStart = () => {
    setIsListening(true);
    recognition.start();

    // create interim and final transcripts
    // from https://developers.google.com/web/updates/2013/01/Voice-Driven-Web-Apps-Introduction-to-the-Web-Speech-API
    let finalTranscript = "";
    recognition.onresult = (event) => {
      let interimTranscript = "";
      for (let i = event.resultIndex; i < event.results.length; i++) {
        const transcript = event.results[i][0].transcript;
        if (event.results[i].isFinal) finalTranscript += transcript + " ";
        else interimTranscript += transcript;
      }

      // set transcript to state using the setUsername callback we're passing in
      finalTranscript
        ? callbackFn(finalTranscript)
        : callbackFn(interimTranscript);
    };
  };

  return (
    <>
      <Button
        onClick={voiceStart}
        customIcon={isListening ? <AiFillAudio /> : <AiOutlineAudio />}
      />
    </>
  );
};
