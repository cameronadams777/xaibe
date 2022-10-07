<script lang="ts" setup>
import { useRouter } from "vue-router";
import axios from "axios";
import { ref } from "vue";

const router = useRouter();

const firstName = ref("");
const lastName = ref("");
const email = ref("");
const password = ref("");
const confirmPassword = ref("");

const submitForm = async () => {
  try {
    const response = await axios
      .post("http://localhost:5000/api/register", {
        firstName: firstName.value,
        lastName: lastName.value,
        email: email.value,
        password: password.value,
        passwordConfirmation: confirmPassword.value,
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
    <h2>Nice to meet you!</h2>
    <div class="flex flex-col w-1/4 mb-2">
      <label for="firstName" class="font-bold mb-2">First Name</label>
      <input
        v-model="firstName"
        id="firstName"
        name="firstName"
        type="text"
        class="p-1"
      />
    </div>
    <div class="flex flex-col w-1/4 mb-2">
      <label for="lastName" class="font-bold mb-2">Last Name</label>
      <input
        v-model="lastName"
        id="lastName"
        name="lastName"
        type="text"
        class="p-1"
      />
    </div>
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
    <div class="flex flex-col w-1/4 mb-6">
      <label for="confirmPassword" class="font-bold mb-2"
        >Confirm Password</label
      >
      <input
        v-model="confirmPassword"
        id="confirmPassword"
        name="confirmPassword"
        type="password"
        class="p-1.5"
      />
    </div>
    <button
      class="w-1/4 p-2 text-white font-bold bg-indigo-600 hover:bg-indigo-800 rounded-md border-none mb-2 cursor-pointer"
      @click="submitForm"
    >
      Register
    </button>
    <router-link
      to="/login"
      class="text-indigo-600 no-underline hover:underline cursor-pointer"
      >Already have an account?</router-link
    >
  </div>
</template>
