<template>
  <base-modal :is-open="isOpen" class="w-7/12 h-3/4">
    <div class="relative h-full flex flex-col justify-center items-center">
      <base-fab-button class="absolute top-2 right-2" @click="close">
        <x-mark-icon class="w-8 h-8" />
      </base-fab-button>
      <h2 class="text-center text-lg w-48 md:w-64 lg:w-72">
        Please Select a User
      </h2>
      <div v-if="usersList.length" class="w-9/10 mb-4">
        <select-user-list
          :active-user-id="activeUser?.id"
          :selected-user-id="userId"
          :users="usersList"
          @on-select="selectUser"
        />
      </div>
      <div v-if="usersList.length" class="mb-4 w-full flex align-center">
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
import { onMounted, ref } from "vue";
import { storeToRefs } from "pinia";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import SelectUserList from "./select-user-list.vue";
import {
  emptyAddUserToTeamProps,
  useActiveUserStore,
  useModalStore,
  useToastStore,
  useGalataUsersStore,
} from "src/state";
import { ButtonVariant, User, ToastType } from "src/types";
import { inviteNewUser } from "src/api/users";
import { inviteExistingUserToTeam } from 'src/api/teams';
import { mixpanelWrapper } from "src/tools/mixpanel";

const props = defineProps<{
  isOpen: boolean;
  teamId?: string;
}>();

const { getAllUsers } = useGalataUsersStore();
const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);
const { setAddUserToTeamProps } = useModalStore();
const { setActiveToast } = useToastStore();

const newUserEmail = ref("");
const usersList = ref<User[]>([]);
const userId = ref<string | undefined>(undefined);

const selectUser = (selectedUserId: string) => (userId.value = selectedUserId);

const confirm = async () => {
  try {
    // TODO: Add validation errors here as well
    if (!props.teamId) return;
    if (!userId.value && !newUserEmail.value) return;

    if (newUserEmail.value.length) {
      await inviteNewUser({ teamId: props.teamId, email: newUserEmail.value });
      setActiveToast({
        type: ToastType.SUCCESS,
        message: "User invited.",
      });
      close();
      return;
    }

    if (!userId.value) return;

    await inviteExistingUserToTeam({ teamId: props.teamId, userId: userId.value });

    mixpanelWrapper.client.track("Invited new user")

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
    const users = await getAllUsers();
    usersList.value = users.filter(user => user.id !== activeUser?.value?.id) ?? [];
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
