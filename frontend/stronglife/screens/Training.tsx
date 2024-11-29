import React,{useState} from 'react';

import {TouchableOpacity, Text, View, Button ,StyleSheet,Image} from 'react-native';


function Training() {
   
  return (
    <View style={styles.container}>
    <View >
    <Text style={styles.header}>5 <Text style={styles.subTitle}>Ejercicios {'\n'} Restantes</Text></Text>
    
  </View>
 
    </View>
  );
}
const styles = StyleSheet.create({
    container:{
      backgroundColor: 'black',
      flex: 1,
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
        backgroundColor: '#1C1C1E',
        padding: 15,
        paddingLeft: 20,
        fontSize: 20,
        color: 'white',
      },
   
});
export default Training;