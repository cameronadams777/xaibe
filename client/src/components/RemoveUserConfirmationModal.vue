<template>
  <BaseModal :is-open="isOpen" class="w-7/12">
    <div class="relative h-full flex flex-col justify-center items-center">
      <BaseFabButton class="absolute top-2 right-2" @click="close">
        <XMarkIcon class="w-8 h-8" />
      </BaseFabButton>
      <h2 class="text-center text-lg w-48 md:w-64 lg:w-72">
        Are you sure you want to remove this user from your team?
      </h2>
      <div class="w-1/2 flex flex-col md:flex-row">
        <BaseButton
          text="Delete"
          :variant="ButtonVariant.DANGER"
          class="w-full mr-2"
          @click="confirm"
        />
        <BaseButton
          text="Cancel"
          :variant="ButtonVariant.WHITE"
          class="w-full"
          @click="close"
        />
      </div>
    </div>
  </BaseModal>
</template>

<script lang="ts" setup>
import { XMarkIcon } from "@heroicons/vue/24/outline";
import {
  emptyRemoveUserConfirmationProps,
  useModalStore,
  useToastStore,
} from "src/state";
import { ButtonVariant, ToastType } from "src/types";
import { removeUserFromTeam } from "src/api/teams";
import { mixpanelWrapper } from "src/tools/mixpanel";

const props = defineProps<{
  isOpen: boolean;
  teamId?: string;
  userId?: string;
}>();

const { setRemoveUserConfirmationProps } = useModalStore();
const { setActiveToast } = useToastStore();

const close = () =>
  setRemoveUserConfirmationProps(emptyRemoveUserConfirmationProps);

const confirm = async () => {
  try {
    if (!props.teamId || !props.userId) return;
    await removeUserFromTeam({ teamId: props.teamId, userId: props.userId });
    
    mixpanelWrapper.client.track("User removed from team");

    setActiveToast({
      type: ToastType.SUCCESS,
      message: "User removed.",
    });
    close();
    /**
     * TODO: We are reloading the page here to ensure that the user that is
     * added is shown in the list upon closing. Need to refactor this with
     * and emitter when the modals refactor happens.
     */
    window.location.reload();
  } catch (error) {
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred while trying to remove the specified user.",
    });
    close();
  }
};
</script>

