import { defineStore } from "pinia";
import { fetchAllUsers } from "src/api/users";
import { User } from "src/types";

interface IXaibeUsersState {
  users: User[];
}

export const useXaibeUsersStore = defineStore("xaibe_users", {
  state(): IXaibeUsersState {
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
