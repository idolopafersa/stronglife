import { Text } from '@react-navigation/elements'
import { StyleSheet, View } from 'react-native'
import WeeklyCalendar from '../../components/WeeklyCalendar'

export function Kitchen() {
  return (
    <View style={styles.container}>
      <WeeklyCalendar />
      <Text>Updates Screen</Text>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    gap: 10,
  },
})
