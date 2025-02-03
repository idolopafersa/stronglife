import axios from 'axios'

const API_URL = 'http://localhost:7777/api/meal'
axios.defaults.withCredentials = true

export const GetMeal = async (id: number) => {
  try {
    const response = await axios.get(`${API_URL}/get`, {
      params: { id },
      withCredentials: true, // Allow credentials
    })
    return response.data
  } catch (error) {
    console.error('Error fetching meal:', error)
    throw error
  }
}

export const GetAllMeals = async () => {
  try {
    const response = await axios.get(`${API_URL}/getall`, {
      withCredentials: true, // Allow credentials
    })
    return response.data
  } catch (error) {
    console.error('Error fetching all meals:', error)
    throw error
  }
}

// Function to post a new meal
export const PostMeal = async (
  name: string,
  description: string,
  calories: number,
  proteins: number,
  fats: number,
  carbs: number
) => {
  try {
    const response = await axios.post(
      `${API_URL}/post`,
      { name, description, calories, proteins, fats, carbs },
      { withCredentials: true } // Allow credentials
    )
    return response.data
  } catch (error) {
    console.error('Error posting meal:', error)
    throw error
  }
}

// Function to delete a meal
export const DelMeal = async (id: number) => {
  try {
    const response = await axios.delete(`${API_URL}/del`, {
      params: { id },
    })
    return response.data
  } catch (error) {
    console.error('Error deleting meal:', error)
    throw error
  }
}

// Function to update a meal
export const PutMeal = async (
  id: number,
  name: string,
  description: string,
  calories: number,
  proteins: number,
  fats: number,
  carbs: number
) => {
  try {
    const response = await axios.put(
      `${API_URL}/put`,
      { id, name, description, calories, proteins, fats, carbs },
      { withCredentials: true } // Allow credentials
    )
    return response.data
  } catch (error) {
    console.error('Error updating meal:', error)
    throw error
  }
}
