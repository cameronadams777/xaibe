<template>
  <h2 class="text-center text-lg">
    Are you sure you want to change your number of seats?
  </h2>
  <div class="w-full flex flex-col md:flex-row">
    <BaseButton
      text="Confirm"
      :variant="ButtonVariant.PRIMARY"
      class="w-full mr-2"
      @click="confirm"
    />
    <BaseButton
      text="Cancel"
      :variant="ButtonVariant.WHITE"
      class="w-full"
      @click="close"
    />
  </div>
</template>

<script setup lang="ts">
import { ButtonVariant } from "src/types";

const emits = defineEmits<{
  (event: "onClose"): void;
  (event: "onConfirm"): void;
}>();

const confirm = () => emits("onConfirm");
const close = () => emits("onClose");
</script>
