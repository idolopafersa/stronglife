import { useUser } from '@/app/context/UserContext';
import axios from 'axios';

const API_URL = 'https://stronglifeapi.fernandezpablo.es/api';


export const login = async (username: string, password: string) => {
  try {
    console.log('Initiating login request...');
    const response = await axios.post(`${API_URL}/login`, {
      username,
      password,
    },{
      withCredentials: true,
    }
    
  );
    console.log('Login successful', response.data);
    return response.data;
  } catch (error) {
    console.error('Login failed', error);
    throw error;
  }
};

export const register = async (username: string, email: string, password: string) => {
  try {
    console.log('Initiating register request...');
    const response = await axios.post(`${API_URL}/user/register`, {
      username,
      email,
      password,
    });
    console.log('Register successful', response.data);
    return response.data;
  } catch (error) {
    console.error('Registration failed', error);
    throw error;
  }
};