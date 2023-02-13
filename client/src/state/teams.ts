import { defineStore } from "pinia";
import {
  fetchPendingTeamInvites,
  IUpdateInviteStatusInput,
  updateInviteStatus as updateTeamInviteStatus,
} from "src/api/teams";
import { Team, TeamInvite } from "src/types";

interface TeamsState {
  pendingTeamInvites: TeamInvite[];
  activeTeam?: Team;
}

export const useTeamsStore = defineStore("teams_store", {
  state(): TeamsState {
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
        (pendingInvite) => pendingInvite.id === inviteId
      );
      this.pendingTeamInvites[inviteIndex] = invite;
    },
    async getPendingTeamInvites() {
      const invites = await fetchPendingTeamInvites();
      this.pendingTeamInvites = invites;
    },
  },
});
