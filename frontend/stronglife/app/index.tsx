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
import TabBar from '@/components/LittleComponents/TabBar';

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
          tabBarActiveBackgroundColor: '#1C1C1E',
        }}
        tabBar={props => <TabBar {...props} />}
      screenListeners={{
        state: (e) => {
          const route = e.data.state.routes[e.data.state.index];
          if (route.name === 'Training') {
            setejercicio(false);
          } else {
            setejercicio(true);
          }
        }
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
        
        <Tab.Screen 
          name="Training" 
          component={Training}
          options={{
        tabBarIcon: () => <Icon name="warning" size={30} color="#A5D32D" />,
        tabBarLabel: 'Mi Gimnasio',
        headerShown: false,
          }}
        />
      </Tab.Navigator>
    </NavigationContainer>
  );
};

export default App;