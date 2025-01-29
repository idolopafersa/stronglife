import React from 'react';
import { StyleSheet, View, Text } from 'react-native';
import moment from 'moment';

const WeeklyCalendar = ({ selectedDate, onDateChange }) => {
  const startOfWeek = moment(selectedDate).startOf('week'); // Get the start of the week

  // Generate the days of the week
  const weekDays = Array.from({ length: 7 }, (_, i) =>
    startOfWeek.clone().add(i, 'days').format('YYYY-MM-DD')
  );

  // Handle day press
  const handleDayPress = (day) => {
    onDateChange(day); // Call onDateChange to update the selected day in the parent component
  };

  return (
    <View style={styles.calendarContainer}>
      {weekDays.map((day) => {
        const isSelected = selectedDate === day; // Check if the day is selected
        const dayName = moment(day).format('dd'); // Get the day name (L, M, X, etc.)

        return (
          <View
            key={day}
            style={[
              styles.dayContainer,
              isSelected && styles.selectedDay,
            ]}
          >
            <Text
              style={[
                styles.dayText,
                isSelected && styles.selectedDayText,
              ]}
              onPress={() => handleDayPress(day)}
            >
              {dayName}
            </Text>
          </View>
        );
      })}
    </View>
  );
};

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
});

export default WeeklyCalendar;
