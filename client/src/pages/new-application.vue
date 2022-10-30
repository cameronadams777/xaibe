<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { storeToRefs } from "pinia";
import { createNewApplication } from "../api/applications";
import { useActiveUserStore, useApplicationsStore } from "../state";
import TheMainLayout from "../layouts/the-main-layout.vue";
import BaseButton from "../components/base-button.vue";
import { ButtonTextSize } from "../types";

// TODO: Allow users to create team application as well

const router = useRouter();
const { cacheApplication } = useApplicationsStore();
const { getActiveUser, ...activeUserStore } = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);

const applicationName = ref("");
const isSubmitting = ref(false);

const submitForm = async () => {
  try {
    isSubmitting.value = true;
    if (!activeUser?.value)
      throw new Error("Galata Error: User data not available.");
    // TODO: Move this logic to application state
    const application = await createNewApplication({
      userId: activeUser.value.ID,
      applicationName: applicationName.value,
    });
    if (!application)
      throw new Error("Galata Error: Application not generated.");
    cacheApplication(application);
    await getActiveUser();
    router.push(`/applications/${application.ID}`);
    isSubmitting.value = false;
  } catch (error) {
    console.error(error);
    // TODO: Add toast message for better UX
    isSubmitting.value = false;
  }
};
</script>

<template>
  <the-main-layout>
    <div class="w-full flex flex-col justify-center items-center">
      <h2>Create new Application</h2>
      <div class="flex flex-col w-1/4 mb-3">
        <label for="applicationName" class="font-bold mb-2"
          >Application Name</label
        >
        <input
          v-model="applicationName"
          id="applicationName"
          name="applicationName"
          type="applicationName"
          placeholder="Airbrake"
          class="p-1.5"
        />
      </div>
      <base-button
        text="Create"
        :text-size="ButtonTextSize.LARGE"
        :show-spinner="isSubmitting"
        :disabled="isSubmitting"
        :aria-disabled="isSubmitting"
        @click="submitForm"
      />
    </div>
  </the-main-layout>
</template>
