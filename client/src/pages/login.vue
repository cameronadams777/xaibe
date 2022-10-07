<script lang="ts" setup>
import { useRouter } from "vue-router";
import { ref } from "vue";
import axios from "axios";

const router = useRouter();

const email = ref("");
const password = ref("");

const submitForm = async () => {
  try {
    const response = await axios
      .post("http://localhost:5000/api/login", {
        email: email.value,
        password: password.value,
      })
      .then((res) => res.data);

    localStorage.setItem("token", response.data.token);
    router.push("/");
  } catch (error) {
    console.error(error);
  }
};
</script>

<template>
  <div class="w-full h-full flex flex-col justify-center items-center">
    <h2>Welcome Back!</h2>
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
        type="password"
        class="p-1.5"
      />
    </div>
    <div class="w-1/4 flex justify-end mb-3">
      <router-link
        to="/forgot-password"
        class="font-bold bg-transparent text-indigo-600 no-underline hover:underline border-none cursor-pointer"
      >
        Forgot Password?
      </router-link>
    </div>
    <button
      class="w-1/4 mb-2 p-2 text-white font-bold bg-indigo-600 hover:bg-indigo-800 rounded-md border-none cursor-pointer"
      @click="submitForm"
    >
      Log In
    </button>
    <router-link
      to="/register"
      class="text-indigo-600 no-underline hover:underline cursor-pointer"
      >Don't have an account yet?</router-link
    >
  </div>
</template>
