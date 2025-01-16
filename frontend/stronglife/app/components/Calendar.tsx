import React, { useState, useEffect } from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import { StyleSheet } from 'react-native';
import { Button } from 'react-native';
import dayjs from 'dayjs';
import CheckIcon from 'react-native-vector-icons/MaterialIcons'; 
import Arrow from 'react-native-vector-icons/MaterialIcons'; 
import Icon from "react-native-vector-icons/FontAwesome5"
import 'dayjs/locale/es'; 

dayjs.locale('es');

const Calendar = () => {
  const [completedDays, setCompletedDays] = useState<{ [key: string]: boolean }>({});
  const [streak, setStreak] = useState(0);
  const [currentWeek, setCurrentWeek] = useState<string[]>([]);
  const currentDayIndex = new Date().getDay();

  useEffect(() => {
    const week = getCurrentWeek();
    setCurrentWeek(week);
  }, []);

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

  const handleDayPress = (dateString: string) => {
    const updatedCompletedDays = { ...completedDays, [dateString]: !completedDays[dateString] };
    setCompletedDays(updatedCompletedDays);
    setStreak(calculateStreak(updatedCompletedDays));
  };

  const calculateStreak = (completedDays: { [key: string]: boolean }) => {
    const completedDates = Object.keys(completedDays).filter(date => completedDays[date]);
    completedDates.sort((a, b) => dayjs(a).valueOf() - dayjs(b).valueOf());
    
    let maxStreak = 0;
    let currentStreak = 0;
    let previousDate: dayjs.Dayjs | null = null;

    for (const dateString of completedDates) {
      const currentDate = dayjs(dateString);
      if (previousDate && currentDate.diff(previousDate, 'day') === 1) {
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
      <View style={styles.week}>
        {currentWeek.map((day, index) => (
          <TouchableOpacity
            key={index}
            onPress={() => handleDayPress(day)}
            style={[
              styles.dayContainer,
              index === currentDayIndex ? styles.currentDay : {},
              completedDays[day] ? styles.completed : styles.notCompleted,
            ]}
          >
            <Text style={[styles.dayText, completedDays[day] ? styles.completedText : styles.notCompletedText]}>
              <Text>{dayjs(day).date()}</Text>
              <Text>{'\n'}</Text>
              <Text>{dayjs(day).format('ddd')}</Text>
            </Text>
            {completedDays[day] && (
              <CheckIcon name="check" size={15} color="white" style={styles.checkIcon} />
            )}
          </TouchableOpacity>
        ))}
      </View>
      <View>
        <Text style={styles.streakText}>Racha: {streak} días {streak >= 3 && <Icon name="fire" color="tomato" />}</Text>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  week: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    alignItems: 'center',
  },
  container: {
    backgroundColor: '#1C1C1E',
    borderRadius: 20, 
    padding: '7%',
  },
  dayContainer: {
    width: 40,
    height: 60,
    borderRadius: 20,
    justifyContent: 'center',
    alignItems: 'center',
    marginHorizontal: 5,
    padding: 5,
  },
  currentDay: {
    backgroundColor: '#orange', 
  },
  completed: {
    backgroundColor: '#A5D32D', 
  },
  notCompleted: {
    backgroundColor: '#1C1C1E', 
    borderWidth: 1,
    borderColor: '#D3D3D3', 
  },
  dayText: {
    fontSize: 12,
    fontWeight: 'bold',
  },
  completedText: {
    color: 'white',
  },
  notCompletedText: {
    color: 'white',
  },
  checkIcon: {
    position: 'absolute',
    bottom: 5,
  },
  streakText: {
    marginTop: 20,
    fontSize: 16,
    fontWeight: 'bold',
    color: 'white',
    padding: 10,
  },
});

export default Calendar;