<script lang="ts" setup>
import TheMainLayout from "../layouts/the-main-layout.vue";
import { ref } from "vue";

const applicationName = ref("");

const submitForm = () => {};
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
      <button
        class="w-1/4 mb-2 p-2 text-white font-bold bg-indigo-600 hover:bg-indigo-800 rounded-md border-none cursor-pointer"
        @click="submitForm"
      >
        Create
      </button>
    </div>
  </the-main-layout>
</template>
