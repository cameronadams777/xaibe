import { defineStore } from "pinia";
import { fetchActiveUser } from "../api/active-user";
import { IUser } from "../types";

interface ActiveUserState {
  activeUser?: IUser;
}

export const useActiveUserStore = defineStore("active_user", {
  state: (): ActiveUserState => {
    return { activeUser: undefined };
  },
  actions: {
    async getActiveUser() {
      const user = await fetchActiveUser();
      this.activeUser = user;
    },
  },
});
