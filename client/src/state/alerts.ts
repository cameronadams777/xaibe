import compact from "lodash/compact";
import { defineStore } from "pinia";
import { fetchCachedApplicationAlerts } from "../api/alerts";
import { IAlert } from "../types";

interface IAlertStoreState {
  localCacheAlerts: Record<string, IAlert[]>;
}

interface IGetApplicationAlertsInput {
  applicationId: number;
}

export const useAlertsStore = defineStore("alerts", {
  state: (): IAlertStoreState => {
    return { localCacheAlerts: {} };
  },
  actions: {
    getLocalCacheAlerts(applicationId: number): IAlert[] | undefined {
      return this.localCacheAlerts[`application_${applicationId}`];
    },
    async getCachedApplicationAlerts({
      applicationId,
    }: IGetApplicationAlertsInput): Promise<IAlert[]> {
      let alerts = this.getLocalCacheAlerts(applicationId);
      if (!alerts)
        alerts = await fetchCachedApplicationAlerts({ applicationId });
      return compact(alerts) || [];
    },
  },
});
