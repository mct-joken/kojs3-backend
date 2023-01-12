import { WebSocketServer } from "ws";

const ws = new WebSocketServer({ port: 8090 });

/* eslint @typescript-eslint/no-explicit-any:0 */
let Connections: Array<any> = [];

ws.on("connection", (w) => {
  Connections.push(w);
  w.on("close", () => {
    Connections = Connections.filter((conn, i) => {
      return conn !== i;
    });
  });
});

export async function sendMessageToWebSocketStream(mes: string) {
  ws.clients.forEach((client) => {
    client.send(JSON.stringify(mes));
  });
}
