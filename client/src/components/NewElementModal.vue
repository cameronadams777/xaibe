<template>
  <BaseModal :is-open="isOpen">
    <div class="relative h-full flex flex-col justify-center items-center">
      <BaseFabButton
        class="absolute top-2 right-2"
        @click="setIsNewElementModalShown(false)"
      >
        <XMarkIcon class="w-8 h-8" />
      </BaseFabButton>
      <BaseButton
        text="Create New Application"
        :variant="ButtonVariant.PRIMARY"
        class="w-1/2"
        @click="navigateToNewAppPage"
      />
      <BaseButton
        text="Create New Team"
        :variant="ButtonVariant.PRIMARY"
        class="w-1/2"
        @click="navigateToNewTeamPage"
      />
    </div>
  </BaseModal>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router";
import { XMarkIcon } from "@heroicons/vue/24/outline";
import { useModalStore } from "src/state/modals";
import { ButtonVariant } from "src/types";

defineProps<{ isOpen: boolean }>();
const router = useRouter();
const { setIsNewElementModalShown } = useModalStore();

const navigateToNewAppPage = () => {
  router.push("/applications/new");
  setIsNewElementModalShown(false);
};

const navigateToNewTeamPage = () => {
  router.push("/teams/new");
  setIsNewElementModalShown(false);
};
</script>

