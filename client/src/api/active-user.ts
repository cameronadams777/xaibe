import { User, UserSchema } from "src/types";
import * as http from "./http";

export const fetchActiveUser = async (): Promise<User> => {
  const response = await http.get<User>({ url: "api/users/me" });
  UserSchema.parse(response);
  return response;
};
