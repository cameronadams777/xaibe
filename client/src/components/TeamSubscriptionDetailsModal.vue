<template>
  <BaseModal :is-open="isOpen" class="w-7/12 h-3/4">
    <div class="relative h-full p-4 md:p-16 lg:p-4 flex flex-col">
      <BaseFabButton class="absolute top-2 right-2" @click="close">
        <XMarkIcon class="w-8 h-8" />
      </BaseFabButton>
      <h2 class="text-2xl m-0 p-0">
        {{team?.name}}
      </h2>
      <div class="flex flex-col items-center">
        <p><b>Number of Seats:</b></p>
        <PlusMinusNumberInput :value="updatedNumberOfSeats" @on-update="updateSeats" />
      </div>
      <BaseButton :disabled="!hasUpdates" text="Update" class="w-1/2 mt-10 self-center" />
    </div>
  </BaseModal>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import { Team } from "src/types";
import { emptyTeamSubscriptionDetailsProps, useModalStore } from "src/state";
import PlusMinusNumberInput from "src/components/PlusMinusNumberInput.vue";

const props = defineProps<{
  isOpen: boolean;
  team?: Team;
}>();

const updatedNumberOfSeats = ref<number>(0);
const hasUpdates = computed<boolean>(() => updatedNumberOfSeats.value !== props.team?.activeNumberOfSeats);

const router = useRouter();
const { setTeamSubscriptionDetailsProps } = useModalStore();

const updateSeats = (value: number) => updatedNumberOfSeats.value = value;
const close = () => setTeamSubscriptionDetailsProps(emptyTeamSubscriptionDetailsProps); 

onMounted(() => {
  if (!props.team) {
    router.push("/500");
    close();
    return;
  }
  updatedNumberOfSeats.value = props.team.activeNumberOfSeats;
});
</script>
