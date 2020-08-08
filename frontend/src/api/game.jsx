import axios from "axios";

const apiUrl = "http://localhost:8000/rooms/";

export const createGame = () => {
  return axios({
    url: apiUrl,
    method: "POST",
  });
};

export const joinGame = (formData) => {
  return axios({
    url: `${apiUrl}formData`,
    method: "GET",
  });
};
