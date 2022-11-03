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
  const authToken = localStorage.getItem("token");
  const response = await invoke<ICreateNewTeamResponse>(
    TauriEvents.CREATE_NEW_TEAM,
    {
      authToken,
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
  const authToken = localStorage.getItem("token");
  const response = await invoke<IFetchTeamByIdResponse>(
    TauriEvents.FETCH_TEAM_BY_ID,
    {
      authToken,
      teamId,
    }
  );
  return response.data;
};

interface IDeleteTeamInput {
  teamId: number;
}

export const deleteTeam = async ({
  teamId,
}: IDeleteTeamInput): Promise<void> => {
  const authToken = localStorage.getItem("token");
  let body: Record<string, any> = {
    teamId,
    authToken,
  };
  await invoke<ITeam>(TauriEvents.DELETE_TEAM, {
    authToken,
    teamId,
  });
};
