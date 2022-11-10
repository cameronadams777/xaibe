import { defineStore } from "pinia";
import {
  fetchAuthToken,
  ILoginInput,
  IRegisterUserInput,
  login,
  registerUser,
} from "src/api/auth";

interface IAuthStoreState {
  token: string;
}

export const useAuthStore = defineStore("auth", {
  state: (): IAuthStoreState => {
    return { token: "" };
  },
  actions: {
    async login(payload: ILoginInput): Promise<void> {
      const token = await login(payload);
      this.token = token;
    },
    async registerUser(payload: IRegisterUserInput): Promise<void> {
      const token = await registerUser(payload);
      this.token = token;
    },
    async fetchAuthToken(): Promise<string> {
      const token = await fetchAuthToken();
      this.token = token;
      return token;
    },
  },
});
