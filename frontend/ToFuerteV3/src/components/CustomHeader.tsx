import { StyleSheet, View, Text, Button, TouchableOpacity } from 'react-native'
import { Image } from 'react-native'
import { useState } from 'react'

import { useNavigation } from '@react-navigation/native'
import { useUser } from '../context/UserContext'

export const CustomHeader = () => {
  const currentUser = useUser()

  return (
    <View style={styles.headerContainer}>
      <Text style={styles.greetingText}>
        ¿Listo para entrenar, {currentUser.currentUser.username}?
      </Text>

      <TouchableOpacity
        onPress={() => {
          screen: 'Profile'
        }}
      >
        <Image
          source={{
            uri: `http://stronglifeapi.fernandezpablo.es/api/user/getimage?user_id=${currentUser.currentUser.id}`,
          }}
          style={styles.userImage}
        />
      </TouchableOpacity>
    </View>
  )
}

const styles = StyleSheet.create({
  headerContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    padding: 20,
    backgroundColor: 'green', // Cambia este color según tus necesidades
  },
  greetingText: {
    color: '#fff',
    fontSize: 18,
  },
  userImage: {
    width: 40,
    height: 40,
    borderRadius: 20,
  },
})
