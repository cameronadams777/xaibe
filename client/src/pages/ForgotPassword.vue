<template>
  <div class="w-full h-full flex flex-col justify-center items-center">
    <div class="w-11/12 md:w-1/2 lg:w-1/4 text-center">
      <h2>Forget something?</h2>
      <div class="flex flex-col text-left mb-4">
        <label for="email" class="font-bold mb-2">Email</label>
        <input
          v-model="formData.email"
          id="email"
          name="email"
          type="email"
          class="p-1"
        />
      </div>
      <base-button
        text="Reset Password"
        :disabled="isSubmitting"
        :aria-disabled="isSubmitting"
        :show-spinner="isSubmitting"
        class="w-full"
        @click="submitForm"
      />
      <div class="w-full text-center">
        <router-link
          to="/login"
          class="text-indigo-600 no-underline hover:underline cursor-pointer"
          >Return to Login</router-link
        >
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { submitResetPasswordRequest } from "src/api/auth";
import { useToastStore } from "src/state";
import { ToastType } from "src/types";
import { reactive, ref } from "vue";

const { setActiveToast } = useToastStore();

const isSubmitting = ref(false);
const emailSent = ref(false);
const formData = reactive({
  email: "",
});

const submitForm = async () => {
  try {
    isSubmitting.value = true;
    await submitResetPasswordRequest({ email: formData.email });
    emailSent.value = true;
    isSubmitting.value = false;
  } catch (error) {
    isSubmitting.value = false;
    setActiveToast({
      message:
        "An error occurred while trying to create your account. Please try again later.",
      type: ToastType.ERROR,
    });
  }
};
</script>
