import { initWebsocket } from "./ws";
import { initRedis } from "./redis";
import { listenRedis } from "./redis/sub";

const start = async () => {
  await initRedis();
  const wss = initWebsocket();
  await listenRedis(wss);

  console.log("websocket + redis listener started");
};

start();
