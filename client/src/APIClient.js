import axios from "axios"

const GET = "get"

const client = axios.create({
  baseURL: "http://localhost:1323/", // If empty, uses the current origin. If developing locally point to API running locally
  json: true,
  validateStatus: function (status) {
    return status < 600;
  },
  withCredentials: true,
});
client.defaults.timeout = 15000;

const APIClient = {

  CreateGame(playerId) {
    return this.perform(GET, `/game/create?player=${playerId}`);
  },

  CreatePlayer() {
    return this.perform(GET, `/player/create`);
  },

  GetPlayer(playerId) {
    return this.perform(GET, `/player/${playerId}`)
  },

  CreateSession(playerId) {
    return this.perform(GET, `/session/create?player=${playerId}`);
  },

  ListGames() {
    return this.perform(GET, `/games`);
  },

  GetGame(id) {
    return this.perform(GET, `/game/${id}`);
  },

  UpdateGameBoard(player, square, circle, gameId) {
    return this.perform(GET, `/game/${gameId}/board/update?player=` + player + `&square=` + square + `&circle=` + circle);
  },

  RemoveCircle(player, square, circle, gameId) {
    return this.perform(GET, `/game/${gameId}/board/circle/${circle}/remove?player=` + player + `&square=` + square);
  },

  RollDice(dice1, dice2, gameId) {
    return this.perform(GET, `/game/${gameId}/roll?dice1=` + dice1 + `&dice2=` + dice2);
  },

  JoinGame(id, player) {
    return this.perform(GET, `/game/join?id=` + id + `&player=` + player);
  },

  LeaveGame(gameId, playerId) {
    return this.perform(GET, `/game/${gameId}/leave?&player=` + playerId);
  },

  OAuthLogin() {
    return this.perform(GET, '/login');
  },

  Logout() {
    return this.perform(GET, '/logout')
  },

  GetHighscores() {
    return this.perform(GET, '/highscores')
  },

  async perform(method, resource, data) {
    return client({
      method,
      url: resource,
      data,
    }).then((req) => {
      if (req.status >= 400) {
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