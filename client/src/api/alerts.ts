import { invoke } from "@tauri-apps/api";
import { TauriEvents } from ".";
import { IAlert } from "../types";

interface IFetchCachedAlertsInput {
  applicationId: number;
}

interface IFetchCacheAlertsResponse {
  status: string;
  message: string;
  data: IAlert[] | undefined;
}

export const fetchCachedApplicationAlerts = async ({
  applicationId,
}: IFetchCachedAlertsInput): Promise<IAlert[] | undefined> => {
  const authToken = localStorage.getItem("token");
  const response = await invoke<IFetchCacheAlertsResponse>(
    TauriEvents.FETCH_CACHED_ALERTS,
    {
      authToken,
      applicationId,
    }
  );
  return response.data;
};
