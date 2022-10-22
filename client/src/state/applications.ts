import { defineStore } from "pinia";
import { IApplication } from "../types";

interface IApplicationsState {
  cachedApplications: Record<number, IApplication>;
}

export const useApplicationsStore = defineStore("applications", {
  state: (): IApplicationsState => {
    return { cachedApplications: {} };
  },
  actions: {
    getCachedApplication(applicationId: number): IApplication | undefined {
      return this.cachedApplications[applicationId];
    },
    cacheApplication(application: IApplication): void {
      this.cachedApplications[application.id] = application;
    },
  },
});
