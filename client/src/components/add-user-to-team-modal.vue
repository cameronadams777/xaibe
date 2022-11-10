<script lang="ts" setup>
import { XMarkIcon } from "@heroicons/vue/24/outline";
import SelectUserList from "./select-user-list.vue";
import {
  emptyAddUserToTeamProps,
  useModalStore,
  useToastStore,
} from "src/state";
import { ButtonVariant, IUser, ToastType } from "src/types";
import { addUserToTeam } from "src/api/teams";
import { onMounted, ref } from "vue";
import { fetchAllUsers } from "src/api/users";

const props = defineProps<{
  isOpen: boolean;
  teamId?: number;
}>();

const { setAddUserToTeamProps } = useModalStore();
const { setActiveToast } = useToastStore();

const usersList = ref<IUser[]>([]);
const userId = ref<number | undefined>(undefined);

const selectUser = (selectedUserId: number) => (userId.value = selectedUserId);

const confirm = async () => {
  try {
    if (!props.teamId || !userId.value) return;
    await addUserToTeam({ teamId: props.teamId, userId: userId.value });
    setActiveToast({
      type: ToastType.SUCCESS,
      message: "User added.",
    });
    close();
  } catch (error) {
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred while trying to add the specified user.",
    });
    close();
  }
};

const close = () => setAddUserToTeamProps(emptyAddUserToTeamProps);

onMounted(async () => {
  try {
    const users = await fetchAllUsers();
    usersList.value = users;
  } catch (error) {
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred while trying to fetch list of users.",
    });
    close();
  }
});
</script>

<template>
  <base-modal :is-open="isOpen" class="w-7/12 h-3/4">
    <div class="relative h-full flex flex-col justify-center items-center">
      <base-fab-button class="absolute top-2 right-2" @click="close">
        <x-mark-icon class="w-8 h-8" />
      </base-fab-button>
      <h2 class="text-center text-lg w-48 md:w-64 lg:w-72">
        Please Select a User
      </h2>
      <div class="w-9/10 mb-4">
        <select-user-list :users="usersList" @on-select="selectUser" />
      </div>
      <div class="w-1/2 flex flex-col md:flex-row">
        <base-button
          text="Add"
          :variant="ButtonVariant.PRIMARY"
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
