import { View } from 'react-native'
import { Button } from '@rneui/themed'
import { Icon } from '@rneui/base/dist/Icon/Icon'
import { Navigation } from '../index'
import { useNavigation } from '@react-navigation/native'

export default function Custom() {
  const navigator = useNavigation()

  function handleExercises() {
    navigator.navigate('ExerciseList')
  }

  function handleRoutines() {
    navigator.navigate('RoutineList')
  }

  function handleMeals() {
    navigator.navigate('MealList')
  }

  return (
    <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
      <Button
        buttonStyle={{ width: 150 }}
        containerStyle={{ margin: 5 }}
        disabledStyle={{
          borderWidth: 2,
          borderColor: '#00F',
        }}
        disabledTitleStyle={{ color: 'green' }}
        icon={<Icon name="fitness-center" size={15} color="#0FF" />}
        loadingProps={{ animating: true }}
        loadingStyle={{}}
        title="Ejercicios"
        titleProps={{}}
        titleStyle={{ marginHorizontal: 5 }}
        onPress={handleExercises}
      />
      <Button
        buttonStyle={{ width: 150 }}
        containerStyle={{ margin: 5 }}
        disabledStyle={{
          borderWidth: 2,
          borderColor: '#00F',
        }}
        disabledTitleStyle={{ color: 'green' }}
        icon={<Icon name="fitness-center" size={15} color="#0FF" />}
        loadingProps={{ animating: true }}
        loadingStyle={{}}
        title="Rutinas"
        titleProps={{}}
        titleStyle={{ marginHorizontal: 5 }}
        onPress={handleRoutines}
      />
      <Button
        buttonStyle={{ width: 150 }}
        containerStyle={{ margin: 5 }}
        disabledStyle={{
          borderWidth: 2,
          borderColor: '#00F',
        }}
        disabledTitleStyle={{ color: 'green' }}
        icon={<Icon name="fitness-center" size={15} color="#0FF" />}
        loadingProps={{ animating: true }}
        onPress={handleMeals}
        title="Comidas"
        titleStyle={{ marginHorizontal: 5 }}
      />
      <Button
        buttonStyle={{ width: 150 }}
        containerStyle={{ margin: 5 }}
        disabledStyle={{
          borderWidth: 2,
          borderColor: '#00F',
        }}
        disabledTitleStyle={{ color: 'green' }}
        icon={<Icon name="fitness-center" size={15} color="#0FF" />}
        loadingProps={{ animating: true }}
        loadingStyle={{}}
        title="Planificación"
        titleProps={{}}
        titleStyle={{ marginHorizontal: 5 }}
      />
    </View>
  )
}
