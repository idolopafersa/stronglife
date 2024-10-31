import React from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';

const HelpScreen = () => {
  return (
  <View style={styles.container}>
    <Text style={styles.title}>Ayuda</Text>
    
      <Text style={styles.subtitle}>¿Cómo se usa esta app?</Text>
      <Text style={styles.question}>¿Qué es un hábito?</Text>
      <Text style={styles.answer}>Un hábito es una acción que se repite de forma regular y automática.</Text>
      <Text style={styles.question}>¿Cómo se crea un hábito?</Text>
      <Text style={styles.answer}>Para crear un hábito, simplemente selecciona la opción "Crear hábito" en la pantalla principal e introduce el nombre del hábito que deseas crear.</Text>
      <Text style={styles.question}>¿Cómo se marca un hábito como completado?</Text>
      <Text style={styles.answer}>Para marcar un hábito como completado, simplemente selecciona el día correspondiente en el calendario.</Text>
      <Text style={styles.question}>¿Cómo se elimina un hábito?</Text>
      <Text style={styles.answer}>Para eliminar un hábito, selecciona la opción "Eliminar hábito" en la pantalla principal y selecciona el hábito que deseas eliminar.</Text>
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