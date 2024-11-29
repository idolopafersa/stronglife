import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import MiGimnasio from '../screens/MiGimnasio';
import MiCocina from '../screens/MiCocina';
import { tabNavigatorStyles } from '@/styles/Appstyles';
import Icon from 'react-native-vector-icons/FontAwesome';
import Planificacion from '@/screens/Planifiacion';
import Header from '@/components/Header';

const Tab = createBottomTabNavigator();

const App = () => {
  return (
    <NavigationContainer independent={true}>
      <Header />
      <Tab.Navigator
        screenOptions={{
          tabBarStyle: tabNavigatorStyles.tabBarStyle,
          tabBarLabelStyle: tabNavigatorStyles.tabBarLabelStyle,
        }}
      >
        <Tab.Screen 
          name="MiGimnasio" 
          component={MiGimnasio}
          options={{
            tabBarIcon: () => <Icon name="heartbeat" size={30} color="#A5D32D" />,
            tabBarLabel: 'Mi Gimnasio',
            headerShown: false,
          }}
        />
        <Tab.Screen 
          name="MiCocina" 
          component={MiCocina}
          options={{
            tabBarIcon: () => <Icon name="cutlery" size={30} color="#A5D32D" />,
            tabBarLabel: 'Mi Cocina',
            headerShown: false,
          }}
        />
        <Tab.Screen 
          name="Planifiacion" 
          component={Planificacion}
          options={{
            tabBarIcon: () => <Icon name="table" size={30} color="#A5D32D" />,
            tabBarLabel: 'Mi Planificacion',
            headerShown: false,
          }}
        />
      </Tab.Navigator>
    </NavigationContainer>
  );
};

export default App;