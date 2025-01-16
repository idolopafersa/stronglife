import React, { useState } from 'react';
import { View, Text, ImageBackground, StyleSheet, Modal, TouchableOpacity } from 'react-native';
import BotonesSelector from '../components/BotonesSelector';
import { Picker } from '@react-native-picker/picker';

const Training = () => {
  const [kilos, setKilos] = useState(0);
  const [repeticiones, setRepeticiones] = useState(0);
  const [showKilosPicker, setShowKilosPicker] = useState(false);
  const pesos = Array.from({ length: 181 }, (_, i) => 20 + i);

  const closeModal = () => setShowKilosPicker(false);

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
          source={require('@/assets/images/trice2.jpeg')}
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
                <Text style={styles.kilos}>{kilos} kilos</Text>
              </View>
            </View>
          </View>
        </ImageBackground>
      </View>
      <Text>Kilos: {kilos}</Text>
      <Text>Repeticiones: {repeticiones}</Text>
      <BotonesSelector 
        onIncrementarKilos={() => setShowKilosPicker(true)} 
        onIncrementarRepeticiones={() => {}} 
      />

      <Modal
        transparent
        visible={showKilosPicker}
        animationType="slide"
        onRequestClose={closeModal}
      >
        <View style={styles.modalContainer}>
          <View style={styles.modalContent}>
            <Text style={styles.modalTitle}>Selecciona los Kilos</Text>
            <Picker
              selectedValue={kilos}
              onValueChange={(itemValue) => setKilos(itemValue)}
              style={styles.picker}
            >
              {pesos.map((kg) => (
                <Picker.Item key={kg} label={`${kg} kg`} value={kg} />
              ))}
            </Picker>
            <TouchableOpacity onPress={closeModal} style={styles.closeButton}>
              <Text style={styles.closeButtonText}>Cerrar</Text>
            </TouchableOpacity>
          </View>
        </View>
      </Modal>
    </View>
  );
};

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
  modalContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
  },
  modalContent: {
    backgroundColor: '#1C1C1E',
    borderRadius: 10,
    padding: 20,
    width: '80%',
    alignItems: 'center',
  },
  modalTitle: {
    fontSize: 20,
    marginBottom: 20,
    color: 'white',
  },
  picker: {
    width: '100%',
    height: 200,
    color: 'white',
  },
  closeButton: {
    marginTop: 20,
    backgroundColor: '#FF6347',
    padding: 10,
    borderRadius: 5,
  },
  closeButtonText: {
    color: 'white',
    fontSize: 16,
  },
});

export default Training;