import { invoke } from "@tauri-apps/api";

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
  const response = await invoke<ILoginResponse>("login", { email, password });
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
  const response = await invoke<IRegisterUserResponse>("register_user", {
    ...input,
  });
  return response.data;
};

interface IFetchAuthTokenResponse {
  token: string;
}

export const fetchAuthToken = async (): Promise<string> => {
  const response = await invoke<IFetchAuthTokenResponse>("fetch_auth_token");
  console.log(response);
  return response.token;
};
