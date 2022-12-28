import { defineStore } from "pinia";
import {
  fetchPendingTeamInvites,
  IUpdateInviteStatusInput,
  updateInviteStatus as updateTeamInviteStatus,
} from "src/api/teams";
import { ITeam, ITeamInvite } from "src/types";

interface ITeamsState {
  pendingTeamInvites: ITeamInvite[];
  activeTeam?: ITeam;
}

export const useTeamsStore = defineStore("teams_store", {
  state(): ITeamsState {
    return {
      activeTeam: undefined,
      pendingTeamInvites: [],
    };
  },
  actions: {
    async updateInviteStatus({
      inviteId,
      status,
    }: IUpdateInviteStatusInput): Promise<void> {
      const invite = await updateTeamInviteStatus({ inviteId, status });
      if (!invite) throw new Error("Galata Error: Invite not found.");
      const inviteIndex = this.pendingTeamInvites.findIndex(
        (pendingInvite) => pendingInvite.ID === inviteId
      );
      this.pendingTeamInvites[inviteIndex] = invite;
    },
    async getPendingTeamInvites() {
      const invites = await fetchPendingTeamInvites();
      this.pendingTeamInvites = invites;
    },
  },
});
