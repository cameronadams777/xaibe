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
      <input v-model="email" id="email" name="email" type="email" class="p-1" />
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
    <div class="mb-4 flex items-center">
      <input
        v-model="allowTelemetry"
        id="allowTelemetry"
        name="allowTelemetry"
        type="checkbox"
        class="mr-2"
      />
      <label for="allowTelemetry">Allow Anonymous Telemetry</label>
    </div>
    <base-button
      text="Register"
      class="w-1/4"
      :disabled="isSubmitting"
      :aria-disabled="isSubmitting"
      :show-spinner="isSubmitting"
      @click="submitForm"
    />
    <router-link
      to="/login"
      class="text-indigo-600 no-underline hover:underline cursor-pointer"
      >Already have an account?</router-link
    >
  </div>
</template>


<script lang="ts" setup>
import { useRouter } from "vue-router";
import { ref } from "vue";
import { ToastType } from "../types";
import { useAuthStore, useToastStore } from "../state";
import { mixpanelWrapper } from "src/tools/mixpanel";

const router = useRouter();
const { setActiveToast } = useToastStore();
const { registerUser } = useAuthStore();

const firstName = ref("");
const lastName = ref("");
const email = ref("");
const password = ref("");
const confirmPassword = ref("");
const allowTelemetry = ref(true);
const isSubmitting = ref(false);

const submitForm = async () => {
  try {
    isSubmitting.value = true;
    localStorage.setItem("allowTelemetry", allowTelemetry.toString());
    await registerUser({
      firstName: firstName.value,
      lastName: lastName.value,
      email: email.value,
      password: password.value,
      passwordConfirmation: confirmPassword.value,
    });
    mixpanelWrapper.client.track("Register user");
    isSubmitting.value = false;
    router.push("/");
  } catch (error) {
    console.log(error);
    isSubmitting.value = false;
    setActiveToast({
      message:
        "An error occurred while trying to create your account. Please try again later.",
      type: ToastType.ERROR,
    });
  }
};
</script>

