<template>
  <div
    class="overflow-hidden transition-all duration-500"
    :class="{ closed: !isOpen, open: isOpen }"
  >
    <button
      class="w-full bg-white border-1 shadow-sm text-left flex justify-between items-center rounded-md"
      @click="isOpen = !isOpen"
    >
      {{ title }}
      <ChevronUpIcon
        class="w-4 h-4 transition-all duration-500"
        :class="{ 'rotate-180': isOpen }"
      />
    </button>
    <AlertSchemaTreeBuilder
      indent
      :root-key="title"
      :schema-object="schemaObject"
      @on-element-select="choose"
    />
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { ChevronUpIcon } from "@heroicons/vue/24/solid";
import AlertSchemaTreeBuilder from "./AlertSchemaTreeBuilder.vue";

defineProps<{
  rootKey: string;
  title: string;
  schemaObject: Record<any, any>;
}>();

const emits = defineEmits<{
  (event: "onElementSelect", newKey: string): void;
}>();

const isOpen = ref(true);

const choose = (newKey: string) => emits("onElementSelect", newKey);
</script>

<style>
.closed {
  max-height: 28px;
}

.open {
  max-height: 1000px;
}
</style>
