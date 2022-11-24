import { invoke } from "@tauri-apps/api";
import { TauriEvents } from ".";
import { IAlert, IAlertSchema } from "src/types";
import { camelizeKeys } from "humps";

export interface ICachedAlerts {
  AlertSchema: IAlertSchema;
  Alerts: IAlert[];
}

interface IFetchCacheAlertsResponse {
  status: string;
  message: string;
  data: {
    string: ICachedAlerts;
  };
}

export const fetchAllCachedAlerts = async (): Promise<{
  string: ICachedAlerts;
}> => {
  const responseString = await invoke<string>(
    TauriEvents.FETCH_ALL_CACHED_ALERTS
  );
  const response: IFetchCacheAlertsResponse = JSON.parse(responseString);
  return response.data;
};

interface IFetchCachedAlertsInput {
  applicationId: number;
}

interface IFetchCacheAlertsByApplicationResponse {
  status: string;
  message: string;
  data: IAlert[] | undefined;
}

export const fetchCachedApplicationAlerts = async ({
  applicationId,
}: IFetchCachedAlertsInput): Promise<IAlert[]> => {
  const response = await invoke<IFetchCacheAlertsByApplicationResponse>(
    TauriEvents.FETCH_CACHED_ALERTS_BY_APPLICATION,
    {
      applicationId,
    }
  );
  return response.data || [];
};
