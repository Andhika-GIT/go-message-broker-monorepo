import { createClient } from "redis"
import { config } from "../config"

export const redisSub = createClient({
    url: config.redisUrl
})

redisSub.on("error", (err) => console.error("Redis error: ",err))

export const initRedis = async () => {
    await redisSub.connect();
    console.log("redis subscriber connected")
}