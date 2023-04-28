<template>
  <BaseModal :is-open="isOpen" class="w-7/12 h-3/4">
    <div class="relative h-full p-16">
      <BaseFabButton class="absolute top-2 right-2" @click="close">
        <XMarkIcon class="w-8 h-8" />
      </BaseFabButton>
      <h2 class="text-2xl">
        {{team?.name}}
      </h2>
      <p><b>Number of Seats:</b> {{team?.activeNumberOfSeats}}</p>
    </div>
  </BaseModal>
</template>

<script lang="ts" setup>
import { onMounted } from "vue";
import { useRouter } from "vue-router";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import { Team } from "src/types";
import { emptyTeamSubscriptionDetailsProps, useModalStore } from "src/state";

const props = defineProps<{
  isOpen: boolean;
  team?: Team;
}>();

const router = useRouter();
const { setTeamSubscriptionDetailsProps } = useModalStore();

const close = () => setTeamSubscriptionDetailsProps(emptyTeamSubscriptionDetailsProps); 

onMounted(() => {
  if (!props.team) {
    router.push("/500");
    close();
  }
});
</script>
