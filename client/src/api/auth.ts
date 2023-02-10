import { invoke } from "@tauri-apps/api/tauri";
import { deserializeCookie } from "src/helpers";
import * as http from "src/helpers/http";

export interface ILoginInput {
  email: string;
  password: string;
}

export const login = async ({
  email,
  password,
}: ILoginInput): Promise<string> => {
  const response = await http.rawPost<string>({
    url: `api/login`, 
    body: {
      email,
      password
    }
  });
  const cookies = deserializeCookie(response.headers['set-cookie']);
  await invoke('store_tokens', {
    authToken: response.data,
    refreshToken: cookies.ucid
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

export const registerUser = async (
  input: IRegisterUserInput
): Promise<string> => {
  const response = await http.post<string>({
    url: "api/register",
    body: input
  });
  return response;
};

interface ISubmitResetPasswordRequestInput {
  email: string;
}

export const submitResetPasswordRequest = async (
  payload: ISubmitResetPasswordRequestInput
): Promise<void> => {
  await http.post({
    url: "api/reset_password/send_code",
    body: payload
  });
};

export const logoutUser = async (): Promise<void> => {
  await invoke("logout_user");
};

export const fetchAuthToken = async (): Promise<string> => {
  const refreshToken = await invoke<string>("get_stored_refresh_token");
  if (!refreshToken?.length) return "";
  const response = await http
    .rawPost<string>({
      url: "api/refresh_token",
      body: {},
      options: {
        headers: {
          Cookie: `ucid=${refreshToken}`
        }
      }
    });
  await invoke('store_tokens', {
    authToken: response.data ?? "", 
    refreshToken
  });
  return response.data;
};
