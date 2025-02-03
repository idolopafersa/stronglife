import axios from 'axios'

const API_URL = 'http://localhost:7777/api'

export const login = async (username: string, password: string) => {
  if (username === '' || password === '') {
    throw new Error('Username or password cannot be empty')
  }
  try {
    console.log('Initiating login request...')
    const response = await axios.post(
      `${API_URL}/login`,
      {
        username,
        password,
      },
      {
        withCredentials: true,
      }
    )
    console.log('Login successful', response.data)
    return response.data
  } catch (error) {
    console.error('Login failed', error)
    throw error
  }
}

export const register = async (
  username: string,
  email: string,
  password: string
) => {
  if (username === '' || password === '' || email === '') {
    throw new Error('Username or password or email cannot be empty')
  }
  let reg = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w\w+)+$/
  if (reg.test(email) === false) {
    throw new Error('Invalid email')
  }
  try {
    console.log('Initiating register request...')
    const response = await axios.post(`${API_URL}/user/register`, {
      username,
      email,
      password,
    })
    console.log('Register successful', response.data)
    return response.data
  } catch (error) {
    console.error('Registration failed', error)
    throw error
  }
}

export const checkcookie = async () => {
  const response = await axios.get(`${API_URL}/cookie`)
  if (response.status == 200) {
    return response
  }
  return null
}
