import React from 'react';
import { Text, View, StyleSheet, TouchableOpacity, Image } from 'react-native';
import Calendar from '../components/Calendar';
import Icon from "react-native-vector-icons/FontAwesome5";

const MiGimnasio = () => {
  return (
    <View style={styles.container}>
      <Calendar />
      <View style={styles.recommendedTrainingContainer}>
        <Text style={styles.text}>Entrenamiento Recomendado</Text>
        <View style={styles.trainingBox}>
          <Text style={styles.trainingName}>Push Workout</Text>
          <Text style={styles.trainingDetails}>6 Ejercicios</Text>
          <View style={styles.exerciseDetails}>
            <TouchableOpacity style={styles.startButton} onPress={() => {}}>
              <Text style={styles.buttonText}>Empezar</Text>
            </TouchableOpacity>
          </View>
          <View style={styles.timeContainer}>
            <Icon name="fire" size={20} color="tomato" style={styles.fireIcon} />
            <Text style={styles.timeText}>30 min</Text>
          </View>
        </View>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    backgroundColor: 'black',
    flex: 1,
    padding: 20,
  },
  recommendedTrainingContainer: {
    alignItems: 'center',
    justifyContent: 'center',
    marginTop: 20,
  },
  text: {
    color: 'white',
    fontWeight: 'bold',
    fontSize: 17,
    marginBottom: 10,
  },
  trainingBox: {
    backgroundColor: '#1C1C1E',
    padding: 15,
    borderRadius: 10,
    width: '100%',
    alignItems: 'baseline',
  },
  trainingName: {
    color: 'white',
    fontWeight: 'bold',
    fontSize: 20,
    marginBottom: 10,
  },
  exerciseDetails: {
    alignItems: 'flex-end',
    width: '100%',
    marginBottom: 10,
  },
  trainingDetails: {
    color: 'grey',
    fontSize: 15,
  },
  startButton: {
    backgroundColor: '#A5D32D',
    paddingVertical: 15,
    paddingHorizontal: 20,
    borderRadius: 10,
  },
  buttonText: {
    color: 'white',
    fontWeight: 'bold',
    textAlign: 'center',
  },
  timeContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 10,
    borderRadius: 5,
    padding: 5,
    backgroundColor: 'grey',
  },
  fireIcon: {
    width: 20,
    height: 20,
    marginRight: 5,
  },
  timeText: {
    color: 'white',
    fontSize: 15,
  },
});

export default MiGimnasio;