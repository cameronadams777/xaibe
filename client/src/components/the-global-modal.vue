<script lang="ts" setup>
/**
 * TODO: Refactor modals to be able to be used from anywhere using portals so that props
 * can be passed directly rather than through state.
 */
import { storeToRefs } from "pinia";
import AddUserToTeamModal from "./add-user-to-team-modal.vue";
import DeleteApplicationConfirmationModal from "./delete-application-confirmation-modal.vue";
import DeleteTeamConfirmationModal from "./delete-team-confirmation-modal.vue";
import NewElementModal from "./new-element-modal.vue";
import RemoveUserConfirmationModal from "./remove-user-confirmation-modal.vue";
import { useModalStore } from "../state/modals";

const modalStore = useModalStore();

const {
  addUserToTeamProps,
  isNewElementModalShown,
  isDeleteApplicationConfirmationModalShown,
  isDeleteTeamConfirmationModalShown,
  removeUserConfirmationProps,
} = storeToRefs(modalStore);
</script>

<template>
  <add-user-to-team-modal
    :is-open="addUserToTeamProps.isOpen"
    :team-id="addUserToTeamProps.teamId"
  />
  <delete-application-confirmation-modal
    :is-open="isDeleteApplicationConfirmationModalShown"
  />
  <delete-team-confirmation-modal
    :is-open="isDeleteTeamConfirmationModalShown"
  />
  <new-element-modal :is-open="isNewElementModalShown" />
  <remove-user-confirmation-modal
    :is-open="removeUserConfirmationProps.isOpen"
    :team-id="removeUserConfirmationProps.teamId"
    :user-id="removeUserConfirmationProps.userId"
  />
</template>
