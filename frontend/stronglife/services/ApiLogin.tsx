import axios from 'axios';

const API_URL = 'https://stronglifeapi.fernandezpablo.es/api';

export const login = async (username: string, password: string) => {
  try {
    console.log('Initiating login request...');
    await axios.post(`${API_URL}/login`, {
      username,
      password,
    });
    console.log('Login successful');
  } catch (error) {
    console.error('Login error:', error);
    throw error;
  }
};

export const register = async (username: string, email: string, password: string) => {
  try {
    console.log('Initiating register request...');
    await axios.post(`${API_URL}/user/register`, {
      username,
      email,
      password,
    });
    console.log('Register successful');
  } catch (error) {
    console.error('Registration error:', error);
    throw error;
  }
};