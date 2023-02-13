import { z } from "zod";
import { User, UserSchema } from "src/types";
import * as http from "./http";

const FetchAllUsersResponseSchema = z.lazy(() => z.array(UserSchema));

export const fetchAllUsers = async (): Promise<User[]> => {
  const response = await http.get<User[]>({ url: "api/users" });
  FetchAllUsersResponseSchema.parse(response);
  return response;
};

interface IInviteNewUserInput {
  email: string;
  teamId?: string;
}

export const inviteNewUser = async ({
  teamId,
  email,
}: IInviteNewUserInput): Promise<void> => {
  await http.post({ url: "api/users/invite", body: { teamId, email } });
};
