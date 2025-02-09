import axios from 'axios'
import { GetRoutine } from './RoutineServices'

const API_URL = 'http://localhost:7777/api/day'
axios.defaults.withCredentials = true

export const GetDay = async (date: string) => {
  try {
    const response = await axios.get(`${API_URL}/get`, {
      params: { date },
    })
    return response.data
  } catch (error) {
    console.error('Error fetching day:', error)
    throw error
  }
}

export const GetRoutineDay = async (date: string) => {
  try {
    const routine = await axios.get(`${API_URL}/get`, {
      params: { date },
    })
    return GetRoutine(routine.data.routine_id)
  } catch (error) {
    console.error('Error fetching day:', error)
    throw error
  }
}

export const DayRoutine = async (date: string, routine_id: number) => {
  try {
    const response = await axios.post(`${API_URL}/routine/post`, {
      params: { date, routine_id },
    })
    return response.data
  } catch (error) {
    console.error('Error Adding routine to day:', error)
    throw error
  }
}

export const PostDayMeal = async (date: string, meal_id: number) => {
  try {
    const response = await axios.post(`${API_URL}/meal/post`, {
      params: { date, meal_id },
    })
    return response.data
  } catch (error) {
    console.error('Error adding meal to day:', error)
    throw error
  }
}

export const DelDayMeal = async (date: string, meal_id: number) => {
  try {
    const response = await axios.delete(`${API_URL}/meal/del`, {
      params: { date, meal_id },
    })
    return response.data
  } catch (error) {
    console.error('Error adding meal to day:', error)
    throw error
  }
}
