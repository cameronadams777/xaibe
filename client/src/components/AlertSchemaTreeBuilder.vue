<template>
  <div
    v-for="(key, index) of Object.keys(schemaObject)"
    :key="index"
    :class="{ 'pl-8': indent }"
  >
    <AlertSchemaKeyElement
      v-if="typeof schemaObject[key] !== 'object' || schemaObject[key] == null"
      :root-key="rootKey"
      :label="key"
      :value="schemaObject[key]"
      @on-element-select="choose"
    />
    <AlertSchemaObjectAccordion
      v-else
      :root-key="rootKey"
      :title="key"
      :schema-object="schemaObject[key]"
      @on-element-select="choose"
    />
  </div>
</template>

<script lang="ts" setup>
import AlertSchemaKeyElement from "./AlertSchemaKeyElement.vue";
import AlertSchemaObjectAccordion from "./AlertSchemaObjectAccordion.vue";

const props = defineProps<{
  schemaObject: Record<any, any>;
  rootKey: string;
  indent?: boolean;
}>();

const emits = defineEmits<{
  (event: "onElementSelect", newKey: string): void;
}>();

const choose = (key: string) => {
  let newKey = props.rootKey.length ? `${props.rootKey}.${key}` : key;
  emits("onElementSelect", newKey);
};
</script>

