import { defineStore } from "pinia";
import { fetchAllUsers } from "src/api/users";
import { IUser } from "src/types";

interface IGalataUsersState {
  users: IUser[];
}

export const useGalataUsersStore = defineStore("galata_users", {
  state(): IGalataUsersState {
    return { users: [] };
  },
  actions: {
    async getAllUsers() {
      const users = await fetchAllUsers();
      this.users = users;
    },
  },
});
