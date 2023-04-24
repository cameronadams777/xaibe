<template>
  <TheMainLayout>
    <div class="p-4">
      <h2>{{ activeUser?.firstName }} {{ activeUser?.lastName }}</h2>
      <div class="flex flex-col items-start">
        <label for="allowTelemetry">Anonymous Telemetry</label>
        <select 
          id="allowTelemetry" 
          :value="allowTelemetry" 
          @change="handleTelemetrySelect"
        >
          <option value="true">Enabled</option>
          <option value="false">Disabled</option>
        </select>
      </div>
    </div>
  </TheMainLayout>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { storeToRefs } from "pinia";
import { useActiveUserStore } from "../state/active-user";
import TheMainLayout from "../layouts/TheMainLayout.vue";

const activeUserStore = useActiveUserStore();
const { activeUser } = storeToRefs(activeUserStore);

const allowTelemetry = ref(localStorage.getItem("allowTelemetry") === "true");

const handleTelemetrySelect = (event: Event): void => {
  const newValue = (event.target as HTMLSelectElement).value;
  allowTelemetry.value = newValue === "true";
  localStorage.setItem("allowTelemetry", newValue);
}
</script>

