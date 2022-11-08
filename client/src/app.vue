<template>
  <router-view :key="route.fullPath" />
  <the-global-modal />
  <the-toast-message />
</template>

<script lang="ts" setup>
import { useRoute } from "vue-router";
import TheGlobalModal from "./components/the-global-modal.vue";
import TheToastMessage from "./components/the-toast-message.vue";

const route = useRoute();
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
