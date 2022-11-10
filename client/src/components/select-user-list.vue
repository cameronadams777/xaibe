<template>
  <div
    class="w-full h-full border-gray-400 border-1 rounded-md overflow-hidden"
  >
    <div
      class="p-2 border-b-1 border-gray-400 hover:bg-gray-600 hover:text-white transition-colors duration-300"
    >
      Cameron Adams
    </div>
  </div>
</template>

<script lang="ts" setup>
import { IUser } from "src/types";

defineProps<{ users: IUser[] }>();
</script>
