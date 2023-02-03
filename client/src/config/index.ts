import { z } from "zod";

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL;
const apiWSUrl = import.meta.env.VITE_API_WS_URL;

const configSchema = z.object({
  apiBaseUrl: z.string(),
  apiWSUrl: z.string()
});

export const config: z.infer<typeof configSchema> = configSchema.parse({
  apiBaseUrl,
  apiWSUrl
});


