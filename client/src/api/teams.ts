import { Body } from "@tauri-apps/api/http";
import { NewTeamSubscriptionFormSchema, Team, TeamInvite } from "src/types";
import * as http from "./http";

export const createNewTeam = async (
  form: NewTeamSubscriptionFormSchema
): Promise<{ team: Team; intentId: string; clientSecret: string }> => {
  const response = await http.post<{
    team: Team;
    intentId: string;
    clientSecret: string;
  }>({
    url: `api/teams`,
    body: { ...form },
  });
  return response;
};

interface IFetchTeamByIdInput {
  teamId: string;
}

export const fetchTeamById = async ({
  teamId,
}: IFetchTeamByIdInput): Promise<Team> => {
  const response = await http.get<Team>({
    url: `api/teams/${teamId}`,
  });
  return response;
};

interface IInviteExistingUserToTeamInput {
  userId: string;
  teamId: string;
}

export const inviteExistingUserToTeam = async ({
  userId,
  teamId,
}: IInviteExistingUserToTeamInput): Promise<void> => {
  return http.post({
    url: "api/teams/invites",
    body: {
      user_id: userId,
      team_id: teamId,
    },
  });
};

interface IRemoveUserFromTeamInput {
  teamId: string;
  userId: string;
}

export const removeUserFromTeam = async (
  input: IRemoveUserFromTeamInput
): Promise<void> => {
  await http.del({ url: `api/teams/${input.teamId}/user/${input.userId}` });
};

interface IDeleteTeamInput {
  teamId: string;
}

export const deleteTeam = async ({
  teamId,
}: IDeleteTeamInput): Promise<void> => {
  await http.del({ url: `api/teams/${teamId}` });
};

export const fetchPendingTeamInvites = async (): Promise<TeamInvite[]> => {
  const response = await http.get<TeamInvite[]>({
    url: "api/teams/invites",
  });
  return response;
};

export interface IUpdateInviteStatusInput {
  inviteId: string;
  status: number;
}

export const updateInviteStatus = async ({
  inviteId,
  status,
}: IUpdateInviteStatusInput): Promise<TeamInvite> => {
  const response = await http.patch<TeamInvite>({
    url: "api/teams/invites",
    options: {
      body: Body.json({
        invite_id: inviteId,
        invite_status: status,
      }),
    },
  });
  return response;
};
