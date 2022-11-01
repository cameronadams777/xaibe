<script lang="ts" setup>
defineProps<{ label: string; value: string | number | undefined }>();
</script>

<template>
  <div class="inline-block">
    <button class="text-md font-bold">{{ label }}</button>
    <span class="text-sm">: {{ value }}</span>
  </div>
</template>
