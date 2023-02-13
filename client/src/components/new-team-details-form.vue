<template> 
  <div>
    <base-input v-model="formValues.numberOfSeats" type="number" label="How many seats do you want to include?" />
    <base-input v-model="formValues.businessName" label="Business/Institution name *" />
    <base-input v-model="formValues.addressOne" label="Address (P.O. box, company name, c/o) *" />
    <base-input v-model="formValues.addressTwo" label="Address line 2 (Apartment, suite, unit)" />
    <div class="grid">
      <base-input v-model="formValues.city" label="City *" />
      <base-input v-model="formValues.zip" label="Postal/Zip code" />
      <base-input v-model="formValues.country" label="Country/Region*" />
      <base-input v-model="formValues.state" label="State/Province" />
    </div>
    <base-input v-model="formValues.billingEmail" label="Billing Email *" />
    <base-button text="Continue to payment method" />
  </div>
</template>

<script setup lang="ts">
  import { reactive } from "vue";
  import { z } from "zod"; 
  
  const emit = defineEmits<{ onContinue: (formValues: any) => void; }>();

  const newTeamFormValuesSchema = z.object({
    numberOfSeats: z.number(),
    businessName: z.string(),
    addressOne: z.string(),
    addressTwo: z.string(),
    city: z.string(),
    country: z.string(),
    zip: z.string(),
    state: z.string(),
    billingEmail: z.string(),
  });

  const formValues: z.infer<typeof newTeamFormValuesSchema> = reactive({
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
      newTeamFormValuesSchema.parse(formValues);
      emit.onContinue(formValues);
    } catch (error) {
      console.error(error);
    }
  }
</script>
