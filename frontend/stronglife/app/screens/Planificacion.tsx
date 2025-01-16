import React from 'react';
import { View, Text, StyleSheet } from 'react-native';
import Formulario from '../components/Formulario';

const Planificacion = () => {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Planificación</Text>
      <Formulario />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'black',
  },
  title: {
    fontSize: 24,
    color: 'white',
    marginBottom: 20,
  },
});

export default Planificacion;