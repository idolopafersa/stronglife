import React,{useState} from 'react';

import {TouchableOpacity, Text, View, Button ,StyleSheet,Image, ImageBackground} from 'react-native';


function Training() {
   
  return (
    
    <View style={styles.container}>
    <View >
    <Text style={styles.header}>5 <Text style={styles.subTitle}>Ejercicios {'\n'} Restantes</Text></Text>
    
  </View>
    <View style={styles.subcontainer}>
      
      <Text style={styles.subTitle}>Siguiente: <Text style={{color: 'grey'}}>Aerobics-Sets (1/1)</Text></Text>
    </View>
      <View style={styles.foto}>
      <ImageBackground  
    source={require('@/assets/images/churum.png')}
    resizeMode="cover" style={styles.foto}
  >
        <View style={styles.ejercicioscont}>
        <Text style={styles.ejercicios}>Ejercicios</Text>
        <Text style={styles.ejercicios}>Restantes</Text>
        <Text style={styles.repeticiones}>3 Repeticiones</Text>
        </View>
      </ImageBackground>  
      </View>
    </View>
  );
}
const styles = StyleSheet.create({
    container:{
      backgroundColor: 'black',
      
      
      flex: 1,
      },subcontainer:{
        backgroundColor: '#3A3A3C',
        
        },
      header:{
        flexDirection: 'row',
        alignItems: 'center',
        justifyContent: 'space-between',
        backgroundColor: '#1C1C1E',
        padding: 15,
        paddingLeft: 20,
        color: 'white',
        fontSize: 30,
        fontWeight: 'bold',
      },subTitle:{
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'space-between',
        padding: 15,
        paddingLeft: 20,
        fontSize: 20,
        color: 'white',
        fontWeight: 'bold',
      },ejercicios:{
        fontSize: 42,
        lineHeight: 84,
        fontWeight: 'bold',
        textAlign: 'center',
        backgroundColor: '#000000c0',
        color:'red',
      },
      ejercicioscont:{
        alignItems: 'center',
        
        borderColor: 'yellow',
        borderWidth: 2,
      },
      foto:{
        resizeMode: 'stretch',
        width: '100%',
        height: '100%',
        justifyContent: 'center',
        textAlign: 'center',
      },
      repeticiones:{
        color:'red',
        fontSize: 20,
        fontWeight: 'bold',
        textAlign: 'right'
      },
   
});
export default Training;