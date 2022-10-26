<script lang="ts" setup>
import { storeToRefs } from "pinia";
import NewElementModal from "./new-element-modal.vue";
import DeleteApplicationConfirmationModal from "./delete-application-confirmation-modal.vue";
import DeleteTeamConfirmationModal from "./delete-team-confirmation-modal.vue";
import { useModalStore } from "../state/modals";

const modalStore = useModalStore();

const {
  isNewElementModalShown,
  isDeleteApplicationConfirmationModalShown,
  isDeleteTeamConfirmationModalShown,
} = storeToRefs(modalStore);
</script>

<template>
  <new-element-modal :is-open="isNewElementModalShown" />
  <delete-application-confirmation-modal
    :is-open="isDeleteApplicationConfirmationModalShown"
  />
  <delete-team-confirmation-modal
    :is-open="isDeleteTeamConfirmationModalShown"
  />
</template>
