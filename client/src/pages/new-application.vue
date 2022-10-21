<script lang="ts" setup>
import { ref } from "vue";
import TheMainLayout from "../layouts/the-main-layout.vue";
import { useActiveUserStore } from "../state/active-user";
import { createNewApplication } from "../api/applications";
import { storeToRefs } from "pinia";

const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);

const applicationName = ref("");
// const teamId = ref("");
const isSubmitting = ref(false);

const submitForm = async () => {
  if (!activeUser?.value) {
    // TODO: Handle with toast message here
    console.error(
      "Galata Error: An unknown error occurred. Please try again later."
    );
    return;
  }
  isSubmitting.value = true;
  await createNewApplication({
    applicationName: applicationName.value,
    teamId: undefined,
    userId: activeUser.value.id,
  });
  isSubmitting.value = false;
};
</script>

<template>
  <the-main-layout>
    <div class="w-full h-full flex flex-col justify-center items-center">
      <h2>Create a New Application</h2>
      <div class="flex flex-col w-1/4 mb-2">
        <label for="applicationName" class="font-bold mb-2"
          >Application Name</label
        >
        <input
          v-model="applicationName"
          id="applicationName"
          name="applicationName"
          type="text"
          class="p-1"
        />
      </div>
      {{ /* TODO: Add a select field here to select from teams you're apart of */}}
      <button
        class="w-1/4 mb-2 p-2 text-white font-bold bg-indigo-600 hover:bg-indigo-800 disabled:opacity-50 rounded-md border-none cursor-pointer"
        :disabled="isSubmitting"
        :aria-disabled="isSubmitting"
        @click="submitForm"
      >
        Create
      </button>
    </div>
  </the-main-layout>
</template>
