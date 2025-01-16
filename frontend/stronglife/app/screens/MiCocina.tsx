import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

const MiCocina = () => {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Pantalla de Mi Cocina</Text>
      <Text style={styles.description}>Aquí puedes gestionar tus recetas y comidas.</Text>
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
    fontWeight: 'bold',
  },
  description: {
    fontSize: 16,
    color: 'grey',
    textAlign: 'center',
    margin: 10,
  },
});

export default MiCocina;