import { invoke } from "@tauri-apps/api";
import { TauriEvents } from ".";

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
}: IFetchCachedAlertsInput): Promise<Record<string, any>> => {
  try {
    const authToken = localStorage.getItem("token");
    const responseString = await invoke<string>(
      TauriEvents.FETCH_CACHED_ALERTS,
      {
        authToken,
        applicationId,
      }
    );
    const response: IFetchCacheAlertsResponse = JSON.parse(responseString);
    return response.data.alerts;
  } catch (error) {
    console.error("An error occurred while retrieving cached alerts.");
    return {};
  }
};
