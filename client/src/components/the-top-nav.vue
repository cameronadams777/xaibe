<script lang="ts" setup>
import { Menu, MenuButton, MenuItem, MenuItems } from "@headlessui/vue";
import { Bars3Icon } from "@heroicons/vue/24/solid";
import { UserIcon, PlusIcon } from "@heroicons/vue/24/outline";
import { useModalStore } from "../state/modals";
import { useRouter } from "vue-router";

const emits = defineEmits<{ (event: "onMenuClick"): void }>();

const router = useRouter();
const { setIsNewElementModalShown } = useModalStore();

const userNavigation = [
  {
    name: "Settings",
    cb: () => router.push("/settings"),
  },
  {
    name: "Log Out",
    cb: () => {
      localStorage.removeItem("token"); // TODO: Move this to tauri side
      router.push("/login");
    },
  },
];
</script>

<template>
  <div class="p-4 border-b border-gray-300 flex justify-between">
    <div class="flex justify-center">
      <button
        class="bg-white border-none mr-2 cursor-pointer hover:bg-gray-200 rounded-xl"
        @click="emits('onMenuClick')"
      >
        <bars-3-icon class="h-full w-6" />
      </button>
    </div>
    <div class="flex justify-end">
      <button
        class="w-8 h-8 mr-2 bg-white text-gray-800 flex justify-center items-center text-white focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 border-none rounded-full cursor-pointer hover:shadow-md transition-all duration-500"
        @click="setIsNewElementModalShown(true)"
      >
        <plus-icon class="w-8" />
      </button>
      <Menu as="div" class="relative ml-3">
        <div>
          <MenuButton
            class="w-8 h-8 flex max-w-xs items-center rounded-full bg-white text-sm border-transparent cursor-pointer hover:shadow-md transition-all duration-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
          >
            <span class="sr-only">Open user menu</span>
            <user-icon class="w-4" />
          </MenuButton>
        </div>
        <transition
          enter-active-class="transition ease-out duration-200"
          enter-from-class="transform opacity-0 scale-95"
          enter-to-class="transform opacity-100 scale-100"
          leave-active-class="transition ease-in duration-75"
          leave-from-class="transform opacity-100 scale-100"
          leave-to-class="transform opacity-0 scale-95"
        >
          <MenuItems
            class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
          >
            <MenuItem v-for="item in userNavigation" :key="item.name">
              <button
                class="w-full block px-4 py-2 text-sm bg-white text-gray-700 text-left border-none no-underline hover:bg-gray-300 transition-all duration-300"
                @click="item.cb"
              >
                {{ item.name }}
              </button>
            </MenuItem>
          </MenuItems>
        </transition>
      </Menu>
    </div>
  </div>
</template>
