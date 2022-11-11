import { invoke } from "@tauri-apps/api";
import { IUser } from "src/types";
import { TauriEvents } from ".";

interface IFetchAllUsersResponse {
  status: string;
  message: string;
  data: IUser[];
}

export const fetchAllUsers = async (): Promise<IUser[]> => {
  const responseString = await invoke<string>(TauriEvents.FETCH_ALL_USERS);
  const response: IFetchAllUsersResponse = JSON.parse(responseString);
  return response.data;
};
