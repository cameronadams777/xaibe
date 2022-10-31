import { defineStore } from "pinia";
import { IToast, ToastType } from "../types";

interface IToastStoreState {
  activeToast: IToast | undefined;
}

export const useToastStore = defineStore("toast", {
  state: (): IToastStoreState => {
    return {
      activeToast: undefined,
    };
  },
  getters: {
    toastTitleByType(): string {
      if (this.activeToast?.type === ToastType.INFO) return "Info";
      else if (this.activeToast?.type === ToastType.SUCCESS) return "Success";
      else if (this.activeToast?.type === ToastType.ERROR) return "Error";
      return "";
    },
  },
  actions: {
    setActiveToast(newValue: IToast): void {
      this.activeToast = newValue;
      setTimeout(() => (this.activeToast = undefined), 2000);
    },
  },
});
