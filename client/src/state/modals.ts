import { defineStore } from "pinia";

interface ModalStoreState {
  isNewElementModalShown: boolean;
}

export const useModalStore = defineStore("modals", {
  state: (): ModalStoreState => {
    return { isNewElementModalShown: false };
  },
  actions: {
    setIsNewElementModalShown(newValue: boolean): void {
      this.isNewElementModalShown = newValue;
    },
  },
});
