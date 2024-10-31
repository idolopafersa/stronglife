import React, { useState, useEffect } from 'react';
import { View, Text,TouchableOpacity } from 'react-native';
import { Calendar, LocaleConfig } from 'react-native-calendars';
import { StyleSheet } from 'react-native';
import { Button } from 'react-native';
import type {} from '@mui/x-date-pickers/AdapterDayjs';
import { DatePicker } from '@mui/x-date-pickers'
import dayjs from 'dayjs';
import CheckIcon from 'react-native-vector-icons/MaterialIcons'; 
import Arrow from 'react-native-vector-icons/MaterialIcons'; 
import Star from 'react-native-vector-icons/MaterialIcons';
import 'dayjs/locale/es'; 

dayjs.locale('es');
const FinalCalendar = () => {
  const [completedDays, setCompletedDays] = useState<{ [key: string]: boolean }>({});
  const [streak, setStreak] = useState(0);
  const [currentWeek, setCurrentWeek] = useState<string[]>([]);
  const [weekData, setWeekData] = useState<{ date: string, completed: boolean }[]>([]);
  const currentDayIndex = new Date().getDay();

  useEffect(() => {
    const week = getCurrentWeek();
    setCurrentWeek(week);
    setWeekData(getWeekData(week, completedDays));
  }, []);

  useEffect(() => {
    setWeekData(getWeekData(currentWeek, completedDays));
  }, [currentWeek, completedDays]);

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
    const newweek = [...week];
    
    for (let i = 0; i < 7; i++) {
      const d = dayjs(week[i]).add(1, 'week');
      newweek[i] = d.toISOString().split('T')[0];
    }
    return newweek;
  };
  const subWeek = (week: string[]) => {
    const newweek = [...week];
    
    for (let i = 0; i < 7; i++) {
      const d = dayjs(week[i]).subtract(1, 'week');
      newweek[i] = d.toISOString().split('T')[0];
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
    <View>
    <View style={{ flexDirection: 'row', justifyContent: 'space-between', marginHorizontal: 10 }}>
        <TouchableOpacity onPress={() => setCurrentWeek(subWeek(currentWeek))}>
            <Arrow name="arrow-back" size={32} color="black" />
        </TouchableOpacity>
        <Text style={styles.tituloWeek}>{dayjs(currentWeek[0]).format('MMM YYYY')}</Text>
        <TouchableOpacity onPress={() => setCurrentWeek(addWeek(currentWeek))}>
            <Arrow name="arrow-forward" size={32} color="black" />
        </TouchableOpacity>
    </View>
      
    <View style={styles.container}>
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
      <Text style={styles.streakText}>Racha: {streak} días</Text>
      <Star/>
    </View>
    </View>
  );
};

const styles = StyleSheet.create({
    container: {
      flexDirection: 'row',
      justifyContent: 'space-around',
      padding: 'auto',
      alignItems: 'center',
      backgroundColor: 'grey',
      borderRadius: 20, 
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
      backgroundColor: '#FF8C00', 
    },
    completed: {
      backgroundColor: '#00C781', 
    },
    notCompleted: {
      backgroundColor: '#F5F5F5', 
      borderWidth: 1,
      borderColor: '#D3D3D3', 
    },
    dayText: {
      fontSize: 12,
      fontWeight: 'bold',
    },
    completedText: {
      color: 'white',
    },tituloWeek: {
        color: 'black',
        fontSize: 20,
        fontWeight: 'bold',
        textAlign: 'center',
        marginBottom: 10,
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
  

export default FinalCalendar;
