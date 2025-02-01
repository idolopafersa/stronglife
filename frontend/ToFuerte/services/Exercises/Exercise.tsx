import axios from 'axios';
import { ExpoRoot } from 'expo-router';
import { Exercise, NewExercise } from '../types/Exercise';

const API_URL = 'https://stronglifeapi.fernandezpablo.es/api/exercise';


export const getExercises = async () => {
    const response = await axios.get(`${API_URL}/getall`);
    return response.data;
}

export const getExercise = async (id: number) => {
    const response = await axios.get(`${API_URL}/get${id}`);
    return response.data;
}

export const createExercise = async (exercise: NewExercise ) => {
    const response = await axios.post(`${API_URL}/post`, exercise);
    return response.data;
}

export const updateExercise = async (exercise: Exercise ) => {
    const response = await axios.put(`${API_URL}/put`, exercise);
    return response.data;
}

export const deleteExercise = async (id: number) => {
    const response = await axios.delete(`${API_URL}/del${id}`);
    return response.data;
}