<template>
  <h2 class="text-2xl m-0 p-0">
    {{ team?.name }}
  </h2>
  <div class="flex flex-col items-center">
    <p><b>Number of Seats:</b></p>
    <PlusMinusNumberInput
      :value="updatedNumberOfSeats"
      @on-update="updateSeats"
    />
  </div>
  <BaseButton
    :disabled="!hasUpdates"
    text="Update"
    class="w-1/2 mt-10 self-center disabled:cursor-default"
    @click="submit"
  />
</template>

<script setup lang="ts">
import { useToastStore } from "src/state";
import { Team, ToastType } from "src/types";
import { computed, onMounted, ref } from "vue";
import PlusMinusNumberInput from "./PlusMinusNumberInput.vue";

const props = defineProps<{ team: Team }>();
const emits = defineEmits<{
  (event: "onContinue", updatedNumberOfSeats: number): void;
}>();

const { setActiveToast } = useToastStore();

const updatedNumberOfSeats = ref<number>(0);
const hasUpdates = computed<boolean>(
  () => updatedNumberOfSeats.value !== props.team?.activeNumberOfSeats
);

const updateSeats = (value: number) => (updatedNumberOfSeats.value = value);
const submit = () => {
  if (props.team.users.length > updatedNumberOfSeats.value) {
    setActiveToast({
      type: ToastType.ERROR,
      message:
        "You are not able to remove seats that currently belong to users. Please remove users you do not want on the team and try again.",
    });
    return;
  }
  emits("onContinue", updatedNumberOfSeats.value);
};

onMounted(() => {
  updatedNumberOfSeats.value = props.team.activeNumberOfSeats;
});
</script>
