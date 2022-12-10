import type { RequestOptions } from "@tauri-apps/api/http";
import { invoke } from "@tauri-apps/api/tauri";
import { Body, getClient } from "@tauri-apps/api/http";
import { config } from "src/config";

const getToken = async () => {
  return invoke<string>('get_stored_auth_token');
}

const buildRequestOptions = async (options?: RequestOptions) => {
  const authToken = await getToken();
  return {
    ...options,
    headers: {
      ...options?.headers,
      Authorization: `Bearer ${authToken}`
    }
  }
};

interface IGetParams {
  url: string,
  options?: RequestOptions;
}

export const get = async <T>({ url, options }: IGetParams): Promise<T> => {
  const client = await getClient();
  const requestOptions = await buildRequestOptions(options); 
  return client.get<T>(`${config.apiBaseUrl}/${url}`, requestOptions).then(res => res.data);
}

interface IPostParams {
  url: string;
  body: Record<any, any>;
  options?: RequestOptions;
}

export const post = async <T>({ url, body, options }: IPostParams): Promise<T> => {
  const client = await getClient();
  const requestOptions = await buildRequestOptions(options);
  return client.post<T>(`${config.apiBaseUrl}/${url}`, Body.json(body), requestOptions).then(res => res.data);
}

interface IPutParams {
  url: string;
  body: Record<any, any>;
  options?: RequestOptions;
}

export const put = async <T>({ url, body, options }: IPutParams): Promise<T> => {
  const client = await getClient();
  const requestOptions = await buildRequestOptions(options);
  return client.put<T>(`${config.apiBaseUrl}/${url}`, Body.json(body), requestOptions).then(res => res.data);
}

interface IPatchParams {
  url: string;
  options?: RequestOptions;
}

export const patch = async <T>({ url, options }: IPatchParams): Promise<T> => {
  const client = await getClient();
  const requestOptions = await buildRequestOptions(options);
  return client.patch<T>(`${config.apiBaseUrl}/${url}`, requestOptions).then(res => res.data);
}

interface IDeleteParams {
  url: string;
  options?: RequestOptions;
}

export const del = async <T>({ url, options }: IDeleteParams): Promise<T> => {
  const client = await getClient();
  const requestOptions = await buildRequestOptions(options);
  return client.delete<T>(`${config.apiBaseUrl}/${url}`, requestOptions).then(res => res.data);
}

