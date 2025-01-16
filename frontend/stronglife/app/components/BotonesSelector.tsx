import React from 'react';
import { View, Button, StyleSheet, Dimensions } from 'react-native';

interface BotonesSelectorProps {
  onIncrementarKilos: () => void;
  onIncrementarRepeticiones: () => void;
}

const BotonesSelector: React.FC<BotonesSelectorProps> = ({ onIncrementarKilos, onIncrementarRepeticiones }) => {
  return (
    <View style={styles.footerContainer}>
      <View style={styles.buttonContainer}>
        <Button title=" Kilos" color="#FF0000" onPress={onIncrementarKilos} />
        <Button title="Repeticiones" color="#FF0000" onPress={onIncrementarRepeticiones} />
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  footerContainer: {
    position: 'absolute',
    bottom: 0,
    width: '100%',
    height: Dimensions.get('window').height / 5,
    backgroundColor: '#808080',
    justifyContent: 'center',
    alignItems: 'center',
    borderTopStartRadius: 75,
    borderTopEndRadius: 75,
  },
  buttonContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    width: '80%',
  },
});

export default BotonesSelector;