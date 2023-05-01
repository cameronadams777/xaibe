<template>
  <BaseModal :is-open="isOpen" class="w-7/12 h-3/4">
    <div class="relative h-full p-4 md:p-16 lg:p-4 flex flex-col">
      <BaseFabButton class="absolute top-2 right-2" @click="close">
        <XMarkIcon class="w-8 h-8" />
      </BaseFabButton>
      <TeamSubscriptionDetailsModalReviewStep
        v-if="formStep === 'review'"
        :team="team as Team"
        @on-continue="updateFormStep"
      />
      <TeamSubscriptionDetailsModalConfirmationStep
        v-else
        @on-confirm="confirmUpdate"
        @on-close="close"
      />
    </div>
  </BaseModal>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import { Team, ToastType } from "src/types";
import {
  emptyTeamSubscriptionDetailsProps,
  useModalStore,
  useToastStore,
} from "src/state";
import TeamSubscriptionDetailsModalReviewStep from "./TeamSubscriptionDetailsModalReviewStep.vue";
import TeamSubscriptionDetailsModalConfirmationStep from "./TeamSubscriptionDetailsModalConfirmationStep.vue";
import { updateTeamSeatCount } from "src/api/teams";

type TeamSubscriptionDetailsFormStep = "review" | "confirm";

const props = defineProps<{
  isOpen: boolean;
  team?: Team;
}>();

const router = useRouter();
const { setTeamSubscriptionDetailsProps } = useModalStore();
const { setActiveToast } = useToastStore();

const formStep = ref<TeamSubscriptionDetailsFormStep>("review");
const newSeatCount = ref<number>(0);

const updateFormStep = (updatedSeatCount: number): void => {
  newSeatCount.value = updatedSeatCount;
  formStep.value = "confirm";
};

const confirmUpdate = async (): Promise<void> => {
  if (!props.team) return;
  try {
    await updateTeamSeatCount({
      teamId: props.team.id,
      newSeatCount: newSeatCount.value,
    });
    close();
  } catch (error) {
    console.error(error);
    setActiveToast({
      type: ToastType.ERROR,
      message: "An error occurred while trying to update your subscription.",
    });
  }
};

const close = () =>
  setTeamSubscriptionDetailsProps(emptyTeamSubscriptionDetailsProps);

onMounted(() => {
  if (!props.team) {
    router.push("/500");
    close();
    return;
  }
});
</script>
