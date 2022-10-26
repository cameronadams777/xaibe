import { invoke } from "@tauri-apps/api";
import { TauriEvents } from ".";
import { IUser } from "../types";

export const fetchActiveUser = async (): Promise<IUser | undefined> => {
  try {
    const authToken = localStorage.getItem("token");
    const responseString = await invoke<string>(TauriEvents.FETCH_ACTIVE_USER, {
      authToken,
    });
    const response = JSON.parse(responseString);
    console.log(response);
    return response.data.user;
  } catch (error) {
    console.error(
      "Galata Error: An error occurred while retrieving active user:",
      error
    );
    return undefined;
  }
};
