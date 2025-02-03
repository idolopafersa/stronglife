import React, { useState } from 'react'
import { View, StyleSheet } from 'react-native'
import { Button, Card, Input } from '@rneui/themed'
import {
  PutExercise,
  DelExercise,
  PostExercise,
} from '../../../services/ExerciseServices'
import { useNavigation } from '@react-navigation/native'

const EditExercise = ({ route, navigation }: any) => {
  const navigator = useNavigation()
  const { exercise } = route.params

  const isCreating = !exercise // If no exercise is provided, it's creating a new one
  const [name, setName] = useState(isCreating ? '' : exercise.name) // Empty for creation, else use existing name
  const [description, setDescription] = useState(
    isCreating ? '' : exercise.description
  )

  const handleSave = async () => {
    if (!name.trim() || !description.trim()) {
      return
    }

    try {
      if (isCreating) {
        await PostExercise(name, description)
      } else {
        await PutExercise(exercise.id, name, description)
      }

      navigator.navigate('ExerciseList')
    } catch (error) {
      console.error('Error saving exercise:', error)
    }
  }

  // Handle Delete Exercise
  const handleDelete = async () => {
    try {
      if (!isCreating) {
        await DelExercise(exercise.id)
        navigator.navigate('ExerciseList')
      }
    } catch (error) {
      console.error('Error deleting exercise:', error)
    }
  }

  return (
    <View style={styles.container}>
      <Card containerStyle={styles.card}>
        <Card.Title>
          {isCreating ? 'Create Exercise' : 'Edit Exercise'}
        </Card.Title>
        <Card.Divider />

        <Input
          label="Exercise Name"
          value={name}
          onChangeText={setName}
          placeholder="Enter exercise name"
          leftIcon={{ type: 'font-awesome', name: 'dumbbell' }}
        />

        <Input
          label="Description"
          value={description}
          onChangeText={setDescription}
          placeholder="Enter description"
          multiline
          numberOfLines={3}
          leftIcon={{ type: 'font-awesome', name: 'file-alt' }}
        />

        <Button
          title={isCreating ? 'Create Exercise' : 'Save Changes'}
          onPress={handleSave}
          buttonStyle={styles.saveButton}
        />

        {!isCreating && (
          <Button
            title="Delete Exercise"
            onPress={handleDelete}
            buttonStyle={styles.deleteButton}
          />
        )}
      </Card>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f4f4f4',
    justifyContent: 'center',
  },
  card: {
    borderRadius: 12,
    padding: 20,
    elevation: 5, // Shadow for Android
    shadowColor: '#000', // Shadow for iOS
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.2,
    shadowRadius: 4,
  },
  saveButton: {
    backgroundColor: '#28a745',
    borderRadius: 8,
    marginTop: 15,
  },
  deleteButton: {
    backgroundColor: '#FF5733',
    borderRadius: 8,
    marginTop: 10,
  },
})

export default EditExercise
