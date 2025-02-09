import React, { useState, useEffect } from 'react'
import { View, StyleSheet, Text, FlatList } from 'react-native'
import { Exercise } from '../../../interfaces/Exercise'
import { SearchBar, Button, ListItem, Icon } from '@rneui/themed'
import {
  AddRoutineExercise,
  DeleteRoutineExercise,
  GetRoutineExercises,
} from '../../../services/ExercisesRoutines'
import { GetAllExercises } from '../../../services/ExerciseServices'

const EditExercisesInRoutine = ({ route, navigation }: any) => {
  const routine_id = route.params?.routine_id
  const [userExercises, setUserExercises] = useState<Exercise[]>([]) // All exercises for the user
  const [exercises, setExercises] = useState<Exercise[]>([]) // Exercises already in the routine
  const [search, setSearch] = useState('') // Search query

  useEffect(() => {
    fetchExercises()
  }, [routine_id])

  // Fetch exercises for a given routine and all available exercises
  const fetchExercises = async () => {
    try {
      const userdata = await GetAllExercises() // Get all exercises for the user
      setUserExercises(userdata || [])
      const data = await GetRoutineExercises(routine_id) // Get exercises already in the routine
      setExercises(data || [])
    } catch (error) {
      console.error('Error fetching exercises:', error)
    }
  }

  // Add exercise to the routine
  const handleAddToRoutine = async (exercise_id: number) => {
    try {
      await AddRoutineExercise(routine_id, exercise_id)
      const exerciseToAdd = userExercises.find(
        (exercise) => exercise.id === exercise_id
      )
      if (exerciseToAdd) {
        setExercises((prevExercises) => [...prevExercises, exerciseToAdd]) // si se ha añadido, actualizamos estado
      }
    } catch (error) {
      console.error('Error adding exercise to routine:', error)
    }
  }

  // Delete exercise from the routine
  const handleDelete = async (exercise_id: number) => {
    try {
      await DeleteRoutineExercise(routine_id, exercise_id) // Call API to delete exercise from routine
      setExercises(exercises.filter((exercise) => exercise.id !== exercise_id)) // Remove exercise from state
    } catch (error) {
      console.error('Error deleting exercise:', error)
    }
  }

  // Filter exercises: Search through exercises that are not in the routine
  const filteredExercises = userExercises.filter(
    (exercise) =>
      !exercises.some((ex) => ex.id === exercise.id) &&
      exercise.name.toLowerCase().includes(search.toLowerCase())
  )

  const renderAllExercises = ({ item }: { item: Exercise }) => (
    <ListItem bottomDivider>
      <ListItem.Content>
        <ListItem.Title>{item.name}</ListItem.Title>
      </ListItem.Content>
      <Icon
        name="check"
        type="material"
        color="green"
        size={24}
        onPress={() => handleAddToRoutine(item.id)}
      />
    </ListItem>
  )

  const renderRoutineExercises = ({ item }: { item: Exercise }) => (
    <ListItem bottomDivider>
      <ListItem.Content>
        <ListItem.Title>{item.name}</ListItem.Title>
      </ListItem.Content>

      <Icon
        name="close"
        type="material"
        color="red"
        size={24}
        onPress={() => handleDelete(item.id)} // Trigger deletion
      />
    </ListItem>
  )

  //Im sure theres a more efficient way to do this
  return (
    <View style={styles.container}>
      <SearchBar
        placeholder="Buscar Ejercicios"
        value={search}
        onChangeText={setSearch}
        containerStyle={styles.searchBar}
        inputContainerStyle={styles.inputContainer}
      />

      {search ? (
        <>
          {filteredExercises.length === 0 ? (
            <Text style={styles.emptyText}>
              No hay ejercicios disponibles para esta búsqueda.
            </Text>
          ) : (
            <FlatList
              data={filteredExercises}
              renderItem={renderAllExercises}
              keyExtractor={(item) => item.id.toString()}
              contentContainerStyle={styles.list}
            />
          )}
        </>
      ) : (
        // Cuando no se busca
        <>
          {exercises.length === 0 ? (
            <Text style={styles.emptyText}>
              No hay ejercicios en esta rutina.
            </Text>
          ) : (
            <FlatList
              data={exercises}
              renderItem={renderRoutineExercises}
              keyExtractor={(item) => item.id.toString()}
              contentContainerStyle={styles.list}
            />
          )}
        </>
      )}
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 20,
    backgroundColor: '#f9f9f9',
  },
  searchBar: {
    width: '100%',
    marginBottom: 10,
  },
  inputContainer: {
    backgroundColor: '#ffffff',
  },
  list: {
    flexGrow: 1,
  },
  emptyText: {
    textAlign: 'center',
    fontSize: 16,
    marginTop: 20,
    color: '#666',
  },
})

export default EditExercisesInRoutine
