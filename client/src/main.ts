import path from "path";
import { createApp } from "vue";
import App from "./app.vue";
import router from "./router";
import { createPinia } from "pinia";
import "uno.css";

const pinia = createPinia();
const app = createApp(App);

const components = import.meta.glob("./components/base-*.vue", { eager: true });

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

app.use(router);
app.use(pinia);

app.mount("#app");
