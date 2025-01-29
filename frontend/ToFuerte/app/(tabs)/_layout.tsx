import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import React from 'react';
import { Platform } from 'react-native';

import { HapticTab } from '@/components/HapticTab';
import { IconSymbol } from '@/components/ui/IconSymbol';
import TabBarBackground from '@/components/ui/TabBarBackground';
import { Colors } from '@/constants/Colors';
import { useColorScheme } from '@/hooks/useColorScheme';
import Gym from './(gym)/Gym';
import Kitchen from './(kitchen)/Kitchen';
import Params from './(params)/_layout';
import { CustomHeader } from '@/components/CustomHeader';

const Tab = createBottomTabNavigator();

export default function TabLayout() {
  const colorScheme = useColorScheme();

  return (  
    <>
    <CustomHeader/>
    
    <Tab.Navigator
      screenOptions={{
        headerShown: false,  //Quita el header 
        tabBarActiveTintColor: Colors[colorScheme ?? 'light'].tint,
        tabBarButton: HapticTab,
        tabBarBackground: TabBarBackground,
        tabBarStyle: Platform.select({
          ios: {
            // Use a transparent background on iOS to show the blur effect
            position: 'absolute',
          },
          default: {},
        }),
      }}
    >
     
      <Tab.Screen
        name="Gym"
        component={Gym}
        options={{
          title: 'Gimnasio',
          tabBarIcon: ({ color }) => <IconSymbol size={28} name="paperplane.fill" color={color} />,
        }}
      />
      <Tab.Screen
        name="Kitchen"
        component={Kitchen}
        options={{
          title: 'Cocina',
          tabBarIcon: ({ color }) => <IconSymbol size={28} name="paperplane.fill" color={color} />,
        }}
      />
      <Tab.Screen
        name="Params"
        component={Params}
        options={{
          title: 'Ajustes',
          tabBarIcon: ({ color }) => <IconSymbol size={28} name="paperplane.fill" color={color} />,
        }}
      />
    </Tab.Navigator>
    </>
  );
}