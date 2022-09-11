import fastify from "fastify";
import storage from "node-persist";
import { auth, webhooks } from "./routes";

storage
  .init({
    dir: "./store",
  })
  .then(async () => {
    const server = fastify({
      logger: process.env.NODE_ENV === "production",
    });

    server.get("/health", async (_, reply) => {
      return reply.status(200).send({ message: "I am health! 🚀" });
    });

    server.register(auth);
    server.register(webhooks);

    server.listen({ port: 5000 }, (err, address) => {
      if (err) {
        console.error(err);
        process.exit(0);
      }

      console.log(`Server listening at ${address}`);
    });
  });
