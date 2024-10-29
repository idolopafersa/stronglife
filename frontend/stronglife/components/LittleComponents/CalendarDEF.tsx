import React, { useState, useEffect } from 'react';
import { View, Text } from 'react-native';
import { Calendar, LocaleConfig } from 'react-native-calendars';
import { StyleSheet } from 'react-native';
import { Button } from 'react-native';
import type {} from '@mui/x-date-pickers/AdapterDayjs';
import { DatePicker } from '@mui/x-date-pickers'


// Configurar el idioma en español si lo deseas
LocaleConfig.locales['es'] = {
  monthNames: ['Enero', 'Febrero', 'Marzo', 'Abril', 'Mayo', 'Junio', 'Julio', 'Agosto', 'Septiembre', 'Octubre', 'Noviembre', 'Diciembre'],
  monthNamesShort: ['Ene', 'Feb', 'Mar', 'Abr', 'May', 'Jun', 'Jul', 'Ago', 'Sep', 'Oct', 'Nov', 'Dic'],
  dayNames: ['Domingo', 'Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado'],
  dayNamesShort: ['D', 'L', 'M', 'M', 'J', 'V', 'S']
};
LocaleConfig.defaultLocale = 'es';

const CombinedCalendar = () => {
  const [completedDays, setCompletedDays] = useState<{ [key: string]: boolean }>({});
  const [streak, setStreak] = useState(0);
  const [currentWeek, setCurrentWeek] = useState<string[]>([]);
  const [weekData, setWeekData] = useState<{ date: string, completed: boolean }[]>([]);

  useEffect(() => {
    const week = getCurrentWeek();
    setCurrentWeek(week);
    setWeekData(getWeekData(week, completedDays));
  }, [completedDays]);

  const getCurrentWeek = (): string[] => {
    const today = new Date();
    const startOfWeek = today.getDate() - today.getDay();
    const week = [];
    for (let i = 0; i < 7; i++) {
      const date = new Date(today.setDate(startOfWeek + i));
      week.push(date.toISOString().split('T')[0]);
    }
    return week;
  };
  const addWeek = (week: string[]) => {
    const newweek = week;
    
    for (let i = 0; i < 7; i++) {
      const d= new Date(week[i]);
    d.setDate(d.getDate() + d.getDay() + 7);
    d.toDateString();
    newweek[i]= d.toISOString().split('T')[0];
    }
    return newweek;
  };
  const getWeekData = (week: string[], completedDays: { [key: string]: boolean }) => {
    const d = new Date();
    return week.map(date => ({
      date,
      completed: completedDays[date]
    }));
  };

  const handleDayPress = (day: { dateString: string }) => {
    const dateString = day.dateString;
    const updatedCompletedDays = { ...completedDays, [dateString]: !completedDays[dateString] };
    setCompletedDays(updatedCompletedDays);
    setStreak(calculateStreak(updatedCompletedDays));
  };

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

  return (
    <View style={styles.container}>
      <Calendar 
        style={styles.calendar}
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
      <View>
        {/* <Button onClick={addWeek(currentWeek)}> */}
          {/* addWeek */}
        {/* </Button> */}
        <Text>
        {/* currentWeek: {currentWeek.join(', ')} */}
        Semana +1: {addWeek(currentWeek).join(', ')}
        </Text>
        
      </View>
      <View style={styles.weekContainer}>
        {weekData.map(day => (
          <Text key={day.date} style={styles.weekText}>
            {day.date}: {day.completed ? 'Completado' : 'No completado'}
          </Text>
        ))}
      </View>
    </View>
  );
};
const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 10,
    backgroundColor: '#fff',
  },
  calendar: {
    marginBottom: 10,
  },
  streakContainer: {
    alignItems: 'center',
    marginVertical: 10,
  },
  streakText: {
    fontSize: 18,
    fontWeight: 'bold',
  },
  weekContainer: {
    marginTop: 10,
  },
  weekText: {
    fontSize: 16,
    marginVertical: 2,
  },
});

export default CombinedCalendar;
