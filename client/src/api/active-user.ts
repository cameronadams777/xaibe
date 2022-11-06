import { invoke } from "@tauri-apps/api";
import { TauriEvents } from ".";
import { IUser } from "src/types";

interface IFetchActiveUserResponse {
  status: string;
  message: string;
  data: IUser;
}

export const fetchActiveUser = async (): Promise<IUser | undefined> => {
  const authToken = localStorage.getItem("token");
  const response = await invoke<IFetchActiveUserResponse>(
    TauriEvents.FETCH_ACTIVE_USER,
    {
      authToken,
    }
  );
  return response.data;
};
