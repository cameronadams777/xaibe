import * as http from "src/helpers/http";
import { IUser } from "src/types";

export const fetchAllUsers = async (): Promise<IUser[]> => {
  const response = await http.get<IUser[]>({ url: "api/users" });
  return response;
};

interface IInviteNewUserInput {
  email: string;
  teamId?: number;
}

export const inviteNewUser = async ({
  teamId,
  email,
}: IInviteNewUserInput): Promise<void> => {
  await http.post({ url: "api/users/invite", body: { teamId, email } });
};
