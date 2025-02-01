import { useUser } from '@/app/context/UserContext'
import { ThemedText } from '@/components/ThemedText'
import { ThemedView } from '@/components/ThemedView'
import { useRouter } from 'expo-router'
import { Button, StyleSheet, Text } from 'react-native'

export default function Conf() {
  const currentUser = useUser()
  const router = useRouter()

  const handleExercise = () => {
    console.log('Navigating to Exercise')
    router.navigate('./ExerciseList')
  }
  const handleRoutine = () => {
    router.navigate('/Routine')
  }
  const handleMeals = () => {
    router.navigate('/Meal')
  }

  return (
    <ThemedView style={styles.container}>
      <Button
        onPress={handleExercise}
        title="Editar Ejercicios"
        color="#841584"
        accessibilityLabel="Editar Ejercicios añadir, modificar o eliminar"
      />
      <Button
        onPress={handleRoutine}
        title="Editar Rutinas"
        color="#841584"
        accessibilityLabel="Editar Rutinas añadir, modificar o eliminar"
      />

      <Button
        onPress={handleMeals}
        title="Editar Comidas"
        color="#841584"
        accessibilityLabel="Editar Comidas añadir, modificar o eliminar"
      />
    </ThemedView>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 16,
  },
})
