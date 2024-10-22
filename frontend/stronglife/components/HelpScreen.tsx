import React from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';

const HelpScreen = () => {
  return (
    <ScrollView contentContainerStyle={styles.container}>
      <Text style={styles.title}>Pantalla de Ayuda</Text>
      <Text style={styles.subtitle}>¿Cómo puedo ayudarte?</Text>
      <Text style={styles.question}>1. ¿Cómo uso la aplicación?</Text>
      <Text style={styles.answer}>
        La aplicación te permite gestionar tus entrenamientos y rutinas.
      </Text>
      <Text style={styles.question}>2. ¿Cómo puedo crear una cuenta?</Text>
      <Text style={styles.answer}>
        Puedes crear una cuenta haciendo clic en el botón de registro en la pantalla de inicio.
      </Text>
      <Text style={styles.question}>3. ¿Dónde puedo encontrar más información?</Text>
      <Text style={styles.answer}>
        Puedes visitar nuestra página web o contactar con nuestro soporte.
      </Text>
      {/* Agrega más preguntas y respuestas según sea necesario */}
    </ScrollView>
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