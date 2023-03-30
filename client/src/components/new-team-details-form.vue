<template> 
  <div>
    <base-input v-model="formValues.numberOfSeats" id="numberOfSeats" type="number" label="How many seats do you want to include?" />
    <base-input v-model="formValues.businessName" id="businessName" label="Business/Institution name *" />
    <base-input v-model="formValues.teamName" id="teamName" label="Team Name *" />
    <div class="grid grid-cols-2 gap-4">
      <base-input v-model="formValues.addressLineOne" id="addressOne" label="Address (P.O. box, company name, c/o) *" />
      <base-input v-model="formValues.addressLineTwo" id="addressTwo" label="Address line 2 (Apartment, suite, unit)" />
      <base-input v-model="formValues.city" id="city" label="City *" />
      <base-input v-model="formValues.postalCode" id="postalCode" label="Postal/Zip code" />
      <base-input v-model="formValues.country" id="country" label="Country/Region*" />
      <base-input v-model="formValues.state" id="state" label="State/Province" />
    </div>
    <base-input v-model="formValues.billingEmail" id="billingEmail" label="Billing Email *" />
    <base-button text="Continue to payment method" @click="submitForm" />
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";

const emits = defineEmits<{
  (event: "onContinue", formValues: any): void;
}>();

const formValues = reactive({
  teamName: "Dev",
  numberOfSeats: 1,
  businessName: "Galata.app",
  addressLineOne: "5319 Vermilion Ct",
  addressLineTwo: "",
  city: "Dickinson",
  state: "Texas",
  country: "United States",
  postalCode: "77539",
  billingEmail: "cameron.adams@hey.com"
});

const submitForm = (): void => {
  emits("onContinue", formValues);
}
</script>
