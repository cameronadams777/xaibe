import * as http from "src/helpers/http";
import { IAlert, IAlertSchema } from "src/types";

export interface ICachedAlerts {
  AlertSchema: IAlertSchema;
  Alerts: IAlert[];
}

export const fetchAllCachedAlerts = async (): Promise<{
  string: ICachedAlerts;
}> => {
  const response = await http.get<{string: ICachedAlerts}>({ url: "api/alerts" })
  return response;
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
