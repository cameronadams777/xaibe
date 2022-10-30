<script lang="ts" setup>
import { ButtonTextSize, ButtonVariant } from "../types";

const { variant = ButtonVariant.PRIMARY, textSize = ButtonTextSize.MEDIUM } =
  defineProps<{
    text: string;
    showSpinner?: boolean;
    textSize?: ButtonTextSize;
    variant?: ButtonVariant;
  }>();
</script>

<template>
  <button
    class="mb-2 p-2 flex justify-center items-center font-bold disabled:opacity-50 rounded-md border-none cursor-pointer transition-all duration-500"
    :class="{
      'bg-indigo-600 hover:bg-indigo-800 text-white':
        variant === ButtonVariant.PRIMARY,
      'bg-red-500 hover:bg-red-800 text-white':
        variant === ButtonVariant.DANGER,
      'bg-white hover:bg-gray-400 hover:text-white':
        variant === ButtonVariant.WHITE,
    }"
  >
    <span
      :class="{
        'text-sm': textSize === ButtonTextSize.SMALL,
        'text-md': textSize === ButtonTextSize.MEDIUM,
        'text-lg': textSize === ButtonTextSize.LARGE,
        'mr-2': showSpinner,
      }"
      >{{ text }}</span
    >
    <the-spinner :show="showSpinner" />
  </button>
</template>
