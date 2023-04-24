<template>
  <div class="w-full h-full flex flex-col justify-center items-center">
    <h2>Welcome Back!</h2>
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
    <div class="w-1/4 flex justify-end mb-3">
      <router-link
        to="/forgot-password"
        class="font-bold bg-transparent text-indigo-600 no-underline hover:underline border-none cursor-pointer"
      >
        Forgot Password?
      </router-link>
    </div>
    <BaseButton
      text="Log In"
      class="w-1/4"
      :disabled="isSubmitting"
      :aria-disabled="isSubmitting"
      :show-spinner="isSubmitting"
      @click="submitForm"
    />
    <router-link
      to="/register"
      class="text-indigo-600 no-underline hover:underline cursor-pointer"
      >Don't have an account yet?</router-link
    >
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router";
import { ref } from "vue";
import { useAuthStore, useToastStore } from "../state";
import { ToastType } from "../types";
import { mixpanelWrapper } from "src/tools/mixpanel";

const router = useRouter();
const { setActiveToast } = useToastStore();
const { login } = useAuthStore();

const email = ref("");
const password = ref("");
const isSubmitting = ref(false);

const submitForm = async () => {
  try {
    isSubmitting.value = true;
    await login({ email: email.value, password: password.value });
    mixpanelWrapper.client.track("User login");
    isSubmitting.value = false;
    router.push("/");
  } catch (error) {
    setActiveToast({
      message: "An error occurred while attempting to log in.",
      type: ToastType.ERROR,
    });
    isSubmitting.value = false;
  }
};
</script>

