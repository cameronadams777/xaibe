<template>
  <the-main-layout>
    <div class="w-full h-full flex flex-col justify-center items-center">
      <h2>Create a New Team</h2>
      <div class="flex flex-col w-1/4 mb-2">
        <label for="teamName" class="font-bold mb-2">Team Name</label>
        <input
          v-model="teamName"
          id="teamName"
          name="teamName"
          type="text"
          class="p-1"
        />
      </div>
      <base-button
        text="Create"
        :text-size="ButtonTextSize.LARGE"
        :show-spinner="isSubmitting"
        :disabled="isSubmitting"
        :aria-disabled="isSubmitting"
        @click="submitForm"
      />
    </div>
  </the-main-layout>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { createNewTeam } from "../api/teams";
import TheMainLayout from "../layouts/the-main-layout.vue";
import { useToastStore } from "../state";
import { useActiveUserStore } from "../state/active-user";
import { ButtonTextSize, ToastType } from "../types";
import { mixpanelWrapper } from "src/tools/mixpanel";

const router = useRouter();
const { getActiveUser } = useActiveUserStore();
const { setActiveToast } = useToastStore();

const teamName = ref("");
const isSubmitting = ref(false);

const submitForm = async () => {
  try {
    isSubmitting.value = true;
    const team = await createNewTeam({ teamName: teamName.value });
    if (!team) throw new Error("Galata Error: Team not generated.");
    await getActiveUser();
    mixpanelWrapper.client.track("New team created");
    isSubmitting.value = false;
    router.push(`/teams/${team.ID}`);
  } catch (error) {
    setActiveToast({
      message: "An error occurred while creating your new team.",
      type: ToastType.ERROR,
    });
    isSubmitting.value = false;
  }
};
</script>
