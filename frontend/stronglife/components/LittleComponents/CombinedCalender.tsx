import React, { useState } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { Calendar, LocaleConfig } from 'react-native-calendars';
import WhatshotIcon from '@mui/icons-material/Whatshot';

// Configurar el idioma en español si lo deseas
LocaleConfig.locales['es'] = {
  monthNames: ['Enero', 'Febrero', 'Marzo', 'Abril', 'Mayo', 'Junio', 'Julio', 'Agosto', 'Septiembre', 'Octubre', 'Noviembre', 'Diciembre'],
  monthNamesShort: ['Ene', 'Feb', 'Mar', 'Abr', 'May', 'Jun', 'Jul', 'Ago', 'Sep', 'Oct', 'Nov', 'Dic'],
  dayNames: ['Domingo', 'Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado'],
  dayNamesShort: ['D', 'L', 'M', 'M', 'J', 'V', 'S']
};
LocaleConfig.defaultLocale = 'es';

const CombinedCalender = () => {
  const [completedDays, setCompletedDays] = useState<{ [key: string]: boolean }>({});
  const [streak, setStreak] = useState<number>(0);

  const calculateStreak = (completedDays: { [key: string]: boolean }) => {
    const dates = Object.keys(completedDays).filter(date => completedDays[date]);
    dates.sort((a, b) => new Date(a).getTime() - new Date(b).getTime());
    
    let maxStreak = 0;
    let currentStreak = 0;
    let previousDate: Date | null = null;

    for (const dateString of dates) {
      const currentDate = new Date(dateString);
      if (
        previousDate &&
        (currentDate.getTime() - previousDate.getTime() === 86400000)
      ) {
        currentStreak++;
      } else {
        currentStreak = 1; 
      }
      maxStreak = Math.max(maxStreak, currentStreak);
      previousDate = currentDate;
    }

    return maxStreak;
  };

  const handleDayPress = (day: { dateString: string }) => {
    const dateString = day.dateString;
    const updatedCompletedDays = { ...completedDays, [dateString]: !completedDays[dateString] };
    setCompletedDays(updatedCompletedDays);
    setStreak(calculateStreak(updatedCompletedDays));
  };

  return (
    <View style={styles.container}>
      <Calendar style={styles.calendar}
        onDayPress={handleDayPress}
        markedDates={{
          ...Object.keys(completedDays).reduce((acc, date) => {
            acc[date] = { selected: completedDays[date], selectedColor: 'green' };
            return acc;
          }, {} as { [key: string]: { selected: boolean; selectedColor: string } })
        }}
      />
      <View style={styles.streakContainer}>
        
        <Text style={styles.streakText}>Racha: {streak} días</Text>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    padding: 10,
  },
  streakContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginTop: 20,
    justifyContent: 'center',
  },
  icon: {
    marginRight: 8,
    color: 'orange',
  },calendar:{
    backgroundColor:'grey',
  },
  streakText: {
    fontSize: 20,
    fontWeight: 'bold',
  },
});

export default CombinedCalender;
