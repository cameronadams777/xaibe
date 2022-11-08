<template>
  <div class="flex flex-col mb-2" :class="class">
    <label :for="id" class="font-bold mb-2">{{ label }}</label>
    <input v-bind="$attrs" :id="id" :name="id" class="p-1 mb-2" />
  </div>
</template>

<script lang="ts" setup>
withDefaults(
  defineProps<{
    id: string;
    class?: string;
    label?: string;
    error?: string;
  }>(),
  {
    label: "",
    error: "",
    class: "",
  }
);

const emit = defineEmits(["update:value"]);

const updateValue = (event: Event) =>
  emit("update:value", (<HTMLInputElement>event.target)?.value);
</script>
