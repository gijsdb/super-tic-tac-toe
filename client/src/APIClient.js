import axios from "axios"

const POST = "post"
const GET = "get"


const client = axios.create({
  baseURL: "http://localhost:2020/api/v1", // If empty, uses the current origin. If developing locally point to API running locally (HalUpdater)
  json: true,
  validateStatus: function (status) {
    return status < 600; // Reject only if the status code is greater than or equal to 500
  },
});
client.defaults.timeout = 15000;

const APIClient = {

  CreateGame() {
    return this.perform(GET, `/creategame`);
  },

  ListGames() {
    return this.perform(GET, `/games`);
  },

  GetGameBoard(id) {
    return this.perform(GET, `/game?id=` + id);
  },

  UpdateGameBoard(player, square, circle, gameId) {
    return this.perform(GET, `/updateboard?player=` + player + `&square=` + square + `&circle=` + circle + `&gameid=` + gameId);
  },

  JoinGame(id) {
    return this.perform(GET, `/joingame?id=` + id);
  },

  async perform(method, resource, data) {
    return client({
      method,
      url: resource,
      data,
    }).then((req) => {
      //   console.log("req: ", req);
      if (req.status >= 500) {
        if (req.data.Error !== undefined) {
          throw new Error(req.data.Error);
        }
        throw new Error(req.statusText);
      }

      return req.data;
    });
  },
};

export default APIClient;