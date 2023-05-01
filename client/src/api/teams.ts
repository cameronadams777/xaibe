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

type FetchTeamByIdInput = {
  teamId: string;
};

export const fetchTeamById = async ({
  teamId,
}: FetchTeamByIdInput): Promise<Team> => {
  const response = await http.get<Team>({
    url: `api/teams/${teamId}`,
  });
  return response;
};

type InviteExistingUserToTeamInput = {
  userId: string;
  teamId: string;
};

export const inviteExistingUserToTeam = async ({
  userId,
  teamId,
}: InviteExistingUserToTeamInput): Promise<void> => {
  return http.post({
    url: "api/teams/invites",
    body: {
      user_id: userId,
      team_id: teamId,
    },
  });
};

type RemoveUserFromTeamInput = {
  teamId: string;
  userId: string;
};

export const removeUserFromTeam = async (
  input: RemoveUserFromTeamInput
): Promise<void> => {
  await http.del({ url: `api/teams/${input.teamId}/user/${input.userId}` });
};

type DeleteTeamInput = {
  teamId: string;
};

export const deleteTeam = async ({
  teamId,
}: DeleteTeamInput): Promise<void> => {
  await http.del({ url: `api/teams/${teamId}` });
};

type UpdateTeamSeatCountInput = {
  teamId: string;
  newSeatCount: number;
};

export const updateTeamSeatCount = async ({
  teamId,
  newSeatCount,
}: UpdateTeamSeatCountInput) => {
  const response = await http.patch({
    url: "api/teams/subscription",
    options: {
      body: Body.json({
        teamId,
        newSeatCount,
      }),
    },
  });
  return response;
};

export const fetchPendingTeamInvites = async (): Promise<TeamInvite[]> => {
  const response = await http.get<TeamInvite[]>({
    url: "api/teams/invites",
  });
  return response;
};

export type UpdateInviteStatusInput = {
  inviteId: string;
  status: number;
};

export const updateInviteStatus = async ({
  inviteId,
  status,
}: UpdateInviteStatusInput): Promise<TeamInvite> => {
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
