<template>
  <div class="flex flex-col w-1/4 mb-3">
    <label for="applicationname" class="font-bold mb-2">Application name</label>
    <input
      v-model="applicationname"
      id="applicationname"
      name="applicationname"
      type="applicationname"
      placeholder="My Awesome App"
      class="p-1.5 mb-4"
    />
    <div v-if="activeUser?.teams?.length" class="w-full">
      <label class="text-base font-medium text-gray-900"
        >Is this yours or a team's application?</label
      >
      <fieldset class="mt-4 border-none">
        <legend class="sr-only">Notification method</legend>
        <div class="flex space-x-4">
          <div
            v-for="method in applicationOwnershipMethods"
            :key="method.id"
            class="flex items-center"
          >
            <input
              v-model="applyType"
              :id="method.id"
              :value="method.id"
              name="notification-method"
              type="radio"
              class="h-4 w-5 border-gray-300 text-indigo-600 focus:ring-indigo-500"
            />
            <label
              :for="method.id"
              class="ml-3 block text-sm font-medium text-gray-700"
              >{{ method.title }}</label
            >
          </div>
        </div>
      </fieldset>
    </div>
    <div v-if="applyType === 'team'" class="mb-4">
      <label for="team" class="block text-sm font-medium text-gray-700"
        >Team</label
      >
      <select
        v-model="teamId"
        id="team"
        name="team"
        class="mt-1 block w-full rounded-md border-gray-300 py-2 pl-3 pr-10 text-base focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
      >
        <option disabled value="">Please select one</option>
        <option v-for="team of activeUser?.teams" :value="team.ID">
          {{ team.name }}
        </option>
      </select>
    </div>
    <BaseButton
      text="Create"
      :text-size="ButtonTextSize.LARGE"
      :show-spinner="isSubmitting"
      :disabled="isSubmitting"
      :aria-disabled="isSubmitting"
      @click="onCreate"
    />
  </div>
</template>

<script lang="ts" setup>
import { storeToRefs } from "pinia";
import { useActiveUserStore } from "src/state";
import { ButtonTextSize, ApplicationType } from "src/types";
import { onMounted, ref } from "vue";

type NewApplicationApplyType = "user" | "team";

const props = defineProps<{
  applicationType: ApplicationType;
  isSubmitting: boolean;
}>();

const emits = defineEmits<{
  (event: "onCreate", applicationname: string, teamId?: string): void;
}>();

const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);

const applicationOwnershipMethods = [
  { id: "user", title: "User" },
  { id: "team", title: "Team" },
];

const applicationname = ref("");
const teamId = ref<string>("");
const applyType = ref<NewApplicationApplyType>("user");

onMounted(() => {
  if (props.applicationType === ApplicationType.AIRBRAKE)
    applicationname.value = "Airbrake";
  else if (props.applicationType === ApplicationType.NEWRELIC)
    applicationname.value = "NewRelic";
  else if (props.applicationType === ApplicationType.SENTRY)
    applicationname.value = "Sentry";
});

const onCreate = () => {
  emits("onCreate", applicationname.value, teamId.value);
};
</script>
