<script lang="ts" setup>
import { useAttrs, watch } from "vue";

const props = defineProps<{ isOpen: boolean }>();
const attrs = useAttrs();

watch(
  () => props.isOpen,
  (value) => {
    const bodyEl = document.querySelector("body");
    if (!bodyEl) return;
    if (value) bodyEl.style.overflowY = "hidden";
    else bodyEl.style.overflowY = "auto";
  }
);
</script>

<template>
  <Teleport to="#modal-target">
    <transition name="bounce">
      <div
        v-if="isOpen"
        class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 bg-white w-1/2 h-1/3 lg:w-1/4 lg:h-1/4 shadow-md rounded-lg z-10"
        :class="attrs.class"
      >
        <slot />
      </div>
    </transition>
  </Teleport>
</template>

<style>
.bounce-enter-active {
  animation: bounce-in 0.5s;
}
.bounce-leave-active {
  animation: bounce-in 0.5s reverse;
}
@keyframes bounce-in {
  0% {
    transform: translate(-50%, -50%) scale(0);
  }
  100% {
    transform: translate(-50%, -50%) scale(1);
  }
}
</style>
