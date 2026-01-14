import type { UserRole } from "./auth";

export interface Todo {
  id: string;
  title: string;
  description: string;
  dueDate: Date;
  completed: boolean;
  recipientId: string;
  caregiverId: string;
  priority: 'low' | 'medium' | 'high';
}

export type MoodType = 'happy' | 'sad' | 'neutral' | 'anxious' | 'excited';

export interface JournalEntry {
  id: string;
  recipientId: string;
  content: string;
  mood: MoodType;
  createdAt: Date;
  hasVoiceMessage: boolean;
  voiceMessageUrl?: string;
}

export interface Comment {
  id: string;
  journalEntryId: string;
  authorId: string;
  authorRole: UserRole;
  content: string;
  createdAt: Date;
}