import type { LoginCredentials, User } from "@/types/auth";
import { useMutation } from "@tanstack/react-query";
import { apiFetch } from ".";

export const useSignup = () => {
  return useMutation({
    mutationFn: (newUser: Omit<User, "id">) =>
      apiFetch<User>(`/signup`, {
        method: "POST",
        body: JSON.stringify(newUser),
      }),
  });
};

export const useLogin = () => {
  return useMutation({
    mutationFn: (credentials: LoginCredentials) =>
      apiFetch<User>(`/login`, {
        method: "POST",
        body: JSON.stringify(credentials),
      }),
  });
}