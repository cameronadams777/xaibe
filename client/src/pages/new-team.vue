<template>
  <the-main-layout>
    <div class="w-full h-full flex flex-col justify-center items-center">
      <h2>Create a New Team</h2>
      <NewTeamDetailsForm v-if="!stripeClientSecret?.length" @onContinue="createNewTeamWithSubscription"/>
      <PaymentForm v-else :client-secret="stripeClientSecret" @onSubmit="(token: Token) => submitForm(token)"/>
    </div>
  </the-main-layout>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { createNewTeam } from "src/api/teams";
import { createNewStripeCustomer, confirmPaymentIntent } from "src/api/payments"; 
import TheMainLayout from "src/layouts/the-main-layout.vue";
import NewTeamDetailsForm from "src/components/new-team-details-form.vue";
import PaymentForm from "src/components/payment-form.vue";
import { useToastStore } from "src/state";
import { useActiveUserStore } from "src/state/active-user";
import { Team, ToastType } from "src/types";
import { mixpanelWrapper } from "src/tools/mixpanel";
import { storeToRefs } from "pinia";
import { Token } from "@stripe/stripe-js";

const stripeClientSecret = ref("");

const router = useRouter();
const { getActiveUser, ...activeUserStore } = useActiveUserStore();
const { setActiveToast } = useToastStore();

const { activeUser } = storeToRefs(activeUserStore);

const team = ref<Team>();
const isSubmitting = ref<boolean>(false);

const createNewTeamWithSubscription = async (formValues: Record<string, any>): Promise<void> => {
  try {
    if (!activeUser?.value?.stripeId) {
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
    const { team: newTeam, clientSecret } = await createNewTeam({ 
      teamName: formValues.teamName, 
      numberOfSeats: formValues.numberOfSeats, 
    });
    team.value = newTeam;
    stripeClientSecret.value = clientSecret;
  } catch (error) {
    console.error(error)
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred when submitting your information"
    });
  }
};

const submitForm = async (apiToken: Token) => {
  try {
    console.log(stripeClientSecret.value, team.value, apiToken);
    if (!stripeClientSecret.value || !team.value || !apiToken?.card) throw new Error();
    isSubmitting.value = true;
    const success = await confirmPaymentIntent({ 
      paymentIntent: stripeClientSecret.value,
      cardToken: apiToken.card.id
    })
    if (!success) throw new Error();
    await getActiveUser();
    mixpanelWrapper.client.track("New team created");
    isSubmitting.value = false;
    router.push(`/teams/${team.value.id}`);
  } catch (error) {
    setActiveToast({
      message: "An error occurred while creating your new team.",
      type: ToastType.ERROR,
    });
    isSubmitting.value = false;
  }
};
</script>
