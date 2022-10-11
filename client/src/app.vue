<script lang="ts" setup>
import TheMainLayout from "./layouts/the-main-layout.vue";
</script>

<template>
  <the-main-layout>
    <router-view />
  </the-main-layout>
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
