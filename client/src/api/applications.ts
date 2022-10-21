import { invoke } from "@tauri-apps/api";
import { IApplication } from "../types";

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
  try {
    let body: Record<string, any> = {
      applicationName,
    };
    if (teamId != null) body.teamId = teamId;
    else if (teamId != null) body.userId = userId;
    else {
      // TODO: Handle with error toast as we need one of these provided
      console.error(
        "Galata Error: A teamId or userId must be provided when creating an application."
      );
      return;
    }
    const application = await invoke<IApplication>("create_new_application", {
      ...body,
    });
    return application;
  } catch (error) {
    // TODO: Handle with error toast as we need one of these provided
    console.error(
      "Galata Error: An unknown error occurred when trying to create a new application."
    );
  }
};
