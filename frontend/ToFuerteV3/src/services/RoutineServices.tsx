import axios from 'axios'

const API_URL = 'http://localhost:7777/api/routine'
axios.defaults.withCredentials = true

// Function to get a specific exercise by ID
export const GetRoutine = async (routine_id: number) => {
  try {
    const response = await axios.get(`${API_URL}/get`, {
      params: { routine_id },
    })
    return response.data
  } catch (error) {
    console.error('Error fetching exercise:', error)
    throw error
  }
}

// Function to get all exercises
export const GetAllRoutines = async () => {
  try {
    const response = await axios.get(`${API_URL}/getall`)
    return response.data
  } catch (error) {
    console.error('Error fetching all exercises:', error)
    throw error
  }
}

// Function to post a new exercise
export const PostRoutine = async (name: string, description: string) => {
  try {
    const response = await axios.post(`${API_URL}/post`, { name, description })
    return response.data
  } catch (error) {
    console.error('Error posting exercise:', error)
    throw error
  }
}

// Function to delete an exercise
export const DelRoutine = async (routine_id: number) => {
  try {
    const response = await axios.delete(`${API_URL}/del`, {
      params: { routine_id: routine_id },
    })
    return response.data
  } catch (error) {
    console.error('Error deleting exercise:', error)
    throw error
  }
}

// Function to update an exercise
export const PutRoutine = async (
  id: number,
  name: string,
  description: string
) => {
  try {
    const response = await axios.put(`${API_URL}/put`, {
      id,
      name,
      description,
    })
    return response.data
  } catch (error) {
    console.error('Error updating exercise:', error)
    throw error
  }
}
