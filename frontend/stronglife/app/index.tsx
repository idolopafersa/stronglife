import CustomButton from '@/components/LittleComponents/CustomButton';
import { router } from 'expo-router';
import { Image, ImageBackground, SafeAreaView, StyleSheet, Text,  } from 'react-native';

export default function App() {
  return (
    <ImageBackground 
      source={require('../assets/images/Inicio_Sesion.png')}
      style={styles.container}
      resizeMode="cover"
    >
      <Text style={styles.text}> Welcome  to </Text>
      <Text style={styles.text}> STRONGLIFE</Text>
      <Text style={styles.text}> Archive you body goals with us</Text>
      <CustomButton
      title = "Sign in"
      handlePress={() => router.navigate("/sign_in")}></CustomButton>

    </ImageBackground>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'space-evenly',
  },
  text: {
    color: 'white',
    fontSize: 24,
    fontWeight: 'bold',
  },
});
