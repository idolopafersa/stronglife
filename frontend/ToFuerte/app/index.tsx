import React, { useContext } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';

import { UserContext } from '@/app/context/UserContext';
import Login from './(tabs)/(login)/Login';
import Gym from './(tabs)/(gym)/Gym';

const Stack = createStackNavigator();

function Index() {
  const currentUser = useContext(UserContext);

  return (
    
      <Stack.Navigator initialRouteName={currentUser?.isLoggedIn ? "Gym" : "Login"}>
        <Stack.Screen name="Login" component={Login} options={{headerShown: false}} />
        <Stack.Screen name="Gym" component={Gym}  />
      </Stack.Navigator>
    
  );
}

export default Index;