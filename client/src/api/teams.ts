import { Body } from "@tauri-apps/api/http";
import * as http from "src/helpers/http";
import { InviteStatus, ITeam, ITeamInvite } from "src/types";

interface ICreateNewTeamInput {
  teamName: string;
}

export const createNewTeam = async ({
  teamName,
}: ICreateNewTeamInput): Promise<ITeam> => {
  const response = await http.post<ITeam>({
    url: `api/teams`,
    body: { teamName },
  });
  return response;
};

interface IFetchTeamByIdInput {
  teamId: number;
}

export const fetchTeamById = async ({
  teamId,
}: IFetchTeamByIdInput): Promise<ITeam> => {
  const response = await http.get<ITeam>({
    url: `api/teams/${teamId}`,
  });
  return response;
};

interface IInviteExistingUserToTeamInput {
  userId: number;
  teamId: number;
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
  teamId: number;
  userId: number;
}

export const removeUserFromTeam = async (
  input: IRemoveUserFromTeamInput
): Promise<void> => {
  await http.del({ url: `api/teams/${input.teamId}/user/${input.userId}` });
};

interface IDeleteTeamInput {
  teamId: number;
}

export const deleteTeam = async ({
  teamId,
}: IDeleteTeamInput): Promise<void> => {
  await http.del({ url: `api/teams/${teamId}` });
};

interface IFetchPendingTeamInvitesResponse {
  status: string;
  message: string;
  data: ITeamInvite[];
}

export const fetchPendingTeamInvites = async (): Promise<ITeamInvite[]> => {
  const response = await http.get<IFetchPendingTeamInvitesResponse>({
    url: "api/teams/invites",
  });
  return response.data;
};

export interface IUpdateInviteStatusInput {
  inviteId: number;
  status: InviteStatus;
}

export const updateInviteStatus = async ({
  inviteId,
  status,
}: IUpdateInviteStatusInput): Promise<ITeamInvite> => {
  const response = await http.patch<ITeamInvite>({
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
