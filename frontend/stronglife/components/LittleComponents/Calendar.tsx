import React, { useState } from 'react';
import { View, Text, TouchableOpacity, StyleSheet, Dimensions } from 'react-native';

const daysOfWeek = ['L', 'M', 'M', 'J', 'V', 'S', 'D'];
const { width } = Dimensions.get('window');

const Calendar: React.FC = () => {
  const [attendance, setAttendance] = useState<boolean[]>(Array(7).fill(false));

  const toggleAttendance = (index: number) => {
    const updatedAttendance = [...attendance];
    updatedAttendance[index] = !updatedAttendance[index];
    setAttendance(updatedAttendance);
  };

  return (
    <View style={styles.container}>
      {daysOfWeek.map((day, index) => (
        <TouchableOpacity
          key={index}
          onPress={() => toggleAttendance(index)}
          style={[
            styles.dayContainer,
            attendance[index] ? styles.attended : styles.notAttended,
          ]}
        >
          <Text style={styles.dayText}>{day}</Text>
        </TouchableOpacity>
      ))}
    </View>
  );
};

const styles = StyleSheet.create({
    container: {
      flexDirection: 'row',
      justifyContent: 'space-around',
      padding: 10,
    },
    dayContainer: {
        width: width * 0.12, 
      padding: 20,
      borderRadius: 10,
      borderWidth: 2,
      borderColor: 'black',
    },
    attended: {
      backgroundColor: 'orange',
    },
    notAttended: {
      backgroundColor: 'darkgrey',
    },
    dayText: {
      color: 'white',
      fontSize: 18,
      fontWeight: 'bold',
      width: width * 0.03, 
    },
  });

export default Calendar;