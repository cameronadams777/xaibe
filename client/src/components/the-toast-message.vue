<script lang="ts" setup>
import { storeToRefs } from "pinia";
import { useToastStore } from "../state";
import { ToastType } from "../types";

const toastStore = useToastStore();
const { activeToast, toastTitleByType } = storeToRefs(toastStore);
</script>

<template>
  <transition name="slide-fade">
    <div
      v-if="activeToast != null"
      class="absolute top-5 right-5 p-4 bg-white rounded-md shadow-md"
    >
      <div class="relative">
        <h3
          class="m-0"
          :class="{
            'text-indigo-500': activeToast.type === ToastType.INFO,
            'text-red-500': activeToast.type === ToastType.ERROR,
            'text-green-500': activeToast.type === ToastType.SUCCESS,
          }"
        >
          {{ toastTitleByType }}
        </h3>
        <p class="m-0">{{ activeToast?.message }}</p>
      </div>
    </div>
  </transition>
</template>

<style>
.slide-fade-enter-active {
  transition: all 0.3s ease;
}
.slide-fade-leave-active {
  transition: all 0.8s;
}
.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(10px);
  opacity: 0;
}
</style>
