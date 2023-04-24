<template>
  <AddUserToTeamModal
    v-if="addUserToTeamProps.isOpen"
    :is-open="addUserToTeamProps.isOpen"
    :team-id="addUserToTeamProps.teamId"
  />
  <DeleteApplicationConfirmationModal
    v-if="isDeleteApplicationConfirmationModalShown"
    :is-open="isDeleteApplicationConfirmationModalShown"
  />
  <DeleteTeamConfirmationModal
    v-if="isDeleteTeamConfirmationModalShown"
    :is-open="isDeleteTeamConfirmationModalShown"
  />
  <NewElementModal
    v-if="isNewElementModalShown"
    :is-open="isNewElementModalShown" 
  />
  <RemoveUserConfirmationModal
    v-if="removeUserConfirmationProps.isOpen"
    :is-open="removeUserConfirmationProps.isOpen"
    :team-id="removeUserConfirmationProps.teamId"
    :user-id="removeUserConfirmationProps.userId"
  />
</template>

<script lang="ts" setup>
/**
 * TODO: Refactor modals to be able to be used from anywhere using portals so that props
 * can be passed directly rather than through state.
 */
import { storeToRefs } from "pinia";
import AddUserToTeamModal from "./AddUserToTeamModal.vue";
import DeleteApplicationConfirmationModal from "./DeleteApplicationConfirmationModal.vue";
import DeleteTeamConfirmationModal from "./DeleteTeamConfirmationModal.vue";
import NewElementModal from "./NewElementModal.vue";
import RemoveUserConfirmationModal from "./RemoveUserConfirmationModal.vue";
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

