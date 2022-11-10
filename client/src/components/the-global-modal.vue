<script lang="ts" setup>
/**
 * TODO: Refactor modals to be able to be used from anywhere using portals so that props
 * can be passed directly rather than through state.
 */
import { storeToRefs } from "pinia";
import NewElementModal from "./new-element-modal.vue";
import DeleteApplicationConfirmationModal from "./delete-application-confirmation-modal.vue";
import DeleteTeamConfirmationModal from "./delete-team-confirmation-modal.vue";
import RemoveUserConfirmationModal from "./remove-user-confirmation-modal.vue";
import { useModalStore } from "../state/modals";

const modalStore = useModalStore();

const {
  isNewElementModalShown,
  isDeleteApplicationConfirmationModalShown,
  isDeleteTeamConfirmationModalShown,
  removeUserConfirmationProps,
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
  <remove-user-confirmation-modal
    :is-open="removeUserConfirmationProps.isOpen"
    :team-id="removeUserConfirmationProps.teamId"
    :user-id="removeUserConfirmationProps.userId"
  />
</template>
