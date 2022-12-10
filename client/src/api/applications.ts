import * as http from "src/helpers/http";
import { Body } from "@tauri-apps/api/http";
import { IAlertSchema, IApplication } from "src/types";
import { camelizeKeys } from "humps";

interface IFetchApplicationByIdInput {
  applicationId: number;
}

interface IFetchApplicationByIdResponse {
  status: string;
  message: string;
  data: IApplication;
}

export const fetchApplicationById = async ({
  applicationId,
}: IFetchApplicationByIdInput): Promise<IApplication | undefined> => {
  const response = await http.get<IFetchApplicationByIdResponse>({
    url: `api/applications/${applicationId}`
  })
  return response.data;
};

export interface ICreateNewApplicationInput {
  applicationName: string;
  alertSchema?: IAlertSchema;
  teamId?: number;
  userId?: number;
}

interface ICreateNewApplicationResponse {
  status: string;
  message: string;
  data?: IApplication;
}

export const createNewApplication = async ({
  teamId,
  userId,
  applicationName,
  alertSchema,
}: ICreateNewApplicationInput): Promise<IApplication | undefined> => {
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

  const response = await http.post<ICreateNewApplicationResponse>({
    url: "api/application",
    body
  });
  
  return response.data;
};

export interface IAddSchemaToApplicationInput {
  applicationId: number;
  title: string;
  description: string;
  link: string;
}

interface IAddSchemaToApplicationResponse {
  status: string;
  message: string;
  data: IApplication;
}

export const addSchemaToApplication = async (
  input: IAddSchemaToApplicationInput
): Promise<IApplication> => {
  const { applicationId, ...rest } = input;
  const response = await http.patch<IAddSchemaToApplicationResponse>({
    url: `api/applications/${applicationId}`,
    options: {
      body: Body.json(rest)
    }
  });
  return response.data;
};

interface IDeleteApplicationInput {
  applicationId: number;
}

export const deleteApplication = async ({
  applicationId,
}: IDeleteApplicationInput): Promise<void> => {
  await http.del({ url: `api/applications/${applicationId}` });
};
