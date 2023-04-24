<template>
  <div
    class="p-4 flex justify-between items-center border-b border-gray-300 font-bold text-black no-underline transition-all duration-300"
  >
    <p>{{ item.title }}</p>
    <p>{{ item.description }}</p>
  </div>
</template>

<script lang="ts" setup>
interface IListItem {
  title: string;
  description: string;
}
defineProps<{ item: IListItem }>();
</script>
