import storage from "node-persist";
import { SupportedApplication } from "../types";
import { FastifyPluginCallback } from "fastify";

const webhooks: FastifyPluginCallback = (server, options, done) => {
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

  done();
};

export default webhooks;
