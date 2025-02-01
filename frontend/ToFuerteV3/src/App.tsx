import { Assets as NavigationAssets } from '@react-navigation/elements'
import { Asset } from 'expo-asset'
import * as SplashScreen from 'expo-splash-screen'
import * as React from 'react'
import { Navigation } from './navigation'
import { UserContext, UserContextProps } from './context/UserContext'
import moment from 'moment'
import { useState } from 'react'
import { checkcookie } from './services/UserServices'
import { NavigationContainer } from '@react-navigation/native'

Asset.loadAsync([
  ...NavigationAssets,
  require('./assets/newspaper.png'),
  require('./assets/bell.png'),
])

SplashScreen.preventAutoHideAsync()

export function App() {
  const [currentUser, setCurrentUser] = useState<UserContextProps>({
    username: '',
    id: 0,
    email: '',
    dayselected: moment().format('YYYY-MM-DD'),
    isLoggedIn: false,
  })

  React.useEffect(() => {
    const fetchUser = async () => {
      const response = await checkcookie()
      if (response) {
        setCurrentUser({
          ...currentUser,
          username: response.data.username,
          id: response.data.id,
          isLoggedIn: true,
        })
      }
    }

    fetchUser()
  }, [])

  return (
    <UserContext.Provider value={{ currentUser, setCurrentUser }}>
      <Navigation
        onReady={() => {
          SplashScreen.hideAsync()
        }}
      />
    </UserContext.Provider>
  )
}
