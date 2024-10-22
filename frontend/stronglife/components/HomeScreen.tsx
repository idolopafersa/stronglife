import React,{useState} from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import {TouchableOpacity, Text, View, Button ,StyleSheet,Image} from 'react-native';

function HomeScreen({ navigation }) {
    const [showGif, setShowGif] = useState(false);
  return (



    <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>

      <Text>Pantalla de Inicio</Text>
      <TouchableOpacity
              style={estilos.button} // Estilos personalizados
              onPress={() => navigation.navigate('Details')}
            >
              <Text style={estilos.buttonText}>Ir a Detalles</Text>
            </TouchableOpacity>
 <TouchableOpacity style={estilos.button} // Estilos personalizados
              onPress={() => navigation.navigate('Help')}
            >
               <Text style={estilos.buttonText}>Ayuda</Text>
            </TouchableOpacity>


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
   gif: {
      width: 200,
      height: 200,
      marginVertical: 20,
    },
});
export default HomeScreen;