import { z } from "zod";

const configSchema = z.object({
  appEnv: z.enum(["local", "staging", "production"]),
  apiBaseUrl: z.string(),
  apiWSUrl: z.string(),
  mixpanelToken: z.string(),
  mixpanelDebug: z.boolean(),
  sentryDSN: z.string(),
  stripePublishableKey: z.string(),
});

export const config: z.infer<typeof configSchema> = configSchema.parse({
  appEnv: import.meta.env.VITE_APP_ENV,
  apiBaseUrl: import.meta.env.VITE_API_BASE_URL,
  apiWSUrl: import.meta.env.VITE_API_WS_URL,
  mixpanelToken: import.meta.env.VITE_MIXPANEL_TOKEN,
  mixpanelDebug: import.meta.env.VITE_MIXPANEL_DEBUG === "true",
  sentryDSN: import.meta.env.VITE_SENTRY_DSN,
  stripePublishableKey: import.meta.env.VITE_STRIPE_PUBLISHABLE_KEY,
});
