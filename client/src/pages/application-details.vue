<script lang="ts" setup>
import { storeToRefs } from "pinia";
import { ref } from "vue";
import { useRoute } from "vue-router";
import { useActiveUserStore } from "../state/active-user";
import TheMainLayout from "../layouts/the-main-layout.vue";
import { IApplication } from "../types";

const route = useRoute();

const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);
</script>

<template>
  <the-main-layout>
    <div>
      <h2>Hello World</h2>
    </div>
  </the-main-layout>
</template>
