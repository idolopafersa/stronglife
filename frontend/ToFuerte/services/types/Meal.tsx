export interface NewMeal {
  name: string;
  description: string;
  calories: number;
  proteins: number;
  fats: number;
  carbs: number;
}

export interface Meal extends NewMeal {
  id: number;
}