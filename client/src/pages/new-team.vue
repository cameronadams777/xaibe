<template>
  <the-main-layout>
    <div class="w-full h-full flex flex-col justify-center items-center">
      <h2>Create a New Team</h2>
      <NewTeamDetailsForm v-if="!hasMetadata" @onContinue="updateTeamDetails"/>
      <PaymentForm v-else/>
    </div>
  </the-main-layout>
</template>

<script lang="ts" setup>
import { computed, ref } from "vue";
import { useRouter } from "vue-router";
import { createNewTeam } from "src/api/teams";
import TheMainLayout from "src/layouts/the-main-layout.vue";
import NewTeamDetailsForm from "src/components/new-team-details-form.vue";
import PaymentForm from "src/components/payment-form.vue";
import { useToastStore } from "src/state";
import { useActiveUserStore } from "src/state/active-user";
import { ToastType, NewTeamDetailsFormSchema } from "src/types";
import { mixpanelWrapper } from "src/tools/mixpanel";

const metadata = ref<NewTeamDetailsFormSchema>();
const paymentToken = ref("");

const router = useRouter();
const { getActiveUser } = useActiveUserStore();
const { setActiveToast } = useToastStore();

const teamName = ref("");
const isSubmitting = ref(false);

const hasMetadata = computed(() => metadata.value != null);

const updateTeamDetails = (formValues: NewTeamDetailsFormSchema) => {
  metadata.value = formValues;
};

// TODO: Make product details type available for use here
const getPaymentToken = (): void => {
  // Submit payment details to api
  // Set payment token with value received
  // Display toast if error occurs
}

const submitForm = async () => {
  try {
    isSubmitting.value = true;
    const team = await createNewTeam({ teamName: teamName.value });
    if (!team) throw new Error("Xaibe Error: Team not generated.");
    await getActiveUser();
    mixpanelWrapper.client.track("New team created");
    isSubmitting.value = false;
    router.push(`/teams/${team.id}`);
  } catch (error) {
    setActiveToast({
      message: "An error occurred while creating your new team.",
      type: ToastType.ERROR,
    });
    isSubmitting.value = false;
  }
};
</script>
