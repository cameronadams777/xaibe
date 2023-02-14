<script lang="ts" setup>
import { storeToRefs } from "pinia";
import { useActiveUserStore } from "../state/active-user";
import TheMainLayout from "../layouts/the-main-layout.vue";

const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);
</script>

<template>
  <the-main-layout>
    <div class="p-4">
      <h2>{{ activeUser?.firstName }} {{ activeUser?.lastName }}</h2>
    </div>
  </the-main-layout>
</template>
