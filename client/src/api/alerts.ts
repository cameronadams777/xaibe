import * as http from "src/helpers/http";
import { IAlert, IAlertSchema } from "src/types";

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
  const response = await http.get<IFetchCacheAlertsResponse>({ url: "api/alerts" })
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
  const response = await http.get<IFetchCacheAlertsByApplicationResponse>({
    url: `api/alerts/applications/${applicationId}`
  })
  return response.data || [];
};
