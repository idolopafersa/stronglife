import React, { useState } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import MiGimnasio from './screens/MiGimnasio';
import MiCocina from './screens/MiCocina';
import Planificacion from './screens/Planificacion';
import Training from './screens/Training';
import Login from './screens/Login';
import { tabNavigatorStyles } from './styles/Appstyles';
import TabBar from './components/LittleComponents/TabBar';

const Tab = createBottomTabNavigator();

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  return (
    <NavigationContainer independent={true}>
      {isLoggedIn ? (
        <Tab.Navigator
          screenOptions={{
            tabBarStyle: tabNavigatorStyles.tabBarStyle,
            tabBarLabelStyle: tabNavigatorStyles.tabBarLabelStyle,
            tabBarInactiveBackgroundColor: '#1C1C1E',
            tabBarActiveBackgroundColor: '#1C1C1E',
          }}
          tabBar={props => (
            <TabBar
              {...props}
              activeBackgroundColor="#1C1C1E"
              activeTintColor="white"
              inactiveBackgroundColor="#1C1C1E"
              inactiveTintColor="gray"
            />
          )}
        >
          <Tab.Screen 
            name="MiGimnasio" 
            component={MiGimnasio}
            options={{ headerShown: false }}
          />
          <Tab.Screen 
            name="MiCocina" 
            component={MiCocina}
            options={{ headerShown: false }}
          />
          <Tab.Screen 
            name="Planificacion" 
            component={Planificacion}
            options={{ headerShown: false }}
          />
          <Tab.Screen 
            name="Training" 
            component={Training}
            options={{ headerShown: false }}
          />
        </Tab.Navigator>
      ) : (
        <Login setIsLoggedIn={setIsLoggedIn} />
      )}
    </NavigationContainer>
  );
};

export default App;