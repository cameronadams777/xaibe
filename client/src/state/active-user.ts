import { defineStore } from "pinia";
import { fetchActiveUser } from "../api/active-user";
import { IUser } from "../types";

interface IActiveUserState {
  activeUser?: IUser;
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
