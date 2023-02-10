import * as http from "src/helpers/http";
import { IUser } from "src/types";

export const fetchActiveUser = async (): Promise<IUser> => {
  const response = await http.get<IUser>({ url: "api/users/me" });
  return response;
};
