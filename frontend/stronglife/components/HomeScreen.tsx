import React,{useState} from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import {TouchableOpacity, Text, View, Button ,StyleSheet,Image} from 'react-native';
import LabelBottomNavigation from './LittleComponents/LabelBottomNavigation';
import Anothercalendar from './LittleComponents/Anothercalendar';
import CombinedCalender from './LittleComponents/CombinedCalender';
function HomeScreen({ navigation }) {
    const [showGif, setShowGif] = useState(false);
  return (



    <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center',backgroundColor:'grey' }}>
      <Anothercalendar/>
      <Text>Pantalla de Inicio</Text>
      <TouchableOpacity
              style={estilos.button} 
              onPress={() => navigation.navigate('Details')}
            >
              <Text style={estilos.buttonText}>Ir a Detalles</Text>
            </TouchableOpacity>
 <TouchableOpacity style={estilos.button} 
              onPress={() => navigation.navigate('Help')}
            >
               <Text style={estilos.buttonText}>Ayuda</Text>
            </TouchableOpacity>
            <CombinedCalender/>

    </View>
  );
}
const estilos = StyleSheet.create({
  button: {
    backgroundColor: 'gold',  
    padding: 10,
    borderRadius: 5,
    alignItems: 'center',
  },
  buttonText: {
    color: 'white',  
    fontWeight: 'bold',
  },
  container:{
      flex: 1,
      },
   
});
export default HomeScreen;