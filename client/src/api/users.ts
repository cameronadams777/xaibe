import { invoke } from "@tauri-apps/api";
import { IUser } from "src/types";
import { TauriEvents } from ".";

interface IFetchAllUsersResponse {
  status: string;
  message: string;
  data: IUser[];
}

export const fetchAllUsers = async (): Promise<IUser[]> => {
  const response = await invoke<IFetchAllUsersResponse>(
    TauriEvents.FETCH_ALL_USERS
  );
  return response.data;
};
