import { invoke } from "@tauri-apps/api";
import { IApplication, IUser } from "../types";

export const fetchActiveUser = async (): Promise<IUser | undefined> => {
  try {
    const authToken = localStorage.getItem("token");
    const responseString = await invoke<string>("fetch_active_user", {
      authToken,
    });
    const response = JSON.parse(responseString);
    return response.data.user;
  } catch (error) {
    console.error("An error occurred while retrieving active user.");
    return undefined;
  }
};
