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
    v-if="addUserToTeamProps.isOpen"
    :is-open="addUserToTeamProps.isOpen"
    :team-id="addUserToTeamProps.teamId"
  />
  <delete-application-confirmation-modal
    v-if="isDeleteApplicationConfirmationModalShown"
    :is-open="isDeleteApplicationConfirmationModalShown"
  />
  <delete-team-confirmation-modal
    v-if="isDeleteTeamConfirmationModalShown"
    :is-open="isDeleteTeamConfirmationModalShown"
  />
  <new-element-modal 
    v-if="isNewElementModalShown"
    :is-open="isNewElementModalShown" 
  />
  <remove-user-confirmation-modal
    v-if="removeUserConfirmationProps.isOpen"
    :is-open="removeUserConfirmationProps.isOpen"
    :team-id="removeUserConfirmationProps.teamId"
    :user-id="removeUserConfirmationProps.userId"
  />
</template>
