<template>
  <router-view :key="route.fullPath" />
  <the-global-modal />
  <the-toast-message />
</template>

<script lang="ts" setup>
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

watch(token, async (tokenValue) => {
  if (!tokenValue || hasConnected.value) return;

  let socket = new WebSocket(`ws://localhost:5000/api/ws?token=${token.value}`);

  socket.onopen = function (e) {
    alert("[open] Connection established");
    alert("Sending to server");
    hasConnected.value = true;
  };

  socket.onmessage = function (event) {
    alert(`[message] Data received from server: ${event.data}`);
  };

  socket.onclose = function (event) {
    if (event.wasClean) {
      alert(
        `[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`
      );
    } else {
      console.log(event);
      // e.g. server process killed or network down
      // event.code is usually 1006 in this case
      alert("[close] Connection died");
    }
  };

  socket.onerror = function (error) {
    alert(`[error]`);
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
