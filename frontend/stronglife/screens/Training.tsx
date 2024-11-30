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
      <View style={styles.ejercicios} >
      <ImageBackground
    source={require('@/assets/images/churum.png')}
    
  >
        <Text style={styles.ejercicios}>Ejercicios</Text>
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
        flexDirection: 'column',
        fontSize: 40,
        color: 'red',
        fontWeight: 'bold',
        flex: 1,
      justifyContent: 'center',
      alignItems: 'center',
      }
   
});
export default Training;