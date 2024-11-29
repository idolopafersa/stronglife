import React,{useState} from 'react';

import {TouchableOpacity, Text, View, Button ,StyleSheet,Image} from 'react-native';


function Planificacion() {
   
  return (

    <View style={styles.container}>
    
      
      
    <View style={{  alignItems: 'center', justifyContent: 'center' }}>
      
      <Text>Pantalla de Inicio</Text>
      
    
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
export default Planificacion;