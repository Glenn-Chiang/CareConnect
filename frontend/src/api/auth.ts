import type { LoginPostData, SignUpPostData, User } from "@/types/auth";
import { useMutation } from "@tanstack/react-query";
import { apiFetch } from ".";
import { useAuth } from "@/auth/AuthProvider";

export const useSignup = () => {
  return useMutation({
    mutationFn: (data: SignUpPostData) =>
      apiFetch<User>(`/signup`, {
        method: "POST",
        body: JSON.stringify(data),
      }),
  });
};

export const useLogin = () => {
  const { login } = useAuth();
  return useMutation({
    mutationFn: (credentials: LoginPostData) =>
      apiFetch<User>(`/login`, {
        method: "POST",
        body: JSON.stringify(credentials),
      }),
    onSuccess: (user) => login(user),
  });
};
