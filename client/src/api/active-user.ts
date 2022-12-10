import * as http from "src/helpers/http";
import { IUser } from "src/types";

interface IFetchActiveUserResponse {
  status: string;
  message: string;
  data: IUser;
}

export const fetchActiveUser = async (): Promise<IUser | undefined> => {
  const response = await http.get<IFetchActiveUserResponse>({ url: "api/users/me" });
  return response.data;
};
