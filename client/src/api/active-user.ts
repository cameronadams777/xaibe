import { invoke } from "@tauri-apps/api";
import { TauriEvents } from ".";
import { IUser } from "../types";

export const fetchActiveUser = async (): Promise<IUser | undefined> => {
  const authToken = localStorage.getItem("token");
  const responseString = await invoke<string>(TauriEvents.FETCH_ACTIVE_USER, {
    authToken,
  });
  const response = JSON.parse(responseString);
  return response.data.user;
};
