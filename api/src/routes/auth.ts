import { FastifyPluginCallback } from "fastify";

const auth: FastifyPluginCallback = (_server, _options, done) => {
  done();
};

export default auth;
