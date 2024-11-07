import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import HomeScreen from '../HomeScreen';
import HelpScreen from '../HelpScreen';

const Bottomtab = createBottomTabNavigator({
  screens: {
    Home: HomeScreen,
    Help: HelpScreen,
  },
});
export default Bottomtab;