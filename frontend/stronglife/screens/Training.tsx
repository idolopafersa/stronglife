import React from 'react';
import { View, Text, ImageBackground, StyleSheet } from 'react-native';

function Training() {
  return (
    <View style={styles.container}>
      <View>
        <Text style={styles.header}>
          5 <Text style={styles.subTitle}>Ejercicios {'\n'} Restantes</Text>
        </Text>
      </View>
      <View style={styles.subcontainer}>
        <Text style={styles.subTitle}>
          Siguiente: <Text style={{ color: 'grey' }}>Aerobics-Sets (1/1)</Text>
        </Text>
      </View>
      <View style={styles.imageBackground}>
        <ImageBackground
          source={require('@/assets/images/trice.png')}
          resizeMode="cover"
          style={styles.imageBackground}
        >
          <View style={styles.ejercicioscont}>
            <View style={styles.row}>
              <Text style={styles.ejercicios}>
                Triceps {'\n'} Extensions
              </Text>
              <View style={styles.repeticionesContainer}>
                <Text style={styles.repeticiones}>3 Repeticiones</Text>
                <Text style={styles.kilos}>50 kilos</Text>
              </View>
            </View>
          </View>
        </ImageBackground>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    backgroundColor: 'black',
    flex: 1,
  },
  subcontainer: {
    backgroundColor: '#3A3A3C',
    padding: 10,
  },
  header: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    backgroundColor: '#1C1C1E',
    padding: 15,
    paddingLeft: 20,
    color: 'white',
    fontSize: 30,
    fontWeight: 'bold',
  },
  subTitle: {
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'space-between',
    padding: 15,
    paddingLeft: 20,
    fontSize: 20,
    color: 'white',
    fontWeight: 'bold',
  },
  
  imageBackground: {
    flex: 1,
    width: '100%',
    height: '100%',
    justifyContent: 'center',
    alignItems: 'center',
    
        
  },
  ejercicioscont: {
    
    padding: 10,
    alignItems: 'center',
    position: 'relative',
  },
  row: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    width: '100%',
  },
  ejercicios: {
    fontSize: 30,
    color: 'red',
    fontWeight: 'bold',
  },
  repeticionesContainer: {
    alignItems: 'flex-end',
  },
  repeticiones: {
    position: 'relative',
    top: -10, 
    color: 'red',
    fontSize: 20,
    fontWeight: 'bold',
  },
  kilos: {
    color: 'red',
    fontSize: 20,
    fontWeight: 'bold',
    marginTop: 55,
  },
});

export default Training;