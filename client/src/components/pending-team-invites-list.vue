<template>
  <base-list v-if="pendingTeamInvites.length">
    <div
      v-for="invite of pendingTeamInvites"
      :class="{
        'p-4 flex justify-between items-center font-bold text-black no-underline transition-all duration-300': true,
        'border-b border-gray-300': pendingTeamInvites.length > 1,
      }"
    >
      <div>
        <p class="font-bold m-0 p-0">{{ invite.team.name }}</p>
        <p class="text-gray-400 m-0 p-0">{{ toInviteEnum(invite.status) }}</p>
      </div>
      <div v-if="invite.status === InviteStatus.PENDING" class="flex items-center">
        <button
          class="p-1 border-0 bg-red-500 hover:bg-red-700 text-white rounded-sm flex justify-center items-center cursor-pointer transition-color duration-300"
          @click="
            updateInviteStatusAndFetchTeams({
              inviteId: invite.id,
              status: InviteStatus.REJECTED,
            })
          "
        >
          <x-mark-icon class="w-5 h-5" />
        </button>
        <button
          class="p-1 border-0 bg-green-500 hover:bg-green-700 text-white rounded-sm flex justify-center items-center cursor-pointer transition-color duration-300"
          @click="
            updateInviteStatusAndFetchTeams({
              inviteId: invite.id,
              status: InviteStatus.ACCEPTED,
            })
          "
        >
          <check-icon class="w-5 h-5" />
        </button>
      </div>
    </div>
  </base-list>
</template>

<script lang="ts" setup>
import { storeToRefs } from "pinia";
import { XMarkIcon, CheckIcon } from "@heroicons/vue/24/outline";
import { useActiveUserStore, useTeamsStore, useToastStore } from "src/state";
import { ToastType, InviteStatus } from "src/types";
import { IUpdateInviteStatusInput } from "src/api/teams";

const { getActiveUser } = useActiveUserStore();
const { getPendingTeamInvites, updateInviteStatus, ...teamsStore } =
  useTeamsStore();
const { pendingTeamInvites } = storeToRefs(teamsStore);
const { setActiveToast } = useToastStore();

const updateInviteStatusAndFetchTeams = async ({
  inviteId,
  status,
}: IUpdateInviteStatusInput) => {
  try {
    await updateInviteStatus({ inviteId, status });
    await getActiveUser();
  } catch (error) {
    console.error(error);
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred updating invite status.",
    });
  }
};

const toInviteEnum = (status: number) => {
  const statuses = {
    [InviteStatus.PENDING]: "PENDING",
    [InviteStatus.ACCEPTED]: "ACCEPTED",
    [InviteStatus.REJECTED]: "REJECTED",
  };
  return statuses[status];
};
</script>
