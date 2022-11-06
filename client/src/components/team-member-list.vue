<template>
  <div class="w-full h-full flex flex-col">
    <h5 class="p-0 m-0 ml-1 mb-2">Members</h5>
    <div
      class="w-full h-full flex flex-col mr-4 border-1 border-gray-300 rounded-lg overflow-y-auto"
    >
      <div
        v-if="members.length"
        v-for="member in members"
        :key="member.ID"
        class="p-4 flex justify-between items-center border-b border-gray-300 hover:bg-gray-600 hover:border-gray-600 hover:text-white font-bold text-black no-underline transition-all duration-300"
      >
        {{ member.FirstName }} {{ member.LastName }}
        <base-fab-button>
          <x-mark-icon class="h-full w-4" />
        </base-fab-button>
      </div>
      <div v-else class="w-full h-full flex justify-center items-center">
        <p>No members</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { XMarkIcon } from "@heroicons/vue/24/solid";
import { IUser } from "src/types";

defineProps<{ members: IUser[] }>();
</script>
