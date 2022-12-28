import { Body } from "@tauri-apps/api/http";
import * as http from "src/helpers/http";
import { InviteStatus, ITeam, ITeamInvite } from "src/types";

interface ICreateNewTeamInput {
  teamName: string;
}

interface ICreateNewTeamResponse {
  status: string;
  message: string;
  data?: ITeam;
}

export const createNewTeam = async ({
  teamName,
}: ICreateNewTeamInput): Promise<ITeam | undefined> => {
  const response = await http.post<ICreateNewTeamResponse>({
    url: `api/teams`,
    body: { teamName },
  });
  return response.data;
};

interface IFetchTeamByIdInput {
  teamId: number;
}

interface IFetchTeamByIdResponse {
  status: string;
  message: string;
  data?: ITeam;
}

export const fetchTeamById = async ({
  teamId,
}: IFetchTeamByIdInput): Promise<ITeam | undefined> => {
  const response = await http.get<IFetchTeamByIdResponse>({
    url: `api/teams/${teamId}`,
  });
  return response.data;
};

interface IInviteExistingUserToTeamInput {
  userId: number;
  teamId: number;
}

export const inviteExistingUserToTeam = async ({ userId, teamId }: IInviteExistingUserToTeamInput): Promise<void> => {
  console.log(userId, teamId)
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

interface IUpdateInviteStatusResponse {
  status: string;
  message: string;
  data?: ITeamInvite;
}

export const updateInviteStatus = async ({
  inviteId,
  status,
}: IUpdateInviteStatusInput): Promise<ITeamInvite | undefined> => {
  const response = await http.patch<IUpdateInviteStatusResponse>({
    url: "api/teams/invites",
    options: {
      body: Body.json({
        invite_id: inviteId,
        invite_status: status,
      }),
    },
  });
  return response.data;
};
