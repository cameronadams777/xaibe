import { invoke } from "@tauri-apps/api";
import { TauriEvents } from ".";

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
  const response = await invoke<ILoginResponse>(TauriEvents.LOGIN, {
    email,
    password,
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
  const response = await invoke<IRegisterUserResponse>(
    TauriEvents.REGISTER_USER,
    {
      ...input,
    }
  );
  return response.data;
};

interface ISubmitResetPasswordRequestInput {
  email: string;
}

export const submitResetPasswordRequest = async (
  payload: ISubmitResetPasswordRequestInput
): Promise<void> => {
  await invoke("submit_reset_password_request", { ...payload });
};

export const logoutUser = async (): Promise<void> => {
  await invoke(TauriEvents.LOGOUT_USER);
};

interface IFetchAuthTokenResponse {
  token: string;
}

export const fetchAuthToken = async (): Promise<string> => {
  const response = await invoke<IFetchAuthTokenResponse>(
    TauriEvents.FETCH_AUTH_TOKEN
  );
  return response.token;
};
