import { invoke } from "@tauri-apps/api";
import { TauriEvents } from ".";
import { ITeam } from "../types";

interface ICreateNewTeamInput {
  teamName: string;
}

export const createNewTeam = async ({
  teamName,
}: ICreateNewTeamInput): Promise<ITeam | undefined> => {
  const authToken = localStorage.getItem("token");
  const newTeam = await invoke<ITeam>(TauriEvents.CREATE_NEW_TEAM, {
    authToken,
    teamName,
  });
  return newTeam;
};
