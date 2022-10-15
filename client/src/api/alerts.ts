import { invoke } from "@tauri-apps/api";

interface IFetchCachedAlertsInput {
  applicationId: number;
  serviceToken: string;
}

export const fetchCachedAlerts = async ({
  applicationId,
  serviceToken,
}: IFetchCachedAlertsInput): Promise<Record<string, any>> => {
  try {
    const authToken = localStorage.getItem("token");
    const alerts = await invoke<Record<string, any>>("fetch_cached_alerts", {
      authToken,
      applicationId,
      serviceToken,
    });
    return alerts;
  } catch (error) {
    console.error("An error occurred while retrieving cached alerts.");
    return {};
  }
};
