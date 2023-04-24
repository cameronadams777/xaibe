import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "./app.vue";
import router from "./router";
import { config } from "./config";
import { mixpanelWrapper } from "./tools/mixpanel";
import * as Sentry from "@sentry/vue";
import "uno.css";

const pinia = createPinia();
const app = createApp(App);

// Make all components that follow the `Base*.vue` naming convention globally available
const components = import.meta.glob("./components/Base*.vue", { eager: true });
Object.entries(components).forEach(([path, definition]) => {
  // Get name of component, based on filename
  // "./components/Fruits.vue" will become "Fruits"
  const componentName = path
    .split("/")
    .pop()
    ?.replace(/\.\w+$/, "");

  if (!componentName) return;

  // Register component on this Vue instance
  app.component(componentName, (definition as any).default);
});

mixpanelWrapper.setup();

Sentry.init({
  app,
  dsn: config.sentryDSN,
  environment: config.appEnv,
  integrations: [
    new Sentry.BrowserTracing({
      routingInstrumentation: Sentry.vueRouterInstrumentation(router),
    }),
  ],
  // Set tracesSampleRate to 1.0 to capture 100%
  // of transactions for performance monitoring.
  // We recommend adjusting this value in production
  tracesSampleRate: 1.0,
});

app.use(router);
app.use(pinia);

app.mount("#app");
