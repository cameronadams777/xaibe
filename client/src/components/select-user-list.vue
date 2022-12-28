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
          user.ID !== selectedUserId,
        'bg-gray-600 text-white': user.ID === selectedUserId,
      }"
      @click="emits('onSelect', user.ID)"
    >
      {{ user.FirstName }} {{ user.LastName }}
    </button>
  </div>
</template>

<script lang="ts" setup>
import { IUser } from "src/types";

defineProps<{
  selectedUserId?: number;
  users: IUser[];
}>();

const emits = defineEmits<{
  (event: "onSelect", userId: number): void;
}>();
</script>
