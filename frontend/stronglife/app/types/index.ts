export type User = {
  id: string;
  email: string;
  password: string;
};

export type TrainingSession = {
  id: string;
  userId: string;
  date: string;
  exercises: Exercise[];
};

export type Exercise = {
  name: string;
  repetitions: number;
  weight: number;
};

export type CalendarDay = {
  date: string;
  completed: boolean;
};