import WebSocket, {WebSocketServer} from "ws";
import { config } from "../config";

export const initWebsocket = () => {
    const wss = new WebSocketServer({
        port: +config.port
    })

    wss.on("connection", (ws) => {
        console.log("fe connected")
        ws.on("close", () => console.log("FE disconnected"))
    })

    return wss;

}