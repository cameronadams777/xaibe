import type { IAlert, AlertSchema } from "src/types";
import * as http from "./http";

export interface ICachedAlerts {
  AlertSchema: AlertSchema;
  Alerts: IAlert[];
}

export const fetchAllCachedAlerts = async (): Promise<{
  string: ICachedAlerts;
}> => {
  const response = await http.get<{string: ICachedAlerts}>({ url: "api/alerts" });
  // TODO: Add validation to request
  return response;
};

interface IFetchCachedAlertsInput {
  applicationId: string;
}

export const fetchCachedApplicationAlerts = async ({
  applicationId,
}: IFetchCachedAlertsInput): Promise<IAlert[]> => {
  const response = await http.get<IAlert[]>({
    url: `api/alerts/applications/${applicationId}`
  });
  // TODO: Add validation to request
  return response || [];
};
