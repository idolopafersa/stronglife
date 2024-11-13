import { ImageBackground, StyleSheet, Text, View } from 'react-native'
import React from 'react'
import { SafeAreaView } from 'react-native-safe-area-context'
import { StatusBar } from 'expo-status-bar'
import CustomButton from '@/components/LittleComponents/CustomButton'

const Onboarding = () => {
  return (
    <>  
    <ImageBackground 
        source={require('../assets/images/Inicio_Sesion.png')}
        style={styles.background}
    >
    
    <SafeAreaView style = {styles.content}>
      <Text style={styles.text}>Onboarding</Text>
      <CustomButton title='Inicia sesion'
        handlePress={() =>{} }></CustomButton>
    </SafeAreaView>
    </ImageBackground>
    
    <StatusBar backgroundColor="#000000" style="light" />
    </>
  )
}

export default Onboarding

const styles = StyleSheet.create({
    background: {
        flex: 1,
        resizeMode: 'cover', // Ajusta la imagen para cubrir toda la pantalla
        justifyContent: 'center', // Centra el contenido si es necesario
      },
      content: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: 'rgba(0, 0, 0, 0.5)', // Oscurece el fondo si es necesario
      },
      text: {
        color: 'white',
        fontSize: 24,
        fontWeight: 'bold',
      },
})