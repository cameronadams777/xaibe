<script lang="ts" setup>
import { useRoute } from "vue-router";
import TheGlobalModal from "./components/the-global-modal.vue";

const route = useRoute();
</script>

<template>
  <router-view :key="route.fullPath" />
  <the-global-modal />
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
