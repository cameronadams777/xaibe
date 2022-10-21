<script lang="ts" setup>
defineProps<{ isOpen: boolean }>();
</script>

<template>
  <div
    v-if="isOpen"
    class="absolute bg-white top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-1/4 h-1/4 shadow-md rounded-lg z-10"
  >
    <slot />
  </div>
</template>

<style>
.popin-from-center {
  animation: popin 0.3s linear;
}

@keyframes popin {
  0% {
    transform: scale(0);
  }
  80% {
    transform: scale(1.25);
  }
  100% {
    transform: scale(1);
  }
}
</style>
