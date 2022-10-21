<script lang="ts" setup>
import { useRouter } from "vue-router";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import BaseModal from "./base-modal.vue";
import { useModalStore } from "../state/modals";

defineProps<{ isOpen: boolean }>();
const router = useRouter();
const { setIsNewElementModalShown } = useModalStore();

const navigateToNewAppPage = () => {
  router.push("/applications/new");
  setIsNewElementModalShown(false);
};

const navigateToNewTeamPage = () => {
  router.push("/teams/new");
  setIsNewElementModalShown(false);
};
</script>

<template>
  <base-modal :is-open="isOpen">
    <div class="relative h-full flex flex-col justify-center items-center">
      <button
        class="absolute top-2 right-2 w-8 h-8 text-black hover:text-gray-200 bg-transparent border-none cursor-pointer"
        @click="setIsNewElementModalShown(false)"
      >
        <x-mark-icon />
      </button>
      <button
        class="w-1/2 mb-2 p-2 bg-white text-black no-underline border-gray-300 border-1 rounded-md text-center hover:bg-gray-600 hover:text-white transition-colors duration-500 cursor-pointer"
        @click="navigateToNewAppPage"
      >
        Create New Application
      </button>
      <button
        class="w-1/2 mb-2 p-2 bg-white text-black no-underline border-gray-300 border-1 rounded-md text-center hover:bg-gray-600 hover:text-white transition-colors duration-500 cursor-pointer"
        @click="navigateToNewTeamPage"
      >
        Create New Team
      </button>
    </div>
  </base-modal>
</template>
