<script lang="ts" setup>
import NewElementModal from "./new-element-modal.vue";
import { useModalStore } from "../state/modals";
import { storeToRefs } from "pinia";

const modalStore = useModalStore();

const { isNewElementModalShown } = storeToRefs(modalStore);
</script>

<template>
  <new-element-modal :is-open="isNewElementModalShown" />
</template>
