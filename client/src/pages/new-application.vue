<template>
  <the-main-layout>
    <div class="w-full flex flex-col justify-center items-center">
      <h2>Create New Application</h2>
      <choose-application-step
        v-if="applicationType == null"
        @on-continue="(type: ApplicationType) => applicationType = type"
      />
      <other-application-type-step
        v-else
        :application-type="applicationType"
        :is-submitting="isSubmitting"
        @on-create="submitForm"
      />
    </div>
  </the-main-layout>
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
import TheMainLayout from "src/layouts/the-main-layout.vue";
import ChooseApplicationStep from "src/components/new-application-choose-application-step.vue";
import OtherApplicationTypeStep from "src/components/new-application-other-application-type-step.vue";
import { ToastType, ApplicationType } from "src/types";
import { getAppSchemaByType } from "src/helpers";
import { mixpanelWrapper } from "src/tools/mixpanel";

// TODO: Allow users to create team application as well

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
      throw new Error("Galata Error: User data not available.");

    if (!applicationType.value)
      throw new Error("Galata Error: Application type not defined.");

    const application = await createNewApplication({
      alertSchema: getAppSchemaByType(applicationType.value),
      userId: activeUser.value.id,
      applicationName,
      teamId,
    });
    if (!application)
      throw new Error("Galata Error: Application not generated.");

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
