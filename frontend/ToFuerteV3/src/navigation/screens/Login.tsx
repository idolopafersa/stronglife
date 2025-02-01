import React, { useState } from 'react'
import { StyleSheet, View, Text, TextInput, Button } from 'react-native'

import { useNavigation } from '@react-navigation/native'
import { login, register } from '../../services/UserServices'
import { useUser } from '../../context/UserContext'

export default function Login() {
  const { currentUser, setCurrentUser } = useUser()
  const navigator = useNavigation()
  const [username, setLoginUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [isRegistering, setIsRegistering] = useState(false)
  const [error, setError] = useState('')

  const handleLogin = async () => {
    try {
      const UserData = await login(username, password)
      setCurrentUser({
        ...currentUser,
        username,
        isLoggedIn: true,
        id: UserData.id,
      })
      console.log('Logging in ' + username)
    } catch (error) {
      setError('error login') // Set error inside catch
    }
  }

  const handleRegister = async () => {
    try {
      await register(username, email, password)
      console.log('Registration successful')
      handleLogin()
    } catch (error) {
      setError("Couldn't register")
    }
  }

  return (
    <View style={styles.container}>
      <Text style={styles.title}>{isRegistering ? 'Register' : 'Login'}</Text>
      {error !== '' && (
        <View style={styles.errorBox}>
          <Text style={styles.errorText}>{error}</Text>
        </View>
      )}
      <TextInput
        style={[styles.input, styles.inputText]}
        placeholder="Username"
        placeholderTextColor="#fff"
        value={username}
        onChangeText={setLoginUsername}
        keyboardType="default"
        autoCapitalize="none"
      />
      {isRegistering && (
        <TextInput
          style={[styles.input, styles.inputText]}
          placeholder="Email"
          placeholderTextColor="#fff"
          value={email}
          onChangeText={setEmail}
          keyboardType="email-address"
          autoCapitalize="none"
        />
      )}
      <TextInput
        style={[styles.input, styles.inputText]}
        placeholder="Password"
        placeholderTextColor="#fff"
        value={password}
        onChangeText={setPassword}
        secureTextEntry
      />
      {isRegistering ? (
        <View>
          <Button title="Register" onPress={handleRegister} />
          <Button
            title="I have an account"
            onPress={() => setIsRegistering(false)}
          />
        </View>
      ) : (
        <View>
          <Button title="Login" onPress={handleLogin} />
          <Button title="Register" onPress={() => setIsRegistering(true)} />
        </View>
      )}
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    padding: 16,
    backgroundColor: '#121212', // Dark theme background
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    color: 'white',
    textAlign: 'center',
    marginBottom: 20,
  },
  input: {
    height: 40,
    borderColor: 'gray',
    borderWidth: 1,
    marginBottom: 12,
    paddingHorizontal: 8,
    borderRadius: 5,
  },
  inputText: {
    color: 'white',
  },
  errorBox: {
    backgroundColor: 'red',
    padding: 10,
    marginBottom: 12,
    borderRadius: 5,
  },
  errorText: {
    color: 'white',
  },
})
