<script lang="ts" setup>
import { useRoute, useRouter } from "vue-router";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import { useModalStore } from "../state/modals";
import { deleteTeam } from "../api/teams";
import { ButtonVariant } from "../types";

defineProps<{ isOpen: boolean }>();

const router = useRouter();
const route = useRoute();
const { setIsDeleteTeamConfirmationModalShown } = useModalStore();

const attemptToDeleteTeam = async () => {
  try {
    const teamId = parseInt(route.params.teamId as string);
    await deleteTeam({ teamId });
    setIsDeleteTeamConfirmationModalShown(false);
    // TODO: Display a successful toast message
    router.push("/");
  } catch (error) {
    // TODO: Display a error toast message
    setIsDeleteTeamConfirmationModalShown(false);
  }
};
</script>

<template>
  <base-modal :is-open="isOpen" class="w-7/12">
    <div class="relative h-full flex flex-col justify-center items-center">
      <base-fab-button
        class="absolute top-2 right-2"
        @click="setIsDeleteTeamConfirmationModalShown(false)"
      >
        <x-mark-icon class="w-8 h-8" />
      </base-fab-button>
      <h2 class="text-lg">Are you sure you want to delete this team?</h2>
      <div class="flex">
        <base-button
          text="Delete"
          :variant="ButtonVariant.DANGER"
          @click="attemptToDeleteTeam"
        />
        <base-button
          text="Cancel"
          :variant="ButtonVariant.WHITE"
          @click="setIsDeleteTeamConfirmationModalShown(false)"
        />
      </div>
    </div>
  </base-modal>
</template>
