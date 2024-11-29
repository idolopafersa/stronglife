import React from 'react';
import { View, Text, StyleSheet } from 'react-native';
import Icon from 'react-native-vector-icons/FontAwesome';
import { Image } from 'react-native';

const Header = () => {
  return (
    <View style={styles.headerContainer}>
      <Text style={styles.headerTitle}>Hello</Text>
      <Text style={styles.headerTitle}>Paco</Text>

      <Image source={require('@/assets/images/churum.png')} style={{ width: 50, height: 50 }} />
    </View>
  );
};

const styles = StyleSheet.create({
  headerContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    backgroundColor: '#000',
    padding: 15,
  },
  headerTitle: {
    color: '#A5D32D',
    fontSize: 20,
    fontWeight: 'bold',
    alignItems: 'stretch',
  },
  icon: {
    marginRight: 10,
  },
});

export default Header;