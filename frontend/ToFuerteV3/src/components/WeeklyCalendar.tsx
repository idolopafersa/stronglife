import React, { useState, useEffect } from 'react'
import { StyleSheet, View, Text } from 'react-native'
import moment from 'moment'
import { useUser } from '../context/UserContext'

const WeeklyCalendar = () => {
  const { currentUser, setCurrentUser } = useUser()

  const [selectedDate, setSelectedDate] = useState(currentUser.dayselected)

  useEffect(() => {
    setSelectedDate(currentUser.dayselected) // Update when context changes
  }, [currentUser.dayselected])

  const startOfWeek = moment(selectedDate).startOf('week')

  const weekDays = Array.from({ length: 7 }, (_, i) =>
    startOfWeek.clone().add(i, 'days').format('YYYY-MM-DD')
  )

  const handleDayPress = (day: string) => {
    setSelectedDate(day) // Update local state for instant UI update
    setCurrentUser((prev) => ({
      ...prev,
      dayselected: day, // Update global context
    }))
  }

  return (
    <View style={styles.calendarContainer}>
      {weekDays.map((day) => {
        const isSelected = selectedDate === day
        const dayName = moment(day).format('dd')

        return (
          <View
            key={day}
            style={[styles.dayContainer, isSelected && styles.selectedDay]}
          >
            <Text
              style={[styles.dayText, isSelected && styles.selectedDayText]}
              onPress={() => handleDayPress(day)}
            >
              {dayName}
            </Text>
          </View>
        )
      })}
    </View>
  )
}

const styles = StyleSheet.create({
  calendarContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    width: 350,
    height: 60,
    borderWidth: 1,
    borderColor: 'green',
    borderRadius: 10,
    paddingHorizontal: 10,
    backgroundColor: 'white',
  },
  dayContainer: {
    justifyContent: 'center',
    alignItems: 'center',
    width: 40,
    height: 40,
    borderRadius: 20,
  },
  selectedDay: {
    backgroundColor: 'orange',
  },
  dayText: {
    fontSize: 16,
    color: 'black',
  },
  selectedDayText: {
    color: 'white',
    fontWeight: 'bold',
  },
})

export default WeeklyCalendar
