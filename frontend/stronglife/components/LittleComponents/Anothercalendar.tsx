import React, { useState } from 'react';
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';
import WhatshotIcon from '@mui/icons-material/Whatshot';
const daysOfWeek = ['L', 'M', 'M', 'J', 'V', 'S', 'D'];

const Anothercalendar: React.FC = () => {
  const [completedDays, setCompletedDays] = useState<boolean[]>(Array(7).fill(false));
  const [streak, setStreak] = useState<number>(0);

  const calculateStreak = (updatedCompletedDays: boolean[]) => {
    let maxStreak = 0;
    let currentStreak = 0;

    for (const dayCompleted of updatedCompletedDays) {
      if (dayCompleted) {
        currentStreak++;
        maxStreak = Math.max(maxStreak, currentStreak);
      } else {
        currentStreak = 0; // Reinicia la racha cuando encuentra un día no completado
      }
    }

    return maxStreak;
  };

  const handleDayPress = (index: number) => {
    const updatedCompletedDays = [...completedDays];
    updatedCompletedDays[index] = !updatedCompletedDays[index]; // Alterna el estado del día seleccionado

    setCompletedDays(updatedCompletedDays);
    setStreak(calculateStreak(updatedCompletedDays)); // Recalcula la racha de días seguidos
  };

  return (
    <View style={styles.container}>
      {daysOfWeek.map((day, index) => (
        <TouchableOpacity
          key={index}
          onPress={() => handleDayPress(index)}
          style={[
            styles.dayContainer,
            completedDays[index] ? styles.completed : styles.notCompleted,
          ]}
        >
          <Text style={styles.dayText}>{day}</Text>
        </TouchableOpacity>
      ))}
      <Text style={styles.streakText}>{streak} </Text>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    padding: 10,
    alignItems: 'center',
  },
  dayContainer: {
    padding: 20,
    borderRadius: 10,
    borderWidth: 2,
    borderColor: 'black',
  },
  completed: {
    backgroundColor: 'green',
  },
  notCompleted: {
    backgroundColor: 'darkgrey',
  },
  dayText: {
    color: 'white',
    fontSize: 18,
    fontWeight: 'bold',
  },
  streakText: {
    marginTop: 20,
    fontSize: 20,
    fontWeight: 'bold',
    padding: 10,
  },
});

export default Anothercalendar;
