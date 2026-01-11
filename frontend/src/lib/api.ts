import type { dto_ErrorResponse } from "./api/models/dto_ErrorResponse";
export type { dto_SuccessResponse } from "./api/models/dto_SuccessResponse";

const API_URL = import.meta.env.VITE_API_URL;

export async function apiFetch(path: string, options: RequestInit = {}) {
  return fetch(`${API_URL}${path}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...(options.headers || {}),
    },
    credentials: "include",
  });
}

export async function apiFetchTyped<T>(
  path: string,
  options: RequestInit = {},
): Promise<{ data: T | null; error: dto_ErrorResponse | null; ok: boolean }> {
  const response = await apiFetch(path, options);

  if (!response.ok) {
    const error = (await response.json()) as dto_ErrorResponse;
    return { data: null, error, ok: false };
  }

  const data = (await response.json()) as T;
  return { data, error: null, ok: true };
}
