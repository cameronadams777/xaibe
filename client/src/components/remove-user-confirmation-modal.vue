<script lang="ts" setup>
import { XMarkIcon } from "@heroicons/vue/24/outline";
import {
  emptyRemoveUserConfirmationProps,
  useModalStore,
  useToastStore,
} from "src/state";
import { ButtonVariant, ToastType } from "src/types";
import { removeUserFromTeam } from "src/api/teams";

const props = defineProps<{
  isOpen: boolean;
  teamId?: number;
  userId?: number;
}>();

const { setRemoveUserConfirmationProps } = useModalStore();
const { setActiveToast } = useToastStore();

const close = () =>
  setRemoveUserConfirmationProps(emptyRemoveUserConfirmationProps);

const confirm = async () => {
  try {
    if (!props.teamId || !props.userId) return;
    await removeUserFromTeam({ teamId: props.teamId, userId: props.userId });
    setActiveToast({
      type: ToastType.SUCCESS,
      message: "User removed.",
    });
    close();
  } catch (error) {
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred while trying to remove the specified user.",
    });
    close();
  }
};
</script>

<template>
  <base-modal :is-open="isOpen" class="w-7/12">
    <div class="relative h-full flex flex-col justify-center items-center">
      <base-fab-button class="absolute top-2 right-2" @click="close">
        <x-mark-icon class="w-8 h-8" />
      </base-fab-button>
      <h2 class="text-center text-lg w-48 md:w-64 lg:w-72">
        Are you sure you want to remove this user from your team?
      </h2>
      <div class="w-1/2 flex flex-col md:flex-row">
        <base-button
          text="Delete"
          :variant="ButtonVariant.DANGER"
          class="w-full mr-2"
          @click="confirm"
        />
        <base-button
          text="Cancel"
          :variant="ButtonVariant.WHITE"
          class="w-full"
          @click="close"
        />
      </div>
    </div>
  </base-modal>
</template>
