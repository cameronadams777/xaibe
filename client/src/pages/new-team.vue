<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { createNewTeam } from "../api/teams";
import TheMainLayout from "../layouts/the-main-layout.vue";

const router = useRouter();

const teamName = ref("");
const isSubmitting = ref(false);

// Ensure that the user is allowed to create teams
// onMounted(() => {})

const submitForm = async () => {
  try {
    isSubmitting.value = true;
    const team = await createNewTeam({ teamName: teamName.value });
    if (!team) throw new Error("Galata Error: Team not generated.");
    isSubmitting.value = false;
    router.push(`/teams/${team.ID}`);
  } catch (error) {
    // TODO: Add toast message for better UX
    console.error(
      "Galata Error: An error occurred while attempting to create a new team:",
      error
    );
    isSubmitting.value = false;
  }
};
</script>

<template>
  <the-main-layout>
    <div class="w-full h-full flex flex-col justify-center items-center">
      <h2>Create a New Team</h2>
      <div class="flex flex-col w-1/4 mb-2">
        <label for="teamName" class="font-bold mb-2">Team Name</label>
        <input
          v-model="teamName"
          id="teamName"
          name="teamName"
          type="text"
          class="p-1"
        />
      </div>
      <button
        class="w-1/4 mb-2 p-2 text-lg text-white font-bold bg-indigo-600 hover:bg-indigo-800 rounded-md border-none cursor-pointer"
        @click="submitForm"
      >
        Create
      </button>
    </div>
  </the-main-layout>
</template>
