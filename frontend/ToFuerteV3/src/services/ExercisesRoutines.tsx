import axios from 'axios'

const API_URL = 'http://localhost:7777/api/exercises/routines'

// 🔹 Get all exercises for a routine
export const GetRoutineExercises = async (routine_id: number) => {
  console.log(`Entering with routine: ${routine_id}`)
  try {
    const response = await axios.get(`${API_URL}/get`, {
      params: { routine_id: routine_id },
    })
    return response.data
  } catch (error) {
    console.error('Error fetching routine exercises:', error)
    throw error
  }
}

// 🔹 Add an exercise to a routine
export const AddRoutineExercise = async (
  routine_id: number,
  exercise_id: number
) => {
  try {
    const response = await axios.post(
      `${API_URL}/post?routine_id=${routine_id}&exercise_id=${exercise_id}`
    )
    return response.data
  } catch (error) {
    console.error('Error adding exercise to routine:', error)
    throw error
  }
}

// 🔹 Delete an exercise from a routine
export const DeleteRoutineExercise = async (
  routine_id: number,
  exercise_id: number
) => {
  try {
    const response = await axios.delete(`${API_URL}/del`, {
      params: { routine_id, exercise_id },
      withCredentials: true,
    })
    return response.data
  } catch (error) {
    console.error('Error deleting exercise from routine:', error)
    throw error
  }
}
