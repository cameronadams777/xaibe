<template>
  <button
    role="button"
    :class="
      twMerge(
        'w-8 h-8 mr-2 bg-transparent flex justify-center items-center text-black focus:outline-none focus:ring-2 focus:ring-indigo-600 focus:ring-offset-2 border-none rounded-full cursor-pointer hover:shadow-md transition-all duration-500',
        attrs.class as string
      )
    "
  >
    <slot />
  </button>
</template>

<script lang="ts" setup>
import { useAttrs } from "vue";
import { twMerge } from "tailwind-merge";

const attrs = useAttrs();
</script>

