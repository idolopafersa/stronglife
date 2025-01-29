import React, { useContext, useState } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
 
import {UserContext, UserContextProps} from '@/app/context/UserContext';
import Gym from './(gym)/Gym';
import Login from './(login)/Login';
const Stack = createStackNavigator();

function Index() {  //Punto de entrada, se encarga de redirigir a la pantalla de login o a la pantalla Gym
    const currentUser = useContext(UserContext);
  return (
    <Stack.Navigator  screenOptions={{headerShown:false}} initialRouteName={currentUser?.isLoggedIn ? "Home" : "Login"}>
      
        <Stack.Screen name="Gym" component={Gym} />
      
        <Stack.Screen name="Login"  component={Login}  />
     
    </Stack.Navigator>
  );
}

export default Index;