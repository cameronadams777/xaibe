<script lang="ts" setup>
import TheSpinner from "./the-spinner.vue";
import { ButtonTextSize, ButtonVariant } from "src/types";
import { useAttrs } from "vue";

withDefaults(
  defineProps<{
    text: string;
    showSpinner?: boolean;
    textSize?: ButtonTextSize;
    variant?: ButtonVariant;
  }>(),
  {
    showSpinner: false,
    textSize: ButtonTextSize.MEDIUM,
    variant: ButtonVariant.PRIMARY,
  }
);

const attrs = useAttrs();
</script>

<template>
  <button
    role="button"
    class="mb-2 p-2 flex justify-center items-center font-bold disabled:opacity-50 rounded-md cursor-pointer transition-all duration-500"
    :class="[
      {
        'bg-indigo-600 hover:bg-indigo-800 text-white border-none':
          variant === ButtonVariant.PRIMARY,
        'bg-red-500 hover:bg-red-800 text-white border-none':
          variant === ButtonVariant.DANGER,
        'bg-white border-1 border-gray-400 hover:bg-indigo-600 hover:border-indigo-600 hover:text-white':
          variant === ButtonVariant.WHITE,
      },
      attrs.class,
    ]"
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
