import React from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';
import Anothercalendar from './LittleComponents/Anothercalendar';
import  Calendar  from './LittleComponents/Calendar';
import LabelBottomNavigation from './LittleComponents/LabelBottomNavigation';
const HelpScreen = () => {
  return (
    <View >
      <Anothercalendar/>
      <Calendar />
      
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    padding: 16,
    flexGrow: 1,
    justifyContent: 'center',
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 16,
  },
  subtitle: {
    fontSize: 18,
    marginBottom: 12,
  },
  question: {
    fontSize: 16,
    fontWeight: 'bold',
    marginTop: 12,
  },
  answer: {
    fontSize: 16,
    marginBottom: 12,
  },
});

export default HelpScreen;