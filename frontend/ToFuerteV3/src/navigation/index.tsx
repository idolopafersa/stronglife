import { createBottomTabNavigator } from '@react-navigation/bottom-tabs'
import { HeaderButton, Text } from '@react-navigation/elements'
import {
  createStaticNavigation,
  StaticParamList,
} from '@react-navigation/native'
import { createNativeStackNavigator } from '@react-navigation/native-stack'
import { Image } from 'react-native'
import bell from '../assets/bell.png'
import newspaper from '../assets/newspaper.png'
import { Home } from './screens/Home'
import { Profile } from './screens/Profile'
import { Settings } from './screens/Settings'
import { Updates } from './screens/Updates'
import { NotFound } from './screens/NotFound'
import Login from './screens/Login'
import { CustomHeader } from '../components/CustomHeader'
import { useUser } from '../context/UserContext'

function useIsSignedIn() {
  const currentUser = useUser()
  return currentUser.currentUser.isLoggedIn
}

function useIsLoggedOut() {
  const currentUser = useUser()
  return !currentUser.currentUser.isLoggedIn
}

const HomeTabs = createBottomTabNavigator({
  screens: {
    Home: {
      screen: Home,
      options: {
        header: (props) => <CustomHeader />,
        title: 'Gym',
        tabBarIcon: ({ color, size }) => (
          <Image
            source={newspaper}
            tintColor={color}
            style={{
              width: size,
              height: size,
            }}
          />
        ),
      },
    },
    Updates: {
      screen: Updates,
      options: {
        header: (props) => <CustomHeader />,
        tabBarIcon: ({ color, size }) => (
          <Image
            source={bell}
            tintColor={color}
            style={{
              width: size,
              height: size,
            }}
          />
        ),
      },
    },
  },
})

const RootStack = createNativeStackNavigator({
  groups: {
    LoggedIn: {
      if: useIsSignedIn,
      screens: {
        HomeTabs: {
          screen: HomeTabs,
          options: {
            title: 'Home',
            headerShown: false,
          },
        },
        Profile: {
          screen: Profile,
          linking: {
            path: ':user(@[a-zA-Z0-9-_]+)',
            parse: {
              user: (value) => value.replace(/^@/, ''),
            },
            stringify: {
              user: (value) => `@${value}`,
            },
          },
        },
        Settings: {
          screen: Settings,
          options: ({ navigation }) => ({
            presentation: 'modal',
            headerRight: () => (
              <HeaderButton onPress={navigation.goBack}>
                <Text>Close</Text>
              </HeaderButton>
            ),
          }),
        },
      },
    },
    LoggedOut: {
      if: useIsLoggedOut,
      screens: {
        Login: {
          screen: Login,
          options: {
            headerShown: false,
          },
        },
        NotFound: {
          screen: NotFound,
          options: {
            title: '404',
          },
          linking: {
            path: '*',
          },
        },
      },
    },
  },
})

export const Navigation = createStaticNavigation(RootStack)

type RootStackParamList = StaticParamList<typeof RootStack>

declare global {
  namespace ReactNavigation {
    interface RootParamList extends RootStackParamList {}
  }
}
