import React, { useState, useEffect } from 'react'
import {
  View,
  FlatList,
  Text,
  StyleSheet,
  ActivityIndicator,
} from 'react-native'
import { ListItem, SearchBar, Button } from '@rneui/themed'
import { GetAllExercises } from '../../../services/ExerciseServices'
import { Exercise } from '../../../interfaces/Exercise'
import { useNavigation } from '@react-navigation/native'

const ExerciseList = () => {
  const navigation = useNavigation<any>()
  const [exercises, setExercises] = useState<Exercise[]>([]) // Type-safe state for exercises
  const [search, setSearch] = useState('') // State to hold search query
  const [loading, setLoading] = useState(true) // Loading state for exercises
  const [error, setError] = useState<string | null>(null) // Error state for handling errors

  useEffect(() => {
    fetchExercises()
  }, [])

  const fetchExercises = async () => {
    try {
      setLoading(true)
      setError(null) // Reset error before fetching

      const data = await GetAllExercises()

      if (!data || !Array.isArray(data)) {
        setExercises([]) // If the data is null or not an array, set an empty array
      } else {
        setExercises(data)
      }
    } catch (error) {
      console.error('Error fetching exercises:', error)
      setError('Error al cargar los ejercicios. Intenta nuevamente.')
    } finally {
      setLoading(false)
    }
  }

  const handleSearch = (text: string) => {
    setSearch(text)
  }

  const handleSelectExercise = (exercise: Exercise) => {
    navigation.navigate('EditExercise', { exercise }) // Navigate to edit screen with selected exercise
  }

  const handleCreate = () => {
    navigation.navigate('EditExercise', {}) // Navigate to create a new exercise
  }

  // Filter exercises based on search query
  const filteredExercises =
    exercises?.filter((exercise) =>
      exercise.name.toLowerCase().includes(search.toLowerCase())
    ) || []

  // Render List Item
  const renderItem = ({ item }: { item: Exercise }) => (
    <ListItem bottomDivider onPress={() => handleSelectExercise(item)}>
      <ListItem.Content>
        <ListItem.Title>{item.name}</ListItem.Title>
      </ListItem.Content>
      <ListItem.Chevron />
    </ListItem>
  )

  return (
    <View style={styles.container}>
      <Button
        title="Crear Ejercicio"
        onPress={handleCreate}
        buttonStyle={styles.createButton}
      />

      <SearchBar
        placeholder="Buscar Ejercicios"
        value={search}
        onChangeText={handleSearch}
        containerStyle={styles.searchBar}
        inputContainerStyle={styles.inputContainer}
      />

      {loading ? (
        <ActivityIndicator size="large" color="#4CAF50" />
      ) : error ? (
        <Text style={styles.errorText}>{error}</Text>
      ) : exercises.length === 0 ? (
        <Text style={styles.emptyText}>Aún no tienes ejercicios.</Text>
      ) : (
        <FlatList
          data={filteredExercises}
          renderItem={renderItem}
          keyExtractor={(item) => item.id.toString()}
          contentContainerStyle={styles.list}
        />
      )}
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 10,
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
  createButton: {
    backgroundColor: '#4CAF50',
    marginBottom: 20,
  },
  errorText: {
    textAlign: 'center',
    color: 'red',
    fontSize: 16,
    marginTop: 20,
  },
  emptyText: {
    textAlign: 'center',
    fontSize: 16,
    marginTop: 20,
    color: '#666',
  },
})

export default ExerciseList
