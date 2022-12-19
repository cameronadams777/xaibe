import { invoke } from "@tauri-apps/api/tauri";
import { deserializeCookie } from "src/helpers";
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
  const response = await http.rawPost<ILoginResponse>({
    url: `api/login`, 
    body: {
      email,
      password
    }
  });
  const cookies = deserializeCookie(response.headers['set-cookie']);
  await invoke('store_tokens', {
    authToken: response.data.data,
    refreshToken: cookies.ucid
  });
  return response.data.data;
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
    url: "api/reset_password/send_code",
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
  const refreshToken = await invoke<string>("get_stored_refresh_token");
  if (!refreshToken?.length) return "";
  const response = await http
    .rawPost<IFetchAuthTokenResponse>({
      url: "api/refresh_token",
      body: {},
      options: {
        headers: {
          Cookie: `ucid=${refreshToken}`
        }
      }
    });
  await invoke('store_tokens', {
    authToken: response.data.token ?? "", 
    refreshToken
  });
  return response.data.token;
};
