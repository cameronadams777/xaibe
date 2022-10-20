import { invoke } from "@tauri-apps/api";

interface IFetchCachedAlertsInput {
  applicationId: number;
  serviceToken: string;
}

interface IFetchCacheAlertsResponse {
  status: string;
  message: string;
  data: Record<string, any>;
}

export const fetchCachedAlerts = async ({
  applicationId,
  serviceToken,
}: IFetchCachedAlertsInput): Promise<Record<string, any>> => {
  try {
    const authToken = localStorage.getItem("token");
    const responseString = await invoke<string>("fetch_cached_alerts", {
      authToken,
      applicationId,
      serviceToken,
    });
    const response: IFetchCacheAlertsResponse = JSON.parse(responseString);
    return response.data.alerts;
  } catch (error) {
    console.error("An error occurred while retrieving cached alerts.");
    return {};
  }
};
