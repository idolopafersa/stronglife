import { StyleSheet, View } from 'react-native'
import WeeklyCalendar from '../../components/WeeklyCalendar'
import DayMeal from '../../components/DayMeal'
import { useState } from 'react'
import { Text } from '@rneui/themed'

export function Kitchen() {
  const [kcal, setkcal] = useState<number>(0)

  return (
    <View style={styles.container}>
      <WeeklyCalendar />
      <Text>Total Calories: {kcal} kcal</Text>
      <View style={styles.mealContainer}>
        <DayMeal kitchenKcal={setkcal} />
      </View>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    gap: 10,
  },
  mealContainer: {
    width: '90%', // Centered width
    justifyContent: 'center',
    alignItems: 'center',
    padding: 10,
  },
})
