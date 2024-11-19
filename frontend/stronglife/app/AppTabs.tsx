import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { Button } from 'react-native';
import HelpScreen from '../components/HelpScreen';
import HomeScreen from '../components/HomeScreen';
import DetailsScreen from '../components/DetailsScreen';
import Icon from "react-native-vector-icons/FontAwesome5"

const Tab = createBottomTabNavigator();

export default function App() {
  return (
    <NavigationContainer>
      <Tab.Navigator initialRouteName="Home"
      screenOptions={{
        animation: 'fade',
      }}>
        
        <Tab.Screen 
          name="Home" 
          component={HomeScreen}
          options={{
            tabBarIcon: () => <Icon name="building" size={30} color="#A5D32D" />,
            tabBarLabel: 'Mi Gimnasio',
            tabBarLabelStyle: {color: '#A5D32D'},
            headerShown: false,
              
          }}
        />
        <Tab.Screen 
          name="Details" 
          component={DetailsScreen}
          options={{
            tabBarIcon: () => <Icon name="user" size={30} color="#A5D32D" />,
            tabBarLabel: 'Mi Cocina',
            tabBarLabelStyle: {color: '#A5D32D'},
            headerTintColor: 'white',
            headerStyle: { backgroundColor: 'tomato' },
          }}
        />
        <Tab.Screen 
          name="Help" 
          component={HelpScreen}
          options={{
            tabBarIcon: () => <Icon name="calendar" size={30} color="#A5D32D" />,
            tabBarLabel: 'Planificacion',
            tabBarLabelStyle: {color: '#A5D32D'},
            
            headerShown: false,
          }}
        />

      </Tab.Navigator>
    </NavigationContainer>
  );
}