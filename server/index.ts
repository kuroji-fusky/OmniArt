import fastify from "fastify"
import * as dotenv from "dotenv"

const app = async () => {
  dotenv.config()

  const app = await fastify({
    logger: true
  })

  app.get("/", async (_, reply) => {
    return reply.send({
      message: "It me"
    })
  })

  app.listen({ port: Number(process.env.PORT || 4444), host: process.env.HOST || "localhost" }, (err) => {
    if (err) {
      app.log.error(`An oopsie as occured`, err)
      process.exit(1)
    }
  })
}

app()