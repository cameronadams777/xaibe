<script lang="ts" setup>
import { useRoute, useRouter } from "vue-router";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import { useModalStore, useToastStore } from "../state";
import { deleteTeam } from "../api/teams";
import { ButtonVariant, ToastType } from "../types";

defineProps<{ isOpen: boolean }>();

const router = useRouter();
const route = useRoute();
const { setIsDeleteTeamConfirmationModalShown } = useModalStore();
const { setActiveToast } = useToastStore();

const attemptToDeleteTeam = async () => {
  try {
    const teamId = parseInt(route.params.teamId as string);
    await deleteTeam({ teamId });
    setIsDeleteTeamConfirmationModalShown(false);
    setActiveToast({
      type: ToastType.SUCCESS,
      message: "Team deleted!",
    });
    router.push("/");
  } catch (error) {
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred while trying to delete the specified team.",
    });
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
