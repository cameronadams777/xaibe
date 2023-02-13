import { Body } from "@tauri-apps/api/http";
import { Team, TeamInvite, TeamInviteSchema, TeamSchema } from "src/types";
import * as http from "./http";

interface ICreateNewTeamInput {
  teamName: string;
}

export const createNewTeam = async ({
  teamName,
}: ICreateNewTeamInput): Promise<Team> => {
  const response = await http.post<Team>({
    url: `api/teams`,
    body: { teamName },
  });
  TeamSchema.parse(response);
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
  TeamSchema.parse(response);
  return response;
};

interface IInviteExistingUserToTeamInput {
  userId: string;
  teamId: string;
}

export const inviteExistingUserToTeam = async ({ userId, teamId }: IInviteExistingUserToTeamInput): Promise<void> => {
  return http.post({ 
    url: "api/teams/invites", 
    body: { 
      user_id: userId, 
      team_id: teamId
    } 
  });
}

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
  TeamInviteSchema.parse(response);
  return response;
};

export interface IUpdateInviteStatusInput {
  inviteId: number;
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
  TeamInviteSchema.parse(response);
  return response;
};
