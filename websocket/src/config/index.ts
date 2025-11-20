import dotenv from "dotenv"
dotenv.config()

export const config = {
    port: process.env.PORT || 8080,
    redisUrl: process.env.REDIS_URL || "redis://redis-pubsub:6379"
}

