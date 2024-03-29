<template>
  <router-view :key="route.fullPath" />
  <TheGlobalModal />
  <TheToastMessage />
</template>

<script lang="ts" setup>
import { invoke } from "@tauri-apps/api";
import { storeToRefs } from "pinia";
import { watch, ref } from "vue";
import { useRoute } from "vue-router";
import TheGlobalModal from "./components/TheGlobalModal.vue";
import TheToastMessage from "./components/TheToastMessage.vue";
import { config } from "./config";
import { useAlertsStore, useAuthStore, useToastStore } from "./state";
import { AlertSchema, ToastType } from "./types";

const route = useRoute();
const { pushAlertToApplication } = useAlertsStore();
const { setActiveToast } = useToastStore();
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

  function heartbeat(): void {
    if (!socket) return;
    if (socket.readyState !== 1) return;
    socket.send("heartbeat");
    setTimeout(heartbeat, 500);
  }

  heartbeat();

  socket.onopen = function (): void {
    console.log("[open] Connection established");
    hasConnected.value = true;
  };
  socket.onmessage = function (event): void {
    if (!event.data || !event.data.length) return;
    const alertResponseData = JSON.parse(event.data);
    const applicationId = alertResponseData.application_id;
    const schema = alertResponseData.alert_schema as AlertSchema;

    const { alert_schema, ...rest } = alertResponseData;

    pushAlertToApplication(applicationId, rest);

    if (!schema.title.length && !schema.link.length && !schema.description.length) {
      // TODO: Make this more informative
      notifyUser({
        title: "New Alert!",
        body: "One of you applications just receive an alertResponseData for the first time!",
      });
      return;
    }

    const titleKeys = schema.title.split(".");
    const title = getElByKey(alertResponseData, titleKeys);
    const descriptionKeys = schema.description.split(".");
    const description = getElByKey(alertResponseData, descriptionKeys);

    notifyUser({ title, body: description });
  };
  socket.onerror = function () {
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

  const notifyUser = ({ title, body }: { title: string, body: string }) => {
    console.log(title, body);
    if(config.appEnv === "local") {
      setActiveToast({
        type: ToastType.INFO,
        message: body,
      });
      return;
    }
    invoke("notify_user", { title, body });
  }
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
