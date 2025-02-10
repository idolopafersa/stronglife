import { useEffect, useState } from 'react'
import { useUser } from '../context/UserContext'
import { Meal } from '../interfaces/Meal'
import { GetDay } from '../services/DayServices'
import { Card, Text } from '@rneui/themed'
import { ImageBackground, View, FlatList } from 'react-native'
import { StyleSheet } from 'react-native'

const DayMeal = ({ kitchenKcal }: any) => {
  const { currentUser } = useUser()
  const [meals, setMeals] = useState<Meal[]>([])

  const fetchMeals = async () => {
    const response = await GetDay(currentUser.dayselected)
    setMeals(response.meals)

    const total = response.meals.reduce(
      (sum: any, meal: { calories: any }) => sum + meal.calories,
      0
    )

    kitchenKcal(total)
  }

  useEffect(() => {
    fetchMeals()
  }, [])

  const renderMealCard = ({ item }: { item: Meal }) => (
    <>
      <Card containerStyle={styles.cardContainer}>
        <ImageBackground
          source={{
            uri: `http://localhost:777/api/meal/getimage?meal_id=${item.id}`,
          }}
          style={styles.imageBackground}
          imageStyle={styles.imageRounded}
        >
          <View style={styles.overlay}>
            <Text style={styles.mealName}>{item.name}</Text>
            <Text style={styles.mealCalories}>{item.calories} kcal</Text>
          </View>
        </ImageBackground>
      </Card>
    </>
  )

  return (
    <View style={styles.container}>
      <FlatList
        data={meals.slice(0, 4)} // Only show 4 meals
        renderItem={renderMealCard}
        keyExtractor={(item) => item.id.toString()}
        numColumns={2} // Show 2 meals per row
        contentContainerStyle={styles.listContainer}
      />
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
  listContainer: {
    alignItems: 'center',
    justifyContent: 'center',
    paddingBottom: 20,
  },
  cardContainer: {
    borderRadius: 15,
    overflow: 'hidden',
    padding: 0,
    width: 140, // Smaller width
    height: 140, // Smaller height
    elevation: 3,
    margin: 10, // Spacing between cards
  },
  imageBackground: {
    width: '100%',
    height: '100%',
    justifyContent: 'flex-end',
  },
  imageRounded: {
    borderRadius: 15,
  },
  overlay: {
    backgroundColor: 'rgba(0, 0, 0, 0.6)',
    padding: 8,
    borderBottomLeftRadius: 15,
    borderBottomRightRadius: 15,
  },
  mealName: {
    fontSize: 14, // Smaller text
    fontWeight: 'bold',
    color: 'white',
    textAlign: 'center',
  },
  mealCalories: {
    fontSize: 12,
    color: 'white',
    textAlign: 'center',
  },
})

export default DayMeal
