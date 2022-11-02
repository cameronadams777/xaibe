<script lang="ts" setup>
import {
  isPermissionGranted,
  requestPermission,
  sendNotification,
} from "@tauri-apps/api/notification";
import { onMounted } from "vue";
import { useRoute } from "vue-router";
import TheGlobalModal from "./components/the-global-modal.vue";
import TheToastMessage from "./components/the-toast-message.vue";
import { io } from "socket.io-client";

const route = useRoute();

onMounted(() => {
  const token = localStorage.getItem("token");
  if (!token) return;
  const socket = io("ws://localhost:5000/api/ws", {
    reconnectionDelayMax: 10000,
    auth: {
      token,
    },
  });

  socket.on("connect", () => console.log("Client connected..."));
  socket.on("data", (data: any) => console.log("Data:", data));
});
</script>

<template>
  <router-view :key="route.fullPath" />
  <the-global-modal />
  <the-toast-message />
</template>

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
