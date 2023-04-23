<template>
  <div class="flex flex-col mb-2" :class="class">
    <label :for="id" class="font-bold mb-2">{{ label }}</label>
    <input 
      v-bind="$attrs" 
      :id="id" 
      :name="id" 
      :type="type" 
      :value="modelValue" 
      class="p-1 mb-2" 
      @input="updateValue" 
    />
  </div>
</template>

<script lang="ts" setup>
withDefaults(
  defineProps<{
    id: string;
    modelValue: string | number;
    type?: string;
    class?: string;
    label?: string;
    error?: string;
  }>(),
  {
    label: "",
    text: "",
    modelValue: "",
    error: "",
    class: "",
  }
);

const emit = defineEmits(["update:modelValue"]);

const updateValue = (event: Event) =>
  emit("update:modelValue", (<HTMLInputElement>event.target)?.value);
</script>
