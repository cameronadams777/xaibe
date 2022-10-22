import { invoke } from "@tauri-apps/api";
import { TauriEvents } from ".";
import { IApplication } from "../types";

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
  const authToken = localStorage.getItem("token");
  const response = await invoke<IFetchApplicationByIdResponse>(
    TauriEvents.FETCH_APPLICATION_BY_ID,
    {
      authToken,
      applicationId,
    }
  );
  return response.data;
};

interface ICreateNewApplicationInput {
  applicationName: string;
  teamId?: number;
  userId?: number;
}

export const createNewApplication = async ({
  teamId,
  userId,
  applicationName,
}: ICreateNewApplicationInput): Promise<IApplication | undefined> => {
  const authToken = localStorage.getItem("token");
  let body: Record<string, any> = {
    applicationName,
    authToken,
  };
  if (teamId != null) body.teamId = teamId;
  else if (teamId != null) body.userId = userId;
  else {
    throw new Error(
      "A teamId or userId must be provided when creating an application."
    );
  }
  const application = await invoke<IApplication>(
    TauriEvents.CREATE_NEW_APPLICATION,
    {
      ...body,
    }
  );
  return application;
};
