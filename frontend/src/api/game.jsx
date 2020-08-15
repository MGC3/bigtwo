import axios from "axios";

const apiUrl = "http://localhost:8000";

export const createGame = () => {
  return axios({
    url: `${apiUrl}/rooms`,
    method: "POST",
  });
};

export const joinGame = (formData) => {
  return axios({
    url: `${apiUrl}/rooms/${formData}`,
    method: "GET",
  });
};
