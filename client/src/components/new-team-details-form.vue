<template> 
  <div>
    <base-input v-model="formValues.numberOfSeats" id="numberOfSeats" type="number" label="How many seats do you want to include?" />
    <base-input v-model="formValues.businessName" id="businessName" label="Business/Institution name *" />
    <div class="grid grid-cols-2 gap-4">
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
import { useToastStore } from "src/state";
import { 
  ToastType, 
  NewTeamDetailsFormSchema, 
  NewTeamDetailsFormValidator
} from "src/types";

const emits = defineEmits<{
  (event: "onContinue", formValues: any): void;
}>();

const { setActiveToast } = useToastStore();

const formValues = reactive<NewTeamDetailsFormSchema>({
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
    NewTeamDetailsFormValidator.parse(formValues);
    emits("onContinue", formValues);
  } catch (error) {
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred submitting the team details."
    });
  }
}
</script>
