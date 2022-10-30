<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { TrashIcon } from "@heroicons/vue/24/outline";
import { useApplicationsStore } from "../state/applications";
import { useModalStore } from "../state/modals";
import TheMainLayout from "../layouts/the-main-layout.vue";
import { IAlert, IApplication } from "../types";
import { fetchApplicationById } from "../api/applications";
import ApplicationAlerts from "../components/application-alerts.vue";
import { useAlertsStore } from "../state/alerts";

const { getCachedApplication, cacheApplication } = useApplicationsStore();
const { getCachedApplicationAlerts } = useAlertsStore();
const { setIsDeleteApplicationConfirmationModalShown } = useModalStore();

const activeApplication = ref<IApplication | undefined>(undefined);
const applicationAlerts = ref<IAlert[]>([]);
const applicationUrl = ref("");

const router = useRouter();
const route = useRoute();

const getActiveApplication = async (applicationId: number) => {
  // TODO: Refactor this into state
  const cachedApplication = getCachedApplication(applicationId);
  if (cachedApplication != null) {
    activeApplication.value = cachedApplication;
    applicationUrl.value = `http://localhost:5000/api/webhook?application_id=${applicationId}`;
    return;
  }
  try {
    const application = await fetchApplicationById({ applicationId });
    if (!application) {
      router.push("/404");
      return;
    }
    cacheApplication(application);
    activeApplication.value = application;
    applicationUrl.value = `http://localhost:5000/api/webhook?application_id=${applicationId}`;
  } catch (error) {
    console.error(
      "Galata Error: An error occurred trying to fetch the requested application:",
      error
    );
    router.push("/");
  }
};

const getApplicationAlerts = async (applicationId: number) => {
  try {
    const cachedAlerts = await getCachedApplicationAlerts({
      applicationId,
    });
    applicationAlerts.value = cachedAlerts;
  } catch (error) {
    console.error(
      "Galata Error: An error occurred fetching alerts for the specified application.",
      error
    );
    return;
  }
};

onMounted(async () => {
  const applicationId = parseInt(route.params.applicationId as string);
  // TODO: Introduce loading component for while data is being fetched and then Promise.all these requests
  await getActiveApplication(applicationId);
  await getApplicationAlerts(applicationId);
});
</script>

<template>
  <the-main-layout>
    <div class="w-full h-full p-4">
      <div class="flex justify-between items-center">
        <h2 class="capitalize">{{ activeApplication?.Name }}</h2>
        <button
          role="button"
          class="w-8 h-8 p-0 m-0 bg-white text-red-500 hover:shadow-md rounded-full border-none transition-all duration-500 cursor-pointer"
          aria-label="Delete Application"
          @click="setIsDeleteApplicationConfirmationModalShown(true)"
        >
          <trash-icon class="w-5 h-5" />
        </button>
      </div>
      <application-alerts
        v-if="activeApplication?.AlertSchema != null"
        :alerts="applicationAlerts"
      />
      <div v-else>
        <h3>Let's Get Started!</h3>
        <p>Utilize the following url to begin receiving applications:</p>
        <p>{{ applicationUrl }}</p>
        <p>
          Need more information? Checkout our docs
          <a href="https://galata.app/docs/applications/setup" target="_blank"
            >here</a
          >.
        </p>
      </div>
    </div>
  </the-main-layout>
</template>
