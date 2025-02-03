import React, { useState, useEffect } from 'react'
import {
  View,
  FlatList,
  Text,
  StyleSheet,
  ActivityIndicator,
} from 'react-native'
import { ListItem, SearchBar, Button } from '@rneui/themed'
import { useNavigation } from '@react-navigation/native'
import { Routine } from '../../../interfaces/Routine'
import { GetAllRoutines } from '../../../services/RoutineServices'

const RoutineList = () => {
  const navigation = useNavigation<any>()
  const [routines, setRoutines] = useState<Routine[]>([])
  const [search, setSearch] = useState('')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    fetchRoutines()
  }, [])

  const fetchRoutines = async () => {
    try {
      setLoading(true)
      setError(null) // Resetear error antes de cargar

      const data = await GetAllRoutines()

      if (!data || !Array.isArray(data)) {
        setRoutines([]) // Si devuelve null, asignamos un array vacío
      } else {
        setRoutines(data)
      }
    } catch (error) {
      console.error('Error fetching routines:', error)
      setError('Error al cargar las rutinas. Intenta nuevamente.')
    } finally {
      setLoading(false)
    }
  }

  const handleSearch = (text: string) => {
    setSearch(text)
  }

  const handleSelectRoutine = (routine: Routine) => {
    navigation.navigate('EditRoutine', { routine })
  }

  const handleCreate = () => {
    navigation.navigate('EditRoutine', {})
  }

  const filteredRoutines =
    routines?.filter((routine) =>
      routine.name.toLowerCase().includes(search.toLowerCase())
    ) || []

  const renderItem = ({ item }: { item: Routine }) => (
    <ListItem bottomDivider onPress={() => handleSelectRoutine(item)}>
      <ListItem.Content>
        <ListItem.Title>{item.name}</ListItem.Title>
      </ListItem.Content>
      <ListItem.Chevron />
    </ListItem>
  )

  return (
    <View style={styles.container}>
      <Button
        title="Crear Rutina"
        onPress={handleCreate}
        buttonStyle={styles.createButton}
      />

      <SearchBar
        placeholder="Buscar Rutinas"
        value={search}
        onChangeText={handleSearch}
        containerStyle={styles.searchBar}
        inputContainerStyle={styles.inputContainer}
      />

      {loading ? (
        <ActivityIndicator size="large" color="#4CAF50" />
      ) : error ? (
        <Text style={styles.errorText}>{error}</Text>
      ) : routines.length === 0 ? (
        <Text style={styles.emptyText}>Aún no tienes rutinas.</Text>
      ) : (
        <FlatList
          data={filteredRoutines}
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

export default RoutineList
