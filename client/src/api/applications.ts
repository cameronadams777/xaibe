import * as http from "src/helpers/http";
import { Body } from "@tauri-apps/api/http";
import { IAlertSchema, IApplication } from "src/types";
import { camelizeKeys } from "humps";

interface IFetchApplicationByIdInput {
  applicationId: number;
}

export const fetchApplicationById = async ({
  applicationId,
}: IFetchApplicationByIdInput): Promise<IApplication> => {
  const response = await http.get<IApplication>({
    url: `api/applications/${applicationId}`
  })
  return response;
};

export interface ICreateNewApplicationInput {
  applicationName: string;
  alertSchema?: IAlertSchema;
  teamId?: number;
  userId?: number;
}

export const createNewApplication = async ({
  teamId,
  userId,
  applicationName,
  alertSchema,
}: ICreateNewApplicationInput): Promise<IApplication> => {
  let body: Record<string, any> = {
    applicationName,
  };
  if (teamId != null) body.teamId = teamId;
  else if (userId != null) body.userId = userId;
  else {
    throw new Error(
      "A teamId or userId must be provided when creating an application."
    );
  }

  if (alertSchema) {
    body.alertSchema = camelizeKeys(alertSchema);
  }

  const response = await http.post<IApplication>({
    url: "api/applications",
    body
  });
  
  return response;
};

export interface IAddSchemaToApplicationInput {
  applicationId: number;
  title: string;
  description: string;
  link: string;
}

export const addSchemaToApplication = async (
  input: IAddSchemaToApplicationInput
): Promise<IApplication> => {
  const { applicationId, ...rest } = input;
  const response = await http.patch<IApplication>({
    url: `api/applications/${applicationId}/alert_schema`,
    options: {
      body: Body.json(rest)
    }
  });
  return response;
};

interface IDeleteApplicationInput {
  applicationId: number;
}

export const deleteApplication = async ({
  applicationId,
}: IDeleteApplicationInput): Promise<void> => {
  await http.del({ url: `api/applications/${applicationId}` });
};
