import React, { useState } from 'react';
import { StyleSheet, Text } from 'react-native';
import { ThemedText } from '@/components/ThemedText';
import { ThemedView } from '@/components/ThemedView';
import { useUser } from '@/app/context/UserContext';
import WeeklyCalendar from '@/components/WeeklyCalendar';
import moment from 'moment';


export default function Gym() {
  const currentUser = useUser();
  const [dayselected, setDaySelected] = useState(moment().format('YYYY-MM-DD')); // Initialize with today's date


  

  return (
    <ThemedView style={styles.container}>
      <ThemedView style={styles.calendarPosition}>
      <WeeklyCalendar selectedDate={dayselected} onDateChange={setDaySelected} />
      </ThemedView>
      
      <ThemedText type="title">
        Welcome to the Home Screen! {currentUser.username}, Today is {dayselected}
      </ThemedText>
      <Text>This is where you can find various components and features of the app.</Text>
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 16,
  },
  calendarPosition: {
    flex: 1,
    justifyContent: 'flex-start', // Align calendar at the top
    alignItems: 'center', // Center horizontally
    width: '100%', // Ensure it takes full width if needed,
    
  }
});


