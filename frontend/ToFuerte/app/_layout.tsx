import { DarkTheme, DefaultTheme, ThemeProvider } from '@react-navigation/native';
import { useFonts } from 'expo-font';
import { Redirect, Stack } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';
import { StatusBar } from 'expo-status-bar';
import { useEffect, useState } from 'react';
import 'react-native-reanimated';

import { useColorScheme } from '@/hooks/useColorScheme';
import { UserContext, UserContextProps } from './context/UserContext';
import Login from './(tabs)/(login)/Login';
import { Header } from '@react-navigation/stack';
import { CustomHeader } from '@/components/CustomHeader';
import { LocaleConfig } from 'react-native-calendars';
import moment from 'moment';

// Prevent the splash screen from auto-hiding before asset loading is complete.
SplashScreen.preventAutoHideAsync();

export default function RootLayout() {
  const [currentUser, setCurrentUser] = useState<UserContextProps>({username: '',id:0 ,email: '',dayselected :moment().format('YYYY-MM-DD') , isLoggedIn: false});
  const colorScheme = useColorScheme();
  const [loaded] = useFonts({
    SpaceMono: require('../assets/fonts/SpaceMono-Regular.ttf'),
  });

  useEffect(() => {
    if (loaded) {
      SplashScreen.hideAsync();
    }
  }, [loaded]);

  if (!loaded) {
    return null;
  }


  LocaleConfig.locales['es'] = {
    monthNames: ['Enero', 'Febrero', 'Marzo', 'Abril', 'Mayo', 'Junio', 'Julio', 'Agosto', 'Septiembre', 'Octubre', 'Noviembre', 'Diciembre'],
    monthNamesShort: ['Ene', 'Feb', 'Mar', 'Abr', 'May', 'Jun', 'Jul', 'Ago', 'Sep', 'Oct', 'Nov', 'Dic'],
    dayNames: ['Domingo', 'Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado'],
    dayNamesShort: ['D', 'L', 'M', 'X', 'J', 'V', 'S'],
    today: 'Hoy'
  };
  LocaleConfig.defaultLocale = 'es';
 

  return (
    <ThemeProvider value={colorScheme === 'dark' ? DarkTheme : DefaultTheme}>
      <UserContext.Provider value={currentUser}>
      
        <Stack  >
            <Stack.Screen name="(tabs)" options={{headerShown : false}}/>  // lo primero que intenta el programa es accedfer al _layout de tabs para ver el principal
            <Stack.Screen name="+not-found" />
        </Stack>
        <StatusBar style="auto" />
      </UserContext.Provider>
    </ThemeProvider>
  );
}
