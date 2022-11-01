<script lang="ts" setup>
import AlertSchemaKeyElement from "./alert-schema-key-element.vue";
import AlertSchemaObjectAccordion from "./alert-schema-object-accordion.vue";

defineProps<{ schemaObject: Record<any, any>; indent?: boolean }>();
</script>

<template>
  <div
    v-for="(key, index) of Object.keys(schemaObject)"
    :key="index"
    :class="{ 'pl-8': indent }"
  >
    <alert-schema-key-element
      v-if="typeof schemaObject[key] !== 'object' || schemaObject[key] == null"
      :label="key"
      :value="schemaObject[key]"
    />
    <alert-schema-object-accordion
      v-else
      :title="key"
      :schema-object="schemaObject[key]"
    />
  </div>
</template>
