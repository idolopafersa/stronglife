import { StyleSheet, Text, ScrollView } from 'react-native'
import React from 'react'
import CustomButton from '@/components/LittleComponents/CustomButton'
import { router } from 'expo-router'
import { SafeAreaView } from 'react-native-safe-area-context'

const sign_in = () => {
  return (
    <SafeAreaView style={styles.background}>
    
        <Text>HOla</Text>
        <CustomButton 
        title='Ir a home(for any case) ' 
        handlePress={() => router.push('/AppTabs')}/>


    </SafeAreaView>
  )
}

export default sign_in

const styles = StyleSheet.create({
  background: {
    flex:1,
    backgroundColor:"#151b23",
    alignItems: 'center',
    justifyContent: 'flex-end',
  }

})