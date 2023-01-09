<script lang="ts" setup>
/**
 * 1. Take in object as prop that resembles an alert received for
 *    the current application.
 * 2. Recursively turn object into dom elements where object keys are buttons
 *    and values are spans.
 * 3. Define an emitter for the key buttons to call when clicked that will send
 *    their path to the root up as a concatenated string
 */
import { computed, reactive, ref } from "vue";
import AlertSchemaTreeBuilder from "./alert-schema-tree-builder.vue";
import { useApplicationsStore, useToastStore } from "src/state";
import { ToastType } from "src/types";
import { mixpanelWrapper } from "src/tools/mixpanel";

const props = defineProps<{
  applicationId: number;
  baseObject: Record<any, any>;
}>();

const { addSchemaToApplication } = useApplicationsStore();
const { setActiveToast } = useToastStore();

const currentStep = ref(0);
const alertSchemaForm = reactive({
  title: "",
  description: "",
  link: "",
});

const formTitle = computed(() => {
  const titles: Record<number, string> = {
    0: "Let's Choose a field for the Alert Title",
    1: "Now, Choose a field for the Description",
    2: "Finally, choose the field that links to your alert",
    3: "Submitting",
  };

  return titles[currentStep.value];
});

const goBack = () => {
  currentStep.value = currentStep.value - 1;
};

const submitForm = async () => {
  try {
    await addSchemaToApplication({
      applicationId: props.applicationId,
      ...alertSchemaForm,
    });
    mixpanelWrapper.client.track("Alert schema submitted");
    // TODO: Refactor code on application details page to utilize
    // the cached applications in state in order to not require this
    window.location.reload();
  } catch (error) {
    console.error(error);
    setActiveToast({
      type: ToastType.ERROR,
      message:
        "An error occurred while trying to submit the schema for your alerts. Please try again.",
    });
  }
};

const sendSchemaObject = async (key: string) => {
  if (currentStep.value === 0) {
    alertSchemaForm.title = key;
  } else if (currentStep.value === 1) {
    alertSchemaForm.description = key;
  } else if (currentStep.value === 2) {
    alertSchemaForm.link = key;
    await submitForm();
  }
  currentStep.value = currentStep.value + 1;
};
</script>

<template>
  <div>
    <div class="flex justify-between items-center">
      <h3>{{ formTitle }}</h3>
      <base-button v-if="currentStep !== 0" text="Back" @click="goBack" />
    </div>
    <alert-schema-tree-builder
      root-key=""
      :schema-object="baseObject"
      @on-element-select="sendSchemaObject"
    />
  </div>
</template>
