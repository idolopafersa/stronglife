import React,{useState} from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import { Button } from 'react-native';
import {TouchableOpacity, Text, View ,StyleSheet,Image, Animated} from 'react-native';
import HelpScreen from '../components/HelpScreen.tsx';
import HomeScreen from '../components/HomeScreen.tsx';
import DetailsScreen from '../components/DetailsScreen.tsx';
<<<<<<< Updated upstream
import Widget from '../components/LittleComponents/Widget.tsx';
=======
import Calendar from '@/components/LittleComponents/Calendar.js';
>>>>>>> Stashed changes



export function HelloWidgetPreviewScreen() {
  return (
    <View style={styles.container}>
      <WidgetPreview
        renderWidget={() => <HelloWidget />}
        width={320}
        height={200}
      />
    </View>
  );
}

<<<<<<< Updated upstream
const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
});
const forFade = ({ current, next }) => {
  const opacity = Animated.add(
    current.progress,
    next ? next.progress : 0
  ).interpolate({
    inputRange: [0, 1, 2],
    outputRange: [0, 1, 0],
  });

  return {
    leftButtonStyle: { opacity },
    rightButtonStyle: { opacity },
    titleStyle: { opacity },
    backgroundStyle: { opacity },
  };
};
=======

>>>>>>> Stashed changes
const Stack = createStackNavigator();

export default function App() {
  return (
    <NavigationContainer independent={true}>
      <Stack.Navigator initialRouteName="Home">
        <Stack.Screen name="Home" component={HomeScreen}
        options={{
                          headerTintColor: 'white',
                          headerStyle: { backgroundColor: 'tomato' },
                          headerRight: () => (
                                      <Button
                                        onPress={() => alert('This is a button!')}
                                        title="Sing up"
                                        color="white"
                                      />
                                    ),
                          }}/>
        <Stack.Screen name="Details" component={DetailsScreen}
        options={{
                          headerTintColor: 'white',
                          headerStyle: { backgroundColor: 'tomato' },}}/>
        <Stack.Screen
        name="Help" component={HelpScreen}
        options={{
                  headerTintColor: 'white',
                  headerStyle: { backgroundColor: 'tomato' }}}
         />
      </Stack.Navigator>
    </NavigationContainer>
  );
}