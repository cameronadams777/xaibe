import { z } from "zod";

const configSchema = z.object({
  apiBaseUrl: z.string(),
  apiWSUrl: z.string(),
  mixpanelToken: z.string(),
  mixpanelDebug: z.boolean(),
  stripePublishableKey: z.string(),
});

export const config: z.infer<typeof configSchema> = configSchema.parse({
  apiBaseUrl: import.meta.env.VITE_API_BASE_URL,
  apiWSUrl: import.meta.env.VITE_API_WS_URL,
  mixpanelToken: import.meta.env.VITE_MIXPANEL_TOKEN,
  mixpanelDebug: import.meta.env.VITE_MIXPANEL_DEBUG === "true",
  stripePublishableKey: import.meta.env.VITE_STRIPE_PUBLISHABLE_KEY,
});
