import React, { useState } from 'react';
import { View, Text, StyleSheet, Image } from 'react-native';

const Header = () => {
  const [showContent, setShowContent] = useState(true);

  return (
    <View style={styles.headerContainer}>
      {showContent ? (
        <>
          <View style={styles.textContainer}>
            <Text style={styles.headerName}>Hello Paco</Text>
            <Text style={styles.subTitle}>Let's start your day !</Text>
          </View>
          <Image
            source={require('@/assets/images/churum.png')}
            style={styles.profileImage}
          />
        </>
      ) : (
        <Text>Hola Mundo</Text>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  headerContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    backgroundColor: '#1C1C1E',
    padding: 15,
    paddingLeft: 20,
  },
  textContainer: {
    flexDirection: 'column',
  },
  headerName: {
    color: '#A5D32D',
    fontSize: 15,
  },
  subTitle: {
    color: 'white',
    fontSize: 20,
    fontWeight: 'bold',
  },
  profileImage: {
    width: 50,
    height: 50,
    borderRadius: 20,
  },
});

export default Header;