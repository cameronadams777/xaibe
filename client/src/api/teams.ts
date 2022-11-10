import { invoke } from "@tauri-apps/api";
import { TauriEvents } from ".";
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
  const response = await invoke<ICreateNewTeamResponse>(
    TauriEvents.CREATE_NEW_TEAM,
    {
      teamName,
    }
  );
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
  const responseString = await invoke<string>(TauriEvents.FETCH_TEAM_BY_ID, {
    teamId,
  });
  const response: IFetchTeamByIdResponse = JSON.parse(responseString);
  return response.data;
};

interface IAddUserToTeamInput {
  teamId: number;
  userId: number;
}

export const addUserToTeam = async (
  input: IAddUserToTeamInput
): Promise<void> => {
  await invoke<string>(TauriEvents.ADD_USER_TO_TEAM, { ...input });
};

interface IRemoveUserFromTeamInput {
  teamId: number;
  userId: number;
}

export const removeUserFromTeam = async (
  input: IRemoveUserFromTeamInput
): Promise<void> => {
  await invoke<string>(TauriEvents.REMOVE_USER_FROM_TEAM, { ...input });
};

interface IDeleteTeamInput {
  teamId: number;
}

export const deleteTeam = async ({
  teamId,
}: IDeleteTeamInput): Promise<void> => {
  let body: Record<string, any> = {
    teamId,
  };
  await invoke<ITeam>(TauriEvents.DELETE_TEAM, {
    teamId,
  });
};
