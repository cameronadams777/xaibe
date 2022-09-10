<script lang="ts" setup>
import { ref } from "vue";

const email = ref("");
const password = ref("");
const confirmPassword = ref("");
</script>

<template>
  <div class="w-full h-full flex flex-col justify-center items-center">
    <h2>Nice to meet you!</h2>
    <div class="flex flex-col w-1/4 mb-2">
      <label for="email" class="font-bold mb-2">Email</label>
      <input v-model="email" id="email" name="email" type="text" class="p-1" />
    </div>
    <div class="flex flex-col w-1/4 mb-3">
      <label for="password" class="font-bold mb-2">Password</label>
      <input
        v-model="password"
        id="password"
        name="password"
        type="text"
        class="p-1.5"
      />
    </div>
    <div class="flex flex-col w-1/4 mb-6">
      <label for="confirmPassword" class="font-bold mb-2"
        >Confirm Password</label
      >
      <input
        v-model="confirmPassword"
        id="confirmPassword"
        name="confirmPassword"
        type="text"
        class="p-1.5"
      />
    </div>
    <button
      class="w-1/4 p-2 text-white font-bold bg-indigo-600 hover:bg-indigo-800 rounded-md border-none mb-2"
    >
      Register
    </button>
    <router-link
      to="/login"
      class="text-indigo-600 no-underline hover:underline"
      >Already have an account?</router-link
    >
  </div>
</template>
