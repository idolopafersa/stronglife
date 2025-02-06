import { Button, Text } from '@react-navigation/elements'
import { StyleSheet, View } from 'react-native'
import { useUser } from '../../context/UserContext'
import WeeklyCalendar from '../../components/WeeklyCalendar'

export function Home() {
  const { currentUser } = useUser()

  return (
    <View style={styles.container}>
      <WeeklyCalendar />
      <Text>Home Screen</Text>
      <Text>
        Open up 'src/App.tsx' to start working on your app!
        {currentUser.dayselected}
      </Text>

      <Button screen="Settings">Go to Settings</Button>
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
