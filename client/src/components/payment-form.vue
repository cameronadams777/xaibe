<template>
  <StripeElements
    v-if="stripeLoaded"
    v-slot="{ elements }"
    ref="elms"
    :stripe-key="config.stripePublishableKey"
    :elements-options="elementsOptions"
  >
    <StripeElement
      ref="card"
      :elements="elements"
      :options="cardOptions"
    />
  </StripeElements>
  <button type="button" @click="pay">Pay</button>
</template>

<script setup lang="ts">
import { onBeforeMount, ref }  from "vue";
import { StripeElement, StripeElements } from "vue-stripe-js";
import { loadStripe } from "@stripe/stripe-js";
import { config } from "src/config";

const emits = defineEmits<{ onFinish: () => void }>(); 

const elementsOptions = ref({
  clientSecret: config.stripeClientSecret,
  appearance: "stripe",
});
const cardOptions = ref({});
const stripeLoaded = ref<boolean>(false);
const card = ref();
const elms = ref();

onBeforeMount(() => {
  const stripePromise = loadStripe(config.stripePublishableKey);
  stripePromise.then(() => {
    stripeLoaded.value = true;
  });
});

const pay = () => {
  try {
    const cardElement = card.value.stripeElement;

    elms.value.instance.createToken(cardElement)
      .then(console.log);
  } catch (error) {
    // TODO: Display toast error 
  }
}
</script>
