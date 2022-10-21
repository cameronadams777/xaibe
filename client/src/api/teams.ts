import { invoke } from "@tauri-apps/api";
import { ITeam } from "../types";

interface ICreateNewTeamInput {
  teamName: string;
}

export const createNewTeam = async ({
  teamName,
}: ICreateNewTeamInput): Promise<ITeam | undefined> => {
  try {
    const newTeam = await invoke<ITeam>("create_new_team", { teamName });
    return newTeam;
  } catch (error) {
    // TODO: Handle with error toast as we need one of these provided
    console.error(
      "Galata Error: An unknown error occurred when trying to create a new team."
    );
  }
};
