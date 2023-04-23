<script lang="ts" setup>
import AlertSchemaKeyElement from "./alert-schema-key-element.vue";
import AlertSchemaObjectAccordion from "./alert-schema-object-accordion.vue";

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

<template>
  <div
    v-for="(key, index) of Object.keys(schemaObject)"
    :key="index"
    :class="{ 'pl-8': indent }"
  >
    <alert-schema-key-element
      v-if="typeof schemaObject[key] !== 'object' || schemaObject[key] == null"
      :root-key="rootKey"
      :label="key"
      :value="schemaObject[key]"
      @on-element-select="choose"
    />
    <alert-schema-object-accordion
      v-else
      :root-key="rootKey"
      :title="key"
      :schema-object="schemaObject[key]"
      @on-element-select="choose"
    />
  </div>
</template>
