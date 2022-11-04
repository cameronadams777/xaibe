<script lang="ts" setup>
import { onMounted } from "vue";
import { useRoute } from "vue-router";
import TheGlobalModal from "./components/the-global-modal.vue";
import TheToastMessage from "./components/the-toast-message.vue";

const route = useRoute();

onMounted(async () => {
  const token = localStorage.getItem("token");
  const socket = new WebSocket(`ws://localhost:5000/api/ws?token=${token}`);
  socket.onclose = () => console.log("Connection closed.");
  socket.onmessage = (evt) => {
    let messages = evt.data.split("\n");
    console.log(messages);
  };

  socket.addEventListener("error", (event) => {
    console.error("WebSocket error: ", event);
  });
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
