<template>
  <div class="inline-block">
    <button class="text-md font-bold" @click="choose">
      {{ label }}
    </button>
    <span class="text-sm">: {{ value }}</span>
  </div>
</template>


<script lang="ts" setup>
  const props = defineProps<{
    rootKey: string;
    label: string;
    value: string | number | undefined;
  }>();
  const emits = defineEmits<{ (event: "onElementSelect", key: string): void }>();
  const choose = () => emits("onElementSelect", props.label);
</script>

