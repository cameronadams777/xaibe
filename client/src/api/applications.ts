import { Body } from "@tauri-apps/api/http";
import { camelizeKeys } from "humps";
import { AlertSchema, Application, ApplicationSchema } from "src/types";
import * as http from "./http";

interface IFetchApplicationByIdInput {
  applicationId: string;
}

export const fetchApplicationById = async ({
  applicationId,
}: IFetchApplicationByIdInput): Promise<Application> => {
  const response = await http.get<Application>({
    url: `api/applications/${applicationId}`
  });
  return response;
};

export interface ICreateNewApplicationInput {
  applicationName: string;
  alertSchema?: AlertSchema;
  teamId: string;
  userId: string;
}

export const createNewApplication = async ({
  teamId,
  userId,
  applicationName,
  alertSchema,
}: ICreateNewApplicationInput): Promise<Application> => {
  let body: Record<string, any> = {
    applicationName,
  };
  if (teamId.length) body.teamId = teamId;
  else if (userId.length) body.userId = userId;
  else {
    throw new Error(
      "A teamId or userId must be provided when creating an application."
    );
  }

  if (alertSchema) {
    body.alertSchema = camelizeKeys(alertSchema);
  }

  const response = await http.post<Application>({
    url: "api/applications",
    body
  });
  
  return response;
};

export interface IAddSchemaToApplicationInput {
  applicationId: string;
  title: string;
  description: string;
  link: string;
}

export const addSchemaToApplication = async (
  input: IAddSchemaToApplicationInput
): Promise<Application> => {
  const { applicationId, ...rest } = input;
  const response = await http.patch<Application>({
    url: `api/applications/${applicationId}/alert_schema`,
    options: {
      body: Body.json(rest)
    }
  });
  return response;
};

interface IDeleteApplicationInput {
  applicationId: string;
}

export const deleteApplication = async ({
  applicationId,
}: IDeleteApplicationInput): Promise<void> => {
  await http.del({ url: `api/applications/${applicationId}` });
};
