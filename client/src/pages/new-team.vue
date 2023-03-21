<template>
  <the-main-layout>
    <div class="w-full h-full flex flex-col justify-center items-center">
      <h2>Create a New Team</h2>
      <NewTeamDetailsForm v-if="!stripeClientSecret.length" @onContinue="createNewTeamWithSubscription"/>
      <PaymentForm v-else :client-secret="stripeClientSecret"/>
    </div>
  </the-main-layout>
</template>

<script lang="ts" setup>
import { computed, ref } from "vue";
import { useRouter } from "vue-router";
import { createNewTeam } from "src/api/teams";
import { createNewStripeCustomer } from "src/api/payments"; 
import TheMainLayout from "src/layouts/the-main-layout.vue";
import NewTeamDetailsForm from "src/components/new-team-details-form.vue";
import PaymentForm from "src/components/payment-form.vue";
import { useToastStore } from "src/state";
import { useActiveUserStore } from "src/state/active-user";
import { ToastType, NewTeamDetailsFormSchema } from "src/types";
import { mixpanelWrapper } from "src/tools/mixpanel";
import { storeToRefs } from "pinia";

const metadata = ref<NewTeamDetailsFormSchema>();
const stripeClientSecret = ref("");
const paymentToken = ref("");

const router = useRouter();
const { getActiveUser, ...activeUserStore } = useActiveUserStore();
const { setActiveToast } = useToastStore();

const { activeUser } = storeToRefs(activeUserStore);

const teamName = ref("");
const isSubmitting = ref(false);

const hasMetadata = computed(() => metadata.value != null);

const createNewTeamWithSubscription = async (formValues: Record<string, any>): Promise<void> => {
  if (!activeUser?.value) {
    router.push("/500");
    return;
  }
  if (!activeUser.value.stripeId) {
    const success = await createNewStripeCustomer(formValues);
    if (!success) {
      setActiveToast({
        type: ToastType.ERROR,
        message: "An unknown error occurred while attempting to create a new team.",
      });
      return;
    }
    await getActiveUser();
  }
  const { clientSecret } = await createNewTeam({ 
    teamName: formValues.teamName, 
    numberOfSeats: formValues.numberOfSeats, 
  });
  stripeClientSecret.value = clientSecret;
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
