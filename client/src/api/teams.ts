import * as http from "src/helpers/http";
import { ITeam } from "src/types";

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
    body: { teamName } 
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
  const response = await http.get<IFetchTeamByIdResponse>({ url: `api/teams/${teamId}` });
  return response.data;
};

interface IAddUserToTeamInput {
  teamId: number;
  userId: number;
}

export const addUserToTeam = async (
  input: IAddUserToTeamInput
): Promise<void> => {
  await http.post({ 
    url: `api/teams/${input.teamId}/user/${input.userId}`,
    body: {}
  });
};

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
  await http.del({ url: `api/teams/${teamId}` })
};
