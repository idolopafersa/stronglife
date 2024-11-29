import React, { useState } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import MiGimnasio from '../screens/MiGimnasio';
import MiCocina from '../screens/MiCocina';
import { tabNavigatorStyles } from '@/styles/Appstyles';
import Icon from 'react-native-vector-icons/FontAwesome';
import Planificacion from '@/screens/Planifiacion';
import Header from '@/components/Header';
import Training from '@/screens/Training';

const Tab = createBottomTabNavigator();

const App = () => {
  const [singedin, setsingedin] = useState(true);
  const [ejercicio, setejercicio] = useState(true);
  return (
    
    <NavigationContainer independent={true}>
      {ejercicio && <Header />}
      <Tab.Navigator
        screenOptions={{
          tabBarStyle: tabNavigatorStyles.tabBarStyle,
          tabBarLabelStyle: tabNavigatorStyles.tabBarLabelStyle,
          tabBarInactiveBackgroundColor: '#1C1C1E',
        }}
      >
        <Tab.Screen 
          name="MiGimnasio" 
          component={MiGimnasio}
          listeners={{
            tabPress: () => setejercicio(true),
              }}
          options={{
        tabBarIcon: () => <Icon name="heartbeat" size={30} color="#A5D32D" />,
        tabBarLabel: 'Mi Gimnasio',
        headerShown: false,
          }}
        />
        <Tab.Screen 
          name="MiCocina" 
          component={MiCocina}
          listeners={{
            tabPress: () => setejercicio(true),
              }}
          options={{
        tabBarIcon: () => <Icon name="cutlery" size={30} color="#A5D32D" />,
        tabBarLabel: 'Mi Cocina',
        headerShown: false,
          }}
        />
        <Tab.Screen 
          name="Planifiacion" 
          component={Planificacion}
          listeners={{
            tabPress: () => setejercicio(true),
              }}
          options={{
        tabBarIcon: () => <Icon name="table" size={30} color="#A5D32D" />,
        tabBarLabel: 'Mi Planificacion',
        headerShown: false,
          }}
        />
        
        <Tab.Screen 
          name="EntrenamientoPrueba" 
          component={Training}
          listeners={{
        tabPress: () => setejercicio(false),
          }}
          options={{
        tabBarIcon: () => <Icon name="heartbeat" size={30} color="#A5D32D" />,
        tabBarLabel: 'Mi Gimnasio',
        headerShown: false,
          }}
        />
      </Tab.Navigator>
    </NavigationContainer>
  );
};

export default App;