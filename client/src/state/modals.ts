import { defineStore } from "pinia";

interface ModalStoreState {
  isNewElementModalShown: boolean;
  isDeleteApplicationConfirmationModalShown: boolean;
}

export const useModalStore = defineStore("modals", {
  state: (): ModalStoreState => {
    return {
      isNewElementModalShown: false,
      isDeleteApplicationConfirmationModalShown: false,
    };
  },
  actions: {
    setIsNewElementModalShown(newValue: boolean): void {
      this.isNewElementModalShown = newValue;
    },
    setIsDeleteApplicationConfirmationModalShown(newValue: boolean): void {
      this.isDeleteApplicationConfirmationModalShown = newValue;
    },
  },
});
