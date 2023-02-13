import { defineStore } from "pinia";
import { Application } from "../types";
import {
  addSchemaToApplication,
  deleteApplication as sendDeleteApplicationRequest,
  IAddSchemaToApplicationInput,
} from "../api/applications";

interface ApplicationsState {
  cachedApplications: Record<string, Application>;
}

export const useApplicationsStore = defineStore("applications", {
  state: (): ApplicationsState => {
    return { cachedApplications: {} };
  },
  actions: {
    getCachedApplication(applicationId: string): Application | undefined {
      return this.cachedApplications[applicationId];
    },
    cacheApplication(application: Application): void {
      this.cachedApplications[application.id] = application;
    },
    async addSchemaToApplication(
      alertSchema: IAddSchemaToApplicationInput
    ): Promise<void> {
      const application = await addSchemaToApplication(alertSchema);
      this.cachedApplications[application.id] = application;
    },
    async deleteApplication(applicationId: string): Promise<void> {
      await sendDeleteApplicationRequest({ applicationId });
      delete this.cachedApplications[applicationId];
    },
  },
});
