<template>
  <base-modal :is-open="isOpen" class="w-7/12 h-3/4">
    <div class="relative h-full flex flex-col justify-center items-center">
      <base-fab-button class="absolute top-2 right-2" @click="close">
        <x-mark-icon class="w-8 h-8" />
      </base-fab-button>
      <h2 class="text-center text-lg w-48 md:w-64 lg:w-72">
        Please Select a User
      </h2>
      <div v-if="selectableUsers.length" class="w-9/10 mb-4">
        <select-user-list
          :active-user-id="activeUser?.ID"
          :selected-user-id="userId"
          :users="selectableUsers"
          @on-select="selectUser"
        />
      </div>
      <div v-if="selectableUsers.length" class="mb-4 w-full flex align-center">
        <hr class="w-3/4 border-0 border-t border-gray-300 mt-3 ml-5 mr-2" />
        <span>or</span>
        <hr class="w-3/4 border-0 border-t border-gray-300 mt-3 mr-5 ml-2" />
      </div>
      <div class="flex flex-col w-9/10 mb-2">
        <label for="newUserEmail" class="font-bold mb-2">Invite New User</label>
        <input
          v-model="newUserEmail"
          placeholder="Email:"
          id="newUserEmail"
          name="newUserEmail"
          type="email"
          class="p-1"
        />
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

<script lang="ts" setup>
import { computed, onMounted, ref } from "vue";
import { storeToRefs } from "pinia";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import SelectUserList from "./select-user-list.vue";
import {
  emptyAddUserToTeamProps,
  useActiveUserStore,
  useModalStore,
  useToastStore,
} from "src/state";
import { ButtonVariant, IUser, ToastType } from "src/types";
import { addUserToTeam } from "src/api/teams";
import { fetchAllUsers, inviteNewUser } from "src/api/users";

const props = defineProps<{
  isOpen: boolean;
  teamId?: number;
}>();

const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);
const { setAddUserToTeamProps } = useModalStore();
const { setActiveToast } = useToastStore();

const newUserEmail = ref("");
const usersList = ref<IUser[]>([]);
const userId = ref<number | undefined>(undefined);

const selectableUsers = computed(
  () =>
    usersList.value?.filter((user) => user.ID !== activeUser?.value?.ID) ?? []
);

const selectUser = (selectedUserId: number) => (userId.value = selectedUserId);

const confirm = async () => {
  try {
    // TODO: Add validation errors here as well
    if (!props.teamId) return;
    if (!userId.value || !newUserEmail.value) return;

    if (newUserEmail.value.length) {
      await inviteNewUser({ email: newUserEmail.value });
      setActiveToast({
        type: ToastType.SUCCESS,
        message: "User invited.",
      });
      return;
    }

    await addUserToTeam({ teamId: props.teamId, userId: userId.value });
    setActiveToast({
      type: ToastType.SUCCESS,
      message: "User added.",
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
    console.error(error);
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred while trying to fetch list of users.",
    });
    close();
  }
});
</script>
