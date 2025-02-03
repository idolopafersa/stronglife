import React, { useState, useEffect } from 'react'
import {
  View,
  FlatList,
  StyleSheet,
  Text,
  ActivityIndicator,
} from 'react-native'
import { ListItem, SearchBar, Button } from '@rneui/themed'
import { useNavigation } from '@react-navigation/native'
import { Meal } from '../../../interfaces/Meal'
import { GetAllMeals } from '../../../services/MealServices'

const MealList = () => {
  const navigation = useNavigation<any>()
  const [meals, setMeals] = useState<Meal[]>([])
  const [search, setSearch] = useState('')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    fetchMeals()
  }, [])

  const fetchMeals = async () => {
    try {
      setLoading(true)
      setError(null) // Resetear error antes de cargar

      const data = await GetAllMeals()

      if (!data || !Array.isArray(data)) {
        setMeals([]) // Si devuelve null, mantenemos un array vacío
      } else {
        setMeals(data)
      }
    } catch (error) {
      console.error('Error fetching meals:', error)
      setError('Error al cargar las comidas. Intenta nuevamente.')
    } finally {
      setLoading(false)
    }
  }

  const handleSearch = (text: string) => {
    setSearch(text)
  }

  const handleSelectMeal = (meal: Meal) => {
    navigation.navigate('EditMeal', { meal })
  }

  const handleCreate = () => {
    navigation.navigate('EditMeal', { meal: null })
  }

  const filteredMeals =
    meals?.filter((meal) =>
      meal.name.toLowerCase().includes(search.toLowerCase())
    ) || []

  const renderItem = ({ item }: { item: Meal }) => (
    <ListItem bottomDivider onPress={() => handleSelectMeal(item)}>
      <ListItem.Content>
        <ListItem.Title>{item.name}</ListItem.Title>
      </ListItem.Content>
      <ListItem.Chevron />
    </ListItem>
  )

  return (
    <View style={styles.container}>
      <Button
        title="Crear Comida"
        onPress={handleCreate}
        buttonStyle={styles.createButton}
      />

      <SearchBar
        placeholder="Buscar Comidas"
        value={search}
        onChangeText={handleSearch}
        containerStyle={styles.searchBar}
        inputContainerStyle={styles.inputContainer}
      />

      {loading ? (
        <ActivityIndicator size="large" color="#4CAF50" />
      ) : error ? (
        <Text style={styles.errorText}>{error}</Text>
      ) : meals.length === 0 ? (
        <Text style={styles.emptyText}>Aún no tienes comidas.</Text>
      ) : (
        <FlatList
          data={filteredMeals}
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

export default MealList
