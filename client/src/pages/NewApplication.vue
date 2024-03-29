<template>
  <TheMainLayout>
    <div class="w-full flex flex-col justify-center items-center">
      <h2>Create New Application</h2>
      <ChooseApplicationStep
        v-if="applicationType == null"
        @on-continue="(type: ApplicationType) => applicationType = type"
      />
      <NewApplicationInformationStep
        v-else
        :application-type="applicationType"
        :is-submitting="isSubmitting"
        @on-create="submitForm"
      />
    </div>
  </TheMainLayout>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { storeToRefs } from "pinia";
import { createNewApplication } from "src/api/applications";
import {
  useActiveUserStore,
  useApplicationsStore,
  useToastStore,
} from "src/state";
import TheMainLayout from "src/layouts/TheMainLayout.vue";
import ChooseApplicationStep from "src/components/NewApplicationChooseApplicationStep.vue";
import NewApplicationInformationStep from "src/components/NewApplicationInformationStep.vue";
import { ToastType, ApplicationType } from "src/types";
import { getAppSchemaByType } from "src/helpers";
import { mixpanelWrapper } from "src/tools/mixpanel";

const router = useRouter();
const { cacheApplication } = useApplicationsStore();
const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);
const { setActiveToast } = useToastStore();

const applicationType = ref<ApplicationType | undefined>(undefined);
const isSubmitting = ref(false);

const submitForm = async (applicationName: string, teamId?: string) => {
  try {
    isSubmitting.value = true;
    // TODO: Move this logic to application state
    if (!activeUser?.value)
      throw new Error("Xaibe Error: User data not available.");

    if (!applicationType.value)
      throw new Error("Xaibe Error: Application type not defined.");

    const application = await createNewApplication({
      alertSchema: getAppSchemaByType(applicationType.value),
      userId: activeUser.value.id,
      applicationName,
      teamId,
    });
    if (!application)
      throw new Error("Xaibe Error: Application not generated.");

    mixpanelWrapper.client.track("New application created");
    cacheApplication(application);
    await activeUserStore.getActiveUser();
    router.push(`/applications/${application.id}`);
    isSubmitting.value = false;
  } catch (error) {
    console.log(error);
    setActiveToast({
      message:
        "An error occurred while creating your new application. Please try again later.",
      type: ToastType.ERROR,
    });
    isSubmitting.value = false;
  }
};
</script>
