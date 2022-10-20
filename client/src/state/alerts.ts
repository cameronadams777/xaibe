import { defineStore } from "pinia";
import { fetchCachedAlerts } from "../api/alerts";

interface IGetAlertsInput {
  applicationId: number;
  serviceToken: string;
}

export const useAlertsStore = defineStore("alerts", {
  state: () => {
    return { alerts: {} };
  },
  actions: {
    async getAlerts({ applicationId, serviceToken }: IGetAlertsInput) {
      const fetchedAlerts = await fetchCachedAlerts({
        applicationId,
        serviceToken,
      });
      this.alerts = fetchedAlerts;
    },
  },
});
