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
import { config } from "./config";
import { useAlertsStore, useAuthStore } from "./state";
import { IAlertSchema } from "./types";

const route = useRoute();
const { pushAlertToApplication } = useAlertsStore();
const authStore = useAuthStore();
const { token } = storeToRefs(authStore);

const hasConnected = ref(false);

const getElByKey = (obj: Record<any, any>, keys: string[]): any => {
  if (keys.length === 1) return obj[keys[0]];
  return getElByKey(obj[keys[0]], keys.slice(1));
};

watch(token, async (tokenValue) => {
  if (!tokenValue || hasConnected.value) return;
  let socket = new WebSocket(
    `${config.apiWSUrl}/api/ws?token=${token.value}`
  );

  function heartbeat() {
    if (!socket) return;
    if (socket.readyState !== 1) return;
    socket.send("heartbeat");
    setTimeout(heartbeat, 500);
  }

  heartbeat();

  socket.onopen = function (e) {
    console.log("[open] Connection established");
    hasConnected.value = true;
  };
  socket.onmessage = function (event) {
    if (!event.data || !event.data.length) return;
    const alertResponseData = JSON.parse(event.data);
    const applicationId = alertResponseData.application_id;
    const schema = alertResponseData.alert_schema as IAlertSchema;

    const { alert_schema, ...rest } = alertResponseData;

    pushAlertToApplication(applicationId, rest); 

    if (schema?.ID === 0) {
      // TODO: Make this more informative
      invoke("notify_user", {
        title: "New Alert!",
        body: "One of you applications just receive an alertResponseData for the first time!",
      });
      return;
    }

    const titleKeys = schema.Title.split(".");
    const title = getElByKey(alertResponseData, titleKeys);
    const descriptionKeys = schema.Description.split(".");
    const description = getElByKey(alertResponseData, descriptionKeys);
    const linkKeys = schema.Link.split(".");
    const link = getElByKey(alertResponseData, linkKeys);

    invoke("notify_user", { title, body: description });
  };
  socket.onerror = function (event) {
    console.log("[error] A socket error occurred.");
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
