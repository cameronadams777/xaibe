import { defineStore } from "pinia";

interface ModalStoreState {
  isNewElementModalShown: boolean;
  isDeleteApplicationConfirmationModalShown: boolean;
  isDeleteTeamConfirmationModalShown: boolean;
}

export const useModalStore = defineStore("modals", {
  state: (): ModalStoreState => {
    return {
      isNewElementModalShown: false,
      isDeleteApplicationConfirmationModalShown: false,
      isDeleteTeamConfirmationModalShown: false,
    };
  },
  actions: {
    setIsNewElementModalShown(newValue: boolean): void {
      this.isNewElementModalShown = newValue;
    },
    setIsDeleteApplicationConfirmationModalShown(newValue: boolean): void {
      this.isDeleteApplicationConfirmationModalShown = newValue;
    },
    setIsDeleteTeamConfirmationModalShown(newValue: boolean): void {
      this.isDeleteTeamConfirmationModalShown = newValue;
    },
  },
});
