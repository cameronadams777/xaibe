<script lang="ts" setup>
import { useRoute, useRouter } from "vue-router";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import { useApplicationsStore, useModalStore, useToastStore } from "src/state";
import { ButtonVariant, ToastType } from "src/types";
import { mixpanelWrapper } from "src/tools/mixpanel";

defineProps<{ isOpen: boolean }>();

const router = useRouter();
const route = useRoute();
const { setIsDeleteApplicationConfirmationModalShown } = useModalStore();
const { deleteApplication } = useApplicationsStore();
const { setActiveToast } = useToastStore();

const attemptToDeleteApplication = async () => {
  try {
    const applicationId = route.params.applicationId as string;
    await deleteApplication(applicationId);
    setIsDeleteApplicationConfirmationModalShown(false);
    setActiveToast({
      type: ToastType.SUCCESS,
      message: "Application deleted!",
    });

    mixpanelWrapper.client.track("Application deleted");

    router.push("/");
  } catch (error) {
    setActiveToast({
      type: ToastType.ERROR,
      message:
        "An error occurred while trying to delete the specified application.",
    });
    setIsDeleteApplicationConfirmationModalShown(false);
  }
};
</script>

<template>
  <base-modal :is-open="isOpen" class="w-7/12">
    <div
      class="relative w-full h-full flex flex-col justify-center items-center"
    >
      <base-fab-button
        class="absolute top-2 right-2"
        @click="setIsDeleteApplicationConfirmationModalShown(false)"
      >
        <x-mark-icon />
      </base-fab-button>
      <h2 class="text-lg">Are you sure you want to delete this application?</h2>
      <div class="flex">
        <base-button
          text="Delete"
          :variant="ButtonVariant.DANGER"
          @click="attemptToDeleteApplication"
        />
        <base-button
          text="Cancel"
          :variant="ButtonVariant.WHITE"
          @click="setIsDeleteApplicationConfirmationModalShown(false)"
        />
      </div>
    </div>
  </base-modal>
</template>
