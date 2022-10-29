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

const { getCachedApplication, cacheApplication } = useApplicationsStore();
const { setIsDeleteApplicationConfirmationModalShown } = useModalStore();

const activeApplication = ref<IApplication | undefined>(undefined);
const applicationAlerts = ref<IAlert[]>([]);
const applicationUrl = ref("");

const router = useRouter();
const route = useRoute();

onMounted(async () => {
  const applicationId = parseInt(route.params.applicationId as string);
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
      "Galata Error: An error occurred trying to fetch the application you wanted:",
      error
    );
    router.push("/");
  }
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
