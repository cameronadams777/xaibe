<template>
  <BaseModal :is-open="isOpen" class="w-7/12">
    <div class="relative h-full flex flex-col justify-center items-center">
      <BaseFabButton
        class="absolute top-2 right-2"
        @click="setIsDeleteTeamConfirmationModalShown(false)"
      >
        <XMarkIcon class="w-8 h-8" />
      </BaseFabButton>
      <h2 class="text-lg">Are you sure you want to delete this team?</h2>
      <div class="flex">
        <BaseButton
          text="Delete"
          :variant="ButtonVariant.DANGER"
          @click="attemptToDeleteTeam"
        />
        <BaseButton
          text="Cancel"
          :variant="ButtonVariant.WHITE"
          @click="setIsDeleteTeamConfirmationModalShown(false)"
        />
      </div>
    </div>
  </BaseModal>
</template>

<script lang="ts" setup>
import { useRoute, useRouter } from "vue-router";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import { useModalStore, useToastStore } from "src/state";
import { deleteTeam } from "src/api/teams";
import { ButtonVariant, ToastType } from "src/types";
import { mixpanelWrapper } from "src/tools/mixpanel";

defineProps<{ isOpen: boolean }>();

const router = useRouter();
const route = useRoute();
const { setIsDeleteTeamConfirmationModalShown } = useModalStore();
const { setActiveToast } = useToastStore();

const attemptToDeleteTeam = async () => {
  try {
    const teamId = route.params.teamId as string;
    await deleteTeam({ teamId });
    setIsDeleteTeamConfirmationModalShown(false);

    mixpanelWrapper.client.track("Team deleted")

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

