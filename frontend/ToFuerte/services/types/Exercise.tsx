export interface NewExercise {
  name: string;
  description: string;
}

export interface Exercise extends NewExercise {
  id: number;
}