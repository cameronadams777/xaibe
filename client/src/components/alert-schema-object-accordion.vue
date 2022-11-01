<script lang="ts" setup>
import { ref } from "vue";
import { ChevronUpIcon } from "@heroicons/vue/24/solid";
import AlertSchemaTreeBuilder from "./alert-schema-tree-builder.vue";

defineProps<{ title: string; schemaObject: Record<any, any> }>();

const isOpen = ref(true);
</script>

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
      <chevron-up-icon
        class="w-4 h-4 transition-all duration-500"
        :class="{ 'rotate-180': isOpen }"
      />
    </button>
    <alert-schema-tree-builder :schema-object="schemaObject" indent />
  </div>
</template>

<style>
.closed {
  max-height: 28px;
}

.open {
  max-height: 1000px;
}
</style>
