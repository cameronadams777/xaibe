<template>
  <router-view :key="route.fullPath" />
  <the-global-modal />
  <the-toast-message />
</template>

<script lang="ts" setup>
import { invoke } from "@tauri-apps/api";
import { storeToRefs } from "pinia";
import { watch, ref } from "vue";
import { useRoute } from "vue-router";
import TheGlobalModal from "./components/the-global-modal.vue";
import TheToastMessage from "./components/the-toast-message.vue";
import { useAuthStore } from "./state";

const route = useRoute();
const authStore = useAuthStore();
const { token } = storeToRefs(authStore);

const hasConnected = ref(false);

const getElByKey = (obj: Record<any, any>, keys: string[]): any => {
  if (keys.length === 1) return obj[keys[0]];
  return getElByKey(obj[keys[0]], keys.slice(1));
};

watch(token, async (tokenValue) => {
  if (!tokenValue || hasConnected.value) return;
  let socket = new WebSocket(`ws://localhost:5000/api/ws?token=${token.value}`);
  socket.onopen = function (e) {
    console.log("[open] Connection established");
    hasConnected.value = true;
  };
  socket.onmessage = function (event) {
    const alert = JSON.parse(event.data);
    const schema = alert.alert_schema;

    if (!schema) {
      // TODO: Make this more informative
      invoke("notify_user", {
        title: "New Alert!",
        body: "One of you applications just receive an alert for the first time!",
      });
      return;
    }

    const titleKeys = schema.Title.split(".");
    const title = getElByKey(alert, titleKeys);
    const descriptionKeys = schema.Description.split(".");
    const body = getElByKey(alert, descriptionKeys);

    invoke("notify_user", { title, body });
  };
  socket.onclose = function (event) {
    if (event.wasClean) {
      console.log(
        `[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`
      );
      return;
    }
    console.log("[close] Connection died: ", event);
  };
});
</script>

<style>
html,
body,
#app {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
}

* {
  font-family: "Source Sans Pro", sans-serif;
}
</style>
