<script lang="ts" setup>
import { useRoute, useRouter } from "vue-router";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import BaseModal from "./base-modal.vue";
import { useModalStore } from "../state/modals";
import { useApplicationsStore } from "../state/applications";

defineProps<{ isOpen: boolean }>();

const router = useRouter();
const route = useRoute();
const { setIsDeleteApplicationConfirmationModalShown } = useModalStore();
const { deleteApplication } = useApplicationsStore();

const attemptToDeleteApplication = async () => {
  try {
    const applicationId = parseInt(route.params.applicationId as string);
    await deleteApplication(applicationId);
    setIsDeleteApplicationConfirmationModalShown(false);
    // TODO: Display a successful toast message
    router.push("/");
  } catch (error) {
    // TODO: Display a error toast message
    setIsDeleteApplicationConfirmationModalShown(false);
  }
};
</script>

<template>
  <base-modal :is-open="isOpen">
    <div
      class="relative w-full h-full flex flex-col justify-center items-center"
    >
      <button
        class="absolute top-2 right-2 w-8 h-8 text-black hover:text-gray-200 bg-transparent border-none cursor-pointer transition-all duration-500"
        @click="setIsDeleteApplicationConfirmationModalShown(false)"
      >
        <x-mark-icon />
      </button>
      <h2 class="text-lg">Are you sure you want to delete this application?</h2>
      <div class="flex">
        <button
          class="px-4 py-2 bg-red-500 hover:bg-red-800 text-white border-none rounded-lg cursor-pointer transition-all duration-500"
          @click="attemptToDeleteApplication"
        >
          Delete
        </button>
        <button
          class="px-4 py-2 bg-white hover:bg-gray-400 hover:text-white border-1 border-gray-300 rounded-lg cursor-pointer transition-all duration-500"
          @click="setIsDeleteApplicationConfirmationModalShown(false)"
        >
          Cancel
        </button>
      </div>
    </div>
  </base-modal>
</template>
