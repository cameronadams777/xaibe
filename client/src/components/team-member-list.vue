<template>
  <div class="w-full h-full flex flex-col">
    <h5 class="p-0 m-0 ml-1 mb-2">Members</h5>
    <div
      class="relative w-full h-full flex flex-col mr-4 border-1 border-gray-300 rounded-lg overflow-y-auto"
    >
      <div
        v-if="members.length"
        v-for="member in members"
        :key="member.ID"
        class="p-4 flex justify-between items-center border-b border-gray-300 font-bold text-black no-underline transition-all duration-300"
      >
        {{ member.FirstName }} {{ member.LastName }}
        <button
          v-if="managerControlsVisible && activeUser?.ID !== member.ID"
          class="p-0 hover:text-red-500 rounded-full border-none bg-white flex justify-center items-center cursor-pointer"
          @click="displayRemoveUserConfirmationModal(member.ID)"
        >
          <x-mark-icon class="h-full w-4" />
        </button>
      </div>
      <div v-else class="w-full h-full flex justify-center items-center">
        <p>No members</p>
      </div>
      <base-fab-button
        v-if="managerControlsVisible"
        class="absolute bottom-2 right-2 bg-indigo-500 text-white"
      >
        <plus-icon class="h-full w-4" />
      </base-fab-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { XMarkIcon, PlusIcon } from "@heroicons/vue/24/solid";
import { storeToRefs } from "pinia";
import { useActiveUserStore, useModalStore } from "src/state";
import { IUser } from "src/types";

const props = defineProps<{
  teamId: number;
  members: IUser[];
  managerControlsVisible: boolean;
}>();

const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);

const { setRemoveUserConfirmationProps } = useModalStore();

const displayRemoveUserConfirmationModal = (userId: number) =>
  setRemoveUserConfirmationProps({
    teamId: props.teamId,
    userId,
    isOpen: true,
  });
</script>
