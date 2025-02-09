import React, { useEffect, useState } from 'react'
import {
  View,
  Text,
  StyleSheet,
  ImageBackground,
  ActivityIndicator,
} from 'react-native'

import { useUser } from '../context/UserContext'
import { GetRoutineDay } from '../services/DayServices'
import { Button } from '@rneui/themed'
import { Routine } from '../interfaces/Routine'

const StartTraining = () => {
  const { currentUser } = useUser()
  const [routine, setRoutine] = useState<Routine | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchRoutine()
  }, [currentUser.dayselected]) // Fetch when the selected day changes

  const fetchRoutine = async () => {
    setLoading(true)
    try {
      const routineData = await GetRoutineDay(currentUser.dayselected)
      setRoutine(routineData)
    } catch (error) {
      console.error('Error fetching routine:', error)
      setRoutine(null)
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return <ActivityIndicator size="large" color="blue" />
  }

  if (!routine) {
    return <Text style={styles.noRoutineText}>No routine found</Text>
  }

  return (
    <ImageBackground
      source={{
        uri: `http://localhost:7777/api/routine/getimage?id${routine.id}`,
      }}
      style={styles.card}
    >
      <View style={styles.overlay}>
        <Text style={styles.routineText}>{routine.name}</Text>
        <Button
          title={'Comenzar entrenamiento'}
          type="solid"
          onPress={() => console.log('Start training')}
        />
      </View>
    </ImageBackground>
  )
}

const styles = StyleSheet.create({
  card: {
    width: 300,
    height: 150,
    justifyContent: 'center',
    alignItems: 'center',
    borderRadius: 15,
    overflow: 'hidden',
    elevation: 5,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.3,
    shadowRadius: 5,
  },
  overlay: {
    ...StyleSheet.absoluteFillObject,
    backgroundColor: 'rgba(0, 0, 0, 0.4)',
    justifyContent: 'center',
    alignItems: 'center',
  },
  routineText: {
    fontSize: 20,
    fontWeight: 'bold',
    color: 'white',
  },
  noRoutineText: {
    fontSize: 18,
    textAlign: 'center',
    marginTop: 20,
  },
})

export default StartTraining
