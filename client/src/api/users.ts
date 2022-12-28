import * as http from "src/helpers/http";
import { IUser } from "src/types";

interface IFetchAllUsersResponse {
  status: string;
  message: string;
  data: IUser[];
}

export const fetchAllUsers = async (): Promise<IUser[]> => {
  const response = await http.get<IFetchAllUsersResponse>({ url: "api/users" });
  return response.data;
};

interface IInviteNewUserInput {
  email: string;
}

export const inviteNewUser = async ({
  email,
}: IInviteNewUserInput): Promise<void> => {
  await http.post({ url: "api/users/invite", body: { email } });
};
