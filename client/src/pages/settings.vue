<script lang="ts" setup>
import { storeToRefs } from "pinia";
import { useRouter } from "vue-router";
import { useActiveUserStore } from "../state/active-user";
import TheMainLayout from "../layouts/the-main-layout.vue";

const router = useRouter();
const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);
</script>

<template>
  <the-main-layout>
    <div class="p-4">
      <h2>{{ activeUser?.FirstName }} {{ activeUser?.LastName }}</h2>
    </div>
  </the-main-layout>
</template>
