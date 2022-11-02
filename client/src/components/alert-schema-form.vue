<script lang="ts" setup>
/**
 * 1. Take in object as prop that resembles an alert received for
 *    the current application.
 * 2. Recursively turn object into dom elements where object keys are buttons
 *    and values are spans.
 * 3. Define an emitter for the key buttons to call when clicked that will send
 *    their path to the root up as a concatenated string
 */
import { ref } from "vue";
import AlertSchemaTreeBuilder from "./alert-schema-tree-builder.vue";

defineProps<{ baseObject: Record<any, any> }>();

const currentStep = ref(0);
const alertSchemaForm = ref({
  title: "",
  description: "",
  link: "",
});

const sendSchemaObject = (key: string) => console.log(key);
</script>

<template>
  <div>
    <alert-schema-tree-builder
      root-key=""
      :schema-object="baseObject"
      @on-element-select="sendSchemaObject"
    />
  </div>
</template>
