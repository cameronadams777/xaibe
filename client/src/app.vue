<script lang="ts" setup>
import TheGlobalModal from "./components/the-global-modal.vue";
</script>

<template>
  <router-view />
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
