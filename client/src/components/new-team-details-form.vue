<template> 
  <div>
    <base-input v-model="formValues.numberOfSeats" id="numberOfSeats" type="number" label="How many seats do you want to include?" />
    <base-input v-model="formValues.businessName" id="businessName" label="Business/Institution name *" />
    <div class="grid grid-cols-2 gap">
      <base-input v-model="formValues.addressOne" id="addressOne" label="Address (P.O. box, company name, c/o) *" />
      <base-input v-model="formValues.addressTwo" id="addressTwo" label="Address line 2 (Apartment, suite, unit)" />
      <base-input v-model="formValues.city" id="city" label="City *" />
      <base-input v-model="formValues.zip" id="postalCode" label="Postal/Zip code" />
      <base-input v-model="formValues.country" id="country" label="Country/Region*" />
      <base-input v-model="formValues.state" id="state" label="State/Province" />
    </div>
    <base-input v-model="formValues.billingEmail" id="billingEmail" label="Billing Email *" />
    <base-button text="Continue to payment method" @click="submitForm" />
  </div>
</template>

<script setup lang="ts">
  import { reactive } from "vue";
  
  const emit = defineEmits<{ onContinue: (formValues: any) => void; }>();

  // TODO: Add validation to form

  const formValues = reactive({
    numberOfSeats: 1,
    businessName: "",
    addressOne: "",
    addressTwo: "",
    city: "",
    state: "",
    country: "",
    zip: "",
    billingEmail: ""
  });

  const submitForm = (): void => {
    try {
      emit.onContinue(formValues);
    } catch (error) {
      console.error(error);
    }
  }
</script>
