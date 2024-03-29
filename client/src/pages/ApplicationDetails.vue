<template>
  <TheMainLayout>
    <div class="w-full h-full p-4">
      <div class="flex justify-between items-center">
        <h2 class="capitalize">{{ activeApplication?.name }}</h2>
        <BaseFabButton
          aria-label="Delete Application"
          @click="setIsDeleteApplicationConfirmationModalShown(true)"
        >
          <TrashIcon class="w-5 h-5 text-red-500" />
        </BaseFabButton>
      </div>
      <div
        v-if="
          activeApplication != null &&
          activeApplication.alertSchemaId != null &&
          applicationAlerts.length > 0
        "
        class="w-full lg:w-1/3 h-48 lg:h-96"
      >
        <AlertsListByApplication
          :alerts="applicationAlerts"
          :alert-schema="activeApplication.alertSchema"
        />
      </div>
      <AlertSchemaForm
        v-else-if="activeApplication != null && applicationAlerts?.length"
        :application-id="activeApplication?.id"
        :base-object="applicationAlerts[0]"
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
  </TheMainLayout>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from "vue";
import { storeToRefs } from "pinia";
import { useRoute, useRouter } from "vue-router";
import { TrashIcon } from "@heroicons/vue/24/outline";
import {
  useAlertsStore,
  useApplicationsStore,
  useModalStore,
  useToastStore,
} from "src/state";
import TheMainLayout from "src/layouts/TheMainLayout.vue";
import AlertSchemaForm from "src/components/AlertSchemaForm.vue";
import AlertsListByApplication from "src/components/AlertsListByApplication.vue";
import { Application, ToastType } from "src/types";
import { fetchApplicationById } from "src/api/applications";
import { config } from "src/config";

const { getCachedApplication, cacheApplication } = useApplicationsStore();
const { getCachedApplicationAlerts, ...alertStore } = useAlertsStore();
const { localCacheAlerts } = storeToRefs(alertStore);
const { setIsDeleteApplicationConfirmationModalShown } = useModalStore();
const { setActiveToast } = useToastStore();

const activeApplication = ref<Application | undefined>(undefined);
const applicationUrl = ref("");

const applicationAlerts = computed(
  () =>
    localCacheAlerts.value?.[`application_${activeApplication.value?.id}`] ?? []
);

const router = useRouter();
const route = useRoute();

const getActiveApplication = async (applicationId: string) => {
  const cachedApplication = getCachedApplication(applicationId);
  if (cachedApplication != null) {
    activeApplication.value = cachedApplication;
    applicationUrl.value = `${config.apiBaseUrl}/api/webhook?application_id=${applicationId}`;
    return;
  }
  try {
    const application = await fetchApplicationById({ applicationId });
    console.log(application);
    if (!application) {
      router.push("/404");
      return;
    }
    cacheApplication(application);
    activeApplication.value = application;
    applicationUrl.value = `${config.apiBaseUrl}/api/webhook?application_id=${applicationId}`;
  } catch (error) {
    setActiveToast({
      message: "An error occurred trying to fetch the requested application.",
      type: ToastType.ERROR,
    });
    router.push("/");
  }
};

onMounted(async () => {
  const applicationId = route.params.applicationId as string;
  // TODO: Introduce loading component while data is being fetched and then Promise.all these requests
  await getActiveApplication(applicationId);
  await getCachedApplicationAlerts({
    applicationId,
  });
});
</script>
