<template>
  <div>
    <base-input
      v-model="formValues.numberOfSeats"
      id="numberOfSeats"
      type="number"
      label="How many seats do you want to include?"
    />
    <base-input
      v-model="formValues.businessName"
      id="businessName"
      label="Business/Institution name *"
    />
    <base-input
      v-model="formValues.teamName"
      id="teamName"
      label="Team Name *"
    />
    <div class="grid grid-cols-2 gap-4">
      <base-input
        v-model="formValues.addressLineOne"
        id="addressOne"
        label="Address (P.O. box, company name, c/o) *"
      />
      <base-input
        v-model="formValues.addressLineTwo"
        id="addressTwo"
        label="Address line 2 (Apartment, suite, unit)"
      />
      <base-input v-model="formValues.city" id="city" label="City *" />
      <base-input
        v-model="formValues.postalCode"
        id="postalCode"
        label="Postal/Zip code"
      />
      <div class="flex flex-col mb-2">
        <label for="country" class="font-bold mb-2">Country</label>
        <select
          id="country"
          v-model="formValues.country"
          class="p-1 mb-2"
          placeholder="Please Select One"
        >
          <option disabled value="">Please Select One</option>
          <option
            v-for="country of countriesList"
            :id="country"
            :value="country"
          >
            {{ country }}
          </option>
        </select>
      </div>
      <div class="flex flex-col mb-2">
        <label for="state" class="font-bold mb-2">State</label>
        <select
          v-model="formValues.state"
          :disabled="!states.length"
          id="state"
          class="p-1 mb-2"
          placeholder="Please Select One"
        >
          <option disabled value="">Please Select One</option>
          <option v-for="state of states" :value="state">{{ state }}</option>
        </select>
      </div>
    </div>
    <base-input
      v-model="formValues.billingEmail"
      id="billingEmail"
      label="Billing Email *"
    />
    <base-button text="Continue to payment method" @click="submitForm" />
  </div>
</template>

<script setup lang="ts">
import { countries } from "src/constants";
import { computed, reactive } from "vue";

const emits = defineEmits<{
  (event: "onContinue", formValues: any): void;
}>();

const countriesList = computed(() =>
  countries.map((country) => country.country)
);

const states = computed(() =>
  formValues.country.length
    ? countries.find((country) => country.country === formValues.country)
        ?.states ?? []
    : []
);

const formValues = reactive({
  teamName: "",
  numberOfSeats: 1,
  businessName: "",
  addressLineOne: "",
  addressLineTwo: "",
  city: "",
  state: "",
  country: "",
  postalCode: "",
  billingEmail: "",
});

const submitForm = (): void => {
  emits("onContinue", formValues);
};
</script>
