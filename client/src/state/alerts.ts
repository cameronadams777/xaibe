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
    pushAlertToApplication(applicationId: number, alert: IAlert) {
      if(!this.localCacheAlerts[`application_${applicationId}`]) {
        this.localCacheAlerts[`application_${applicationId}`] = [alert];
        console.log(this.localCacheAlerts);
        return;
      }
      this.localCacheAlerts[`application_${applicationId}`].push(alert);
    },
    async getCachedApplicationAlerts({
      applicationId,
    }: IGetApplicationAlertsInput): Promise<IAlert[]> {
      const alerts = await fetchCachedApplicationAlerts({ applicationId });
      this.localCacheAlerts[`application_${applicationId}`] = compact(alerts) ?? [];
      return compact(alerts) || [];
    },
  },
});
