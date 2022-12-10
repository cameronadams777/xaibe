import { invoke } from "@tauri-apps/api/tauri";
import * as http from "src/helpers/http";

export interface ILoginInput {
  email: string;
  password: string;
}

interface ILoginResponse {
  status: string;
  message: string;
  data: string;
}

export const login = async ({
  email,
  password,
}: ILoginInput): Promise<string> => {
  const response = await http.post<ILoginResponse>({
    url: `api/login`, 
    body: {
      email,
      password
    }
  });
  return response.data;
};

export interface IRegisterUserInput {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  passwordConfirmation: string;
}

interface IRegisterUserResponse {
  status: string;
  message: string;
  data: string;
}

export const registerUser = async (
  input: IRegisterUserInput
): Promise<string> => {
  const response = await http.post<IRegisterUserResponse>({
    url: "api/register",
    body: input
  });
  return response.data;
};

interface ISubmitResetPasswordRequestInput {
  email: string;
}

export const submitResetPasswordRequest = async (
  payload: ISubmitResetPasswordRequestInput
): Promise<void> => {
  await http.post({
    url: "api/reset-password/send-code",
    body: payload
  });
};

export const logoutUser = async (): Promise<void> => {
  await invoke("logout_user");
};

interface IFetchAuthTokenResponse {
  token: string;
}

export const fetchAuthToken = async (): Promise<string> => {
  const response = await http.get<IFetchAuthTokenResponse>({
    url: "api/refresh_token",
  });
  await invoke('store_auth_token', { token: response.token });
  return response.token;
};
