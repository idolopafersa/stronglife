import dayjs from 'dayjs';
import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

interface WeekdaysProps {
  week: string[];
  diasCompletados: string[];
}

const Weekdays: React.FC<WeekdaysProps> = ({ week, diasCompletados }) => {
  return (
    <View style={styles.weekContainer}>
      {week.map((day) => (
        <Text
          key={day}
          style={[
            styles.dayText,
            diasCompletados.includes(day) && styles.completedDayText,
          ]}
        >
          {dayjs(day).format('ddd')}
        </Text>
      ))}
    </View>
  );
};

const styles = StyleSheet.create({
  weekContainer: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    marginVertical: 10,
  },
  dayText: {
    fontSize: 16,
    color: 'black',
  },
  completedDayText: {
    color: 'green',
    fontWeight: 'bold',
  },
});

export default Weekdays;