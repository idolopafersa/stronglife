import React, { useState } from 'react';
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';
import CheckIcon from 'react-native-vector-icons/MaterialIcons'; // Asegúrate de instalar react-native-vector-icons
const daysOfWeek = ['LU', 'MA', 'MI', 'JU', 'VI', 'SÁ', 'DO'];

const AnotherCalendar: React.FC = () => {
  const [completedDays, setCompletedDays] = useState<boolean[]>(Array(7).fill(false));
  const [currentDayIndex, setCurrentDayIndex] = useState<number>(3); // Cambia esto según el día actual (0 para lunes, 6 para domingo)
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
            index === currentDayIndex ? styles.currentDay : {},
            completedDays[index] ? styles.completed : styles.notCompleted,
          ]}
        >
          <Text style={[styles.dayText, completedDays[index] ? styles.completedText : styles.notCompletedText]}>
            {day}
          </Text>
          {completedDays[index] && (
            <CheckIcon name="check" size={15} color="white" style={styles.checkIcon} />
          )}
        </TouchableOpacity>
      ))}
      <Text style={styles.streakText}>Racha: {streak} días</Text>
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
    width: 40,
    height: 60,
    borderRadius: 20,
    justifyContent: 'center',
    alignItems: 'center',
    marginHorizontal: 5,
  },
  currentDay: {
    backgroundColor: '#FF8C00', // Color naranja para el día actual
  },
  completed: {
    backgroundColor: '#00C781', // Color verde para días completados
  },
  notCompleted: {
    backgroundColor: '#F5F5F5', // Fondo gris claro para días no completados
    borderWidth: 1,
    borderColor: '#D3D3D3', // Borde gris claro
  },
  dayText: {
    fontSize: 12,
    fontWeight: 'bold',
  },
  completedText: {
    color: 'white',
  },
  notCompletedText: {
    color: 'black',
  },
  checkIcon: {
    position: 'absolute',
    bottom: 5,
  },
  streakText: {
    marginTop: 20,
    fontSize: 16,
    fontWeight: 'bold',
    color: 'black',
  },
});

export default AnotherCalendar;
