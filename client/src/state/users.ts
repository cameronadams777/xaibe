import { defineStore } from "pinia";
import { fetchAllUsers } from "src/api/users";
import { User } from "src/types";

interface IGalataUsersState {
  users: User[];
}

export const useGalataUsersStore = defineStore("galata_users", {
  state(): IGalataUsersState {
    return { users: [] };
  },
  actions: {
    async getAllUsers(): Promise<User[]> {
      const users = await fetchAllUsers();
      this.users = users ?? [];
      return users ?? [];
    },
  },
});
