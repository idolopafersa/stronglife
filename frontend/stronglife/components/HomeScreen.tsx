import React,{useState} from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import {TouchableOpacity, Text, View, Button ,StyleSheet,Image} from 'react-native';
import Calendar from './LittleComponents/Calendar';

function HomeScreen({ navigation }) {
    const [showGif, setShowGif] = useState(false);
  return (

    <View>
    <View>
      <Calendar/>
      </View>
    <View style={{  alignItems: 'center', justifyContent: 'center' }}>
      
      <Text>Pantalla de Inicio</Text>
      <TouchableOpacity
              style={estilos.button} 
              onPress={() => navigation.navigate('Details')}
            >
              <Text style={estilos.button}>Ir a Detalles</Text>
            </TouchableOpacity>
 <TouchableOpacity style={estilos.button} 
              onPress={() => navigation.navigate('Help')}
            >
               <Text style={estilos.button}>Ayuda</Text>
            </TouchableOpacity>


    </View>
    </View>
  );
}
const estilos = StyleSheet.create({
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
      flex: 1,
      },
   
});
export default HomeScreen;