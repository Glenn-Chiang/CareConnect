export type UserRole = 'caregiver' | 'recipient';

export type User = {
  id: string;
  role: UserRole
}

export type AuthResponse = {
  user: User;
  token: string;
}
