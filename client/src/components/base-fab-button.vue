<script lang="ts" setup>
import { useAttrs } from "vue";

const attrs = useAttrs();
</script>

<template>
  <button
    class="w-8 h-8 mr-2 bg-white flex justify-center items-center text-black focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 border-none rounded-full cursor-pointer hover:shadow-md transition-all duration-500"
    :class="attrs.class"
  >
    <slot />
  </button>
</template>
