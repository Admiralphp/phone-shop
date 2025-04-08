// src/app/core/models/user.model.ts
export interface User {
  id: string;
  username: string;
  email: string;
  role: 'admin' | 'manager' | 'support';
  createdAt: Date;
  lastLogin: Date;
}