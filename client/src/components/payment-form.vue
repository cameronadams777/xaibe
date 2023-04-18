<template>
  <div class="w-full md:w-1/2 lg:w-1/4">
    <div ref="cardElement" id="card-element"></div>
    <div class="mt-4 w-full flex justify-center items-center">
      <base-button text="Submit" class="w-1/2" @click="submit" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { Stripe, StripeCardElement, loadStripe } from "@stripe/stripe-js";
import { useRouter } from "vue-router";
import { config } from "src/config";
import { useToastStore } from "src/state";
import { ToastType } from "src/types";

const props = defineProps<{ stripeClientSecret: string }>();

const emits = defineEmits<{
  (event: "onSubmit"): void;
}>();

const router = useRouter();
const { setActiveToast } = useToastStore();

const stripe = ref<Stripe | null>(null);
const cardElement = ref<StripeCardElement | null>(null);

onMounted(async () => {
  try {
    stripe.value = await loadStripe(config.stripePublishableKey);
    if (!stripe.value) {
      router.push("/500");
      return;
    }
    const elements = stripe.value.elements();
    elements.update({
      appearance: {
        theme: "stripe",
      },
    });
    cardElement.value = elements.create("card");
    cardElement.value.mount("#card-element");
  } catch (error) {
    console.error(error);
    router.push("/500");
  }
});

const submit = async () => {
  if (!stripe.value || !cardElement.value) {
    router.push("/500");
    return;
  }
  const result = await stripe.value.createToken(cardElement.value);
  if (!result.token) {
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred while submitting payment. Please try again.",
    });
    return;
  }
  const confirmation = await stripe.value.confirmCardPayment(
    props.stripeClientSecret,
    {
      payment_method: {
        card: cardElement.value,
      },
    }
  );
  if (!confirmation.paymentIntent)
    throw new Error("Galata Error: Card payment not confirmed");
  emits("onSubmit");
};
</script>
