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

<template>
  <base-modal :is-open="isOpen">
    <div class="relative h-full flex flex-col justify-center items-center">
      <base-fab-button
        class="absolute top-2 right-2"
        @click="setIsNewElementModalShown(false)"
      >
        <x-mark-icon class="w-8 h-8" />
      </base-fab-button>
      <base-button
        text="Create New Application"
        :variant="ButtonVariant.WHITE"
        class="w-1/2"
        @click="navigateToNewAppPage"
      />
      <base-button
        text="Create New Team"
        :variant="ButtonVariant.WHITE"
        class="w-1/2"
        @click="navigateToNewTeamPage"
      />
    </div>
  </base-modal>
</template>
