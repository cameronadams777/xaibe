<template>
  <div
    class="w-full h-full p-0 border-gray-400 border-1 rounded-md overflow-hidden"
  >
    <button
      v-for="(user, index) of users"
      :class="{
        'w-full m-0 p-2 border-0 transition-colors duration-300 cursor-pointer': true,
        'border-b-1 border-gray-400': index !== users.length && users.length > 1,
        'bg-white hover:bg-gray-600 hover:text-white':
          user.id !== selectedUserId,
        'bg-gray-600 text-white': user.id === selectedUserId,
      }"
      @click="emits('onSelect', user.id)"
    >
      {{ user.firstName }} {{ user.lastName }}
    </button>
  </div>
</template>

<script lang="ts" setup>
import { User } from "src/types";

defineProps<{
  selectedUserId?: string;
  users: User[];
}>();

const emits = defineEmits<{
  (event: "onSelect", userId: string): void;
}>();
</script>
