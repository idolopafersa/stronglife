import React,{useState} from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import {TouchableOpacity, Text, View, Button ,StyleSheet,Image} from 'react-native';
import Calendar from './LittleComponents/Calendar';
import Paper from '@mui/material/Paper';
import LabelBottomNavigation from './LittleComponents/LabelBottomNavigation';
import Calendarmui from './LittleComponents/Calendarmui';
function DetailsScreen() {
  return (
    <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' ,backgroundColor:'grey'}}>
      <Text>Pantalla de Detalles</Text>
      <Calendarmui/>
      
    </View>
  );
}
export default DetailsScreen;