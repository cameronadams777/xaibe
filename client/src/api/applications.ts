import { invoke } from "@tauri-apps/api";
import { TauriEvents } from ".";
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
  const response = await invoke<IFetchApplicationByIdResponse>(
    TauriEvents.FETCH_APPLICATION_BY_ID,
    {
      applicationId,
    }
  );
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

  const response = await invoke<ICreateNewApplicationResponse>(
    TauriEvents.CREATE_NEW_APPLICATION,
    body
  );
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
  alertSchema: IAddSchemaToApplicationInput
): Promise<IApplication> => {
  const response = await invoke<IAddSchemaToApplicationResponse>(
    TauriEvents.ADD_SCHEMA_TO_APPLICATION,
    { ...alertSchema }
  );
  return response.data;
};

interface IDeleteApplicationInput {
  applicationId: number;
}

export const deleteApplication = async ({
  applicationId,
}: IDeleteApplicationInput): Promise<void> => {
  let body: Record<string, any> = {
    applicationId,
  };
  await invoke<IApplication>(TauriEvents.DELETE_APPLICATION, {
    ...body,
  });
};
