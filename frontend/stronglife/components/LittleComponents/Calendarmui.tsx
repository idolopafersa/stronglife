import React, {useState} from 'react';
import {Calendar, LocaleConfig} from 'react-native-calendars';

const Calendarmui = () => {
  const [selected, setSelected] = useState('');

  return (
    <Calendar view="day"
      onDayPress={day => {
        setSelected(day.dateString);
      }}
      markedDates={{
        [selected]: {selected: true, disableTouchEvent: true, selectedDotColor: 'orange'}
      }}
    />
  );
};

export default Calendarmui;