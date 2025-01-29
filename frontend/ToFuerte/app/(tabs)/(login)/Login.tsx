import React, { useState } from 'react';
import { StyleSheet, View, TextInput, Button } from 'react-native';
import { ThemedText } from '@/components/ThemedText';
import { ThemedView } from '@/components/ThemedView';
import {login , register} from '@/services/Login-Register/login';
import { useRouter } from 'expo-router';
import { useUser } from '@/app/context/UserContext';

export default function Login() {
  const router = useRouter();
  const currentUser = useUser();
  const [username, setLoginUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isRegistering, setIsRegistering] = useState(false); // En principio el usuario no se está registrando, cuando haga click en registrar, cambia el estado
  const [error, setError] = useState('');

  const handleLogin = async () => {
    try{
     const UserData =  await login(username, password);
      currentUser.username = username;
      currentUser.isLoggedIn = true;
      currentUser.id = UserData.id;
      console.log('Logging in ' + username); 
      router.push('/(tabs)/(gym)/Gym');
    }catch (error){
    };
  }
    const handleRegister = async () => {
      try{
       await register(username, email, password);
        console.log('registration susscessful');
        handleLogin();
      }catch (error){
        setError('Registration failed:' + error);
      }
    
  };

  return (
    <ThemedView style={styles.container}>
      <ThemedText type="title">Login</ThemedText>
      {error !== '' && (
        <View style={styles.errorBox}>
          <ThemedText style={styles.errorText}>{error}</ThemedText>
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
          <Button title="I have an account" onPress={() => setIsRegistering(false)} />
        </View>
      ) : (
        <View>
          <Button title="Login" onPress={handleLogin} />
          <Button title="Register" onPress={() => setIsRegistering(true)} />
        </View>
      )}
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    padding: 16,
  },
  input: {
    height: 40,
    borderColor: 'gray',
    borderWidth: 1,
    marginBottom: 12,
    paddingHorizontal: 8,
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
});