import { defineStore } from "pinia";
import { IApplication } from "../types";
import {
  addSchemaToApplication,
  deleteApplication as sendDeleteApplicationRequest,
  IAddSchemaToApplicationInput,
} from "../api/applications";

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
      this.cachedApplications[application.ID] = application;
    },
    async addSchemaToApplication(
      alertSchema: IAddSchemaToApplicationInput
    ): Promise<void> {
      const application = await addSchemaToApplication(alertSchema);
      this.cachedApplications[application.ID] = application;
    },
    async deleteApplication(applicationId: number): Promise<void> {
      await sendDeleteApplicationRequest({ applicationId });
      delete this.cachedApplications[applicationId];
    },
  },
});
