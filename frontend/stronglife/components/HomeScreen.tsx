import React,{useState} from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import {TouchableOpacity, Text, View, Button ,StyleSheet,Image} from 'react-native';
import Calendar from './LittleComponents/Calendar';
import CustomButton from './LittleComponents/CustomButton';

function HomeScreen() {
   
  return (

    <View style={styles.container}>
    
      <Calendar/>
      
    <View style={{  alignItems: 'center', justifyContent: 'center' }}>
      
      <Text>Pantalla de Inicio</Text>
      <CustomButton
        title = "homero chino"
        handlePress={() => {}}
        
        />

    
    </View>
    </View>
  );
}
const styles = StyleSheet.create({
  button: {
    backgroundColor: 'gold',  // Color dorado de fondo
    padding: 10,
    borderRadius: 5,
    alignItems: 'center',
  },
  buttonText: {
    color: 'white',  // Color del texto
    fontWeight: 'bold',
  },
  container:{
      backgroundColor: 'black',
      flex: 1,
      },
   
});
export default HomeScreen;