import WebSocket, { WebSocketServer } from "ws";

import { redisSub } from "./index";

export const listenRedis = async (wss: WebSocketServer) => {
  await redisSub.subscribe("notifications", (message) => {
    console.log("redis message", message);

    wss.clients.forEach((client) => {
      if (client.readyState === WebSocket.OPEN) {
        console.log("berhasil terkirim ke client");
        client.send(message);
      }
    });
  });

  console.log("redis subscriber ready");
};
