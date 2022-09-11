import fastify from "fastify";
import storage from "node-persist";

enum SupportedApplication {
  AIRBRAKE = "airbrake",
}

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

    interface IWebhookQueryString {
      application: SupportedApplication;
    }
    server.post<{ Querystring: IWebhookQueryString }>(
      "/webhook",
      async (request, reply) => {
        const { application } = request.query;

        const alert = request.body;

        switch (application) {
          case SupportedApplication.AIRBRAKE:
            break;
        }

        const currentStore = (await storage.getItem("application")) || [];

        await storage.setItem("application", [...currentStore, alert]);

        return reply.status(200).send();
      }
    );

    server.listen({ port: 5000 }, (err, address) => {
      if (err) {
        console.error(err);
        process.exit(0);
      }

      console.log(`Server listening at ${address}`);
    });
  });
