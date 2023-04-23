<template>
  <the-main-layout>
    <div
      class="mt-8 xl:mt-0 px-8 w-full h-full flex flex-col justify-center items-center"
    >
      <h2>Create a New Team</h2>
      <NewTeamDetailsForm
        v-if="!stripeClientSecret?.length"
        @onContinue="createNewTeamWithSubscription"
      />
      <PaymentForm
        v-else
        :stripe-client-secret="stripeClientSecret"
        @onSubmit="(paymentIntent: PaymentIntent, token: Token) => submitForm(paymentIntent, token)"
      />
    </div>
  </the-main-layout>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { createNewTeam } from "src/api/teams";
import {
  createNewStripeCustomer,
  confirmPaymentIntent,
} from "src/api/payments";
import TheMainLayout from "src/layouts/the-main-layout.vue";
import NewTeamDetailsForm from "src/components/new-team-details-form.vue";
import PaymentForm from "src/components/payment-form.vue";
import { useToastStore } from "src/state";
import { useActiveUserStore } from "src/state/active-user";
import { Team, ToastType } from "src/types";
import { mixpanelWrapper } from "src/tools/mixpanel";
import { storeToRefs } from "pinia";
import { loadStripe, PaymentIntent, Token } from "@stripe/stripe-js";
import { config } from "src/config";

const stripeClientSecret = ref("");

const router = useRouter();
const { getActiveUser, ...activeUserStore } = useActiveUserStore();
const { setActiveToast } = useToastStore();

const { activeUser } = storeToRefs(activeUserStore);

const team = ref<Team>();
const isSubmitting = ref<boolean>(false);

const createNewTeamWithSubscription = async (
  formValues: Record<string, any>
): Promise<void> => {
  try {
    if (!activeUser?.value?.stripeId) {
      const success = await createNewStripeCustomer(formValues);
      if (!success) {
        setActiveToast({
          type: ToastType.ERROR,
          message:
            "An unknown error occurred while attempting to create a new team.",
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
    console.error(error);
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred when submitting your information",
    });
  }
};

const submitForm = async () => {
  try {
    if (!team.value)
      throw new Error("Galata Error: Undefined team after payment processing");
    setActiveToast({
      type: ToastType.SUCCESS,
      message: "New Team Created!",
    });
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
