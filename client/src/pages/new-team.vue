<script lang="ts" setup>
import { ref } from "vue";
import TheMainLayout from "../layouts/the-main-layout.vue";

const teamName = ref("");

const submitForm = () => {};
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
        class="w-1/4 mb-2 p-2 text-white font-bold bg-indigo-600 hover:bg-indigo-800 rounded-md border-none cursor-pointer"
        @click="submitForm"
      >
        Create
      </button>
    </div>
  </the-main-layout>
</template>
