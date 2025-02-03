import React, { useState } from 'react'
import { View, StyleSheet } from 'react-native'
import { Button, Card, Input } from '@rneui/themed'
import { PutMeal, DelMeal, PostMeal } from '../../../services/MealServices'
import { useNavigation, RouteProp } from '@react-navigation/native'
import { Meal } from '../../../interfaces/Meal'

type EditMealProps = {
  route: RouteProp<{ params: { meal?: Meal } }, 'params'>
}

const EditMeal: React.FC<EditMealProps> = ({ route }) => {
  const navigation = useNavigation()
  const meal = route.params?.meal
  const isCreating = !meal

  const [name, setName] = useState(meal?.name || '')
  const [description, setDescription] = useState(meal?.description || '')
  const [calories, setCalories] = useState(meal?.calories?.toString() || '')
  const [proteins, setProteins] = useState(meal?.proteins?.toString() || '')
  const [fats, setFats] = useState(meal?.fats?.toString() || '')
  const [carbs, setCarbs] = useState(meal?.carbs?.toString() || '')

  const handleSave = async () => {
    if (!name.trim() || !description.trim()) {
      return
    }

    try {
      if (isCreating) {
        await PostMeal(
          name,
          description,
          parseFloat(calories),
          parseFloat(proteins),
          parseFloat(fats),
          parseFloat(carbs)
        )
      } else {
        await PutMeal(
          meal.id,
          name,
          description,
          parseFloat(calories),
          parseFloat(proteins),
          parseFloat(fats),
          parseFloat(carbs)
        )
      }
      navigation.navigate('MealList' as never)
    } catch (error) {
      console.error('Error saving meal:', error)
    }
  }

  const handleDelete = async () => {
    if (isCreating) return

    try {
      await DelMeal(meal.id)
      navigation.navigate('MealList' as never)
    } catch (error) {
      console.error('Error deleting meal:', error)
    }
  }

  return (
    <View style={styles.container}>
      <Card containerStyle={styles.card}>
        <Card.Title>{isCreating ? 'Create Meal' : 'Edit Meal'}</Card.Title>
        <Card.Divider />

        <Input
          label="Meal Name"
          value={name}
          onChangeText={setName}
          placeholder="Enter meal name"
        />
        <Input
          label="Description"
          value={description}
          onChangeText={setDescription}
          placeholder="Enter description"
          multiline
          numberOfLines={3}
        />
        <Input
          label="Calories"
          value={calories}
          onChangeText={setCalories}
          placeholder="Enter calories"
          keyboardType="numeric"
        />
        <Input
          label="Proteins"
          value={proteins}
          onChangeText={setProteins}
          placeholder="Enter proteins"
          keyboardType="numeric"
        />
        <Input
          label="Fats"
          value={fats}
          onChangeText={setFats}
          placeholder="Enter fats"
          keyboardType="numeric"
        />
        <Input
          label="Carbs"
          value={carbs}
          onChangeText={setCarbs}
          placeholder="Enter carbs"
          keyboardType="numeric"
        />

        <Button
          title={isCreating ? 'Create Meal' : 'Save Changes'}
          onPress={handleSave}
          buttonStyle={styles.saveButton}
        />

        {!isCreating && (
          <Button
            title="Delete Meal"
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
    elevation: 5,
    shadowColor: '#000',
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

export default EditMeal
