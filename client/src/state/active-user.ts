import { defineStore } from "pinia";
import { fetchActiveUser } from "../api/active-user";
import { User } from "../types";

interface IActiveUserState {
  activeUser?: User;
}

export const useActiveUserStore = defineStore("active_user", {
  state: (): IActiveUserState => {
    return { activeUser: undefined };
  },
  actions: {
    async getActiveUser() {
      const user = await fetchActiveUser();
      this.activeUser = user;
    },
  },
});
