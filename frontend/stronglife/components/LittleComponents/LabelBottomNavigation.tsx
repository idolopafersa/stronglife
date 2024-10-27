import * as React from 'react';
import BottomNavigation from '@mui/material/BottomNavigation';
import BottomNavigationAction from '@mui/material/BottomNavigationAction';
import FolderIcon from '@mui/icons-material/Folder';
import RestoreIcon from '@mui/icons-material/Restore';
import FavoriteIcon from '@mui/icons-material/Favorite';
import LocationOnIcon from '@mui/icons-material/LocationOn';
import { useNavigation } from '@react-navigation/native';

export default function LabelBottomNavigation() {
  const [value, setValue] = React.useState('Home');
  const navigation = useNavigation();

  const handleChange = (event: React.SyntheticEvent, newValue: string) => {
    setValue(newValue);
    if (newValue === 'home') {
      navigation.navigate('Home');
    } else if (newValue === 'details') {
      navigation.navigate('Details');
    } else if (newValue === 'help') {
      navigation.navigate('Help');
    }
  };

  return (
    <BottomNavigation  showLabels
    value={value}
    onChange={(event, newValue) => {
      setValue(newValue);
    }}>
      <BottomNavigationAction
        label="Home"
        value="home"
        icon={<RestoreIcon />}
      />
      <BottomNavigationAction
        label="Details"
        value="details"
        icon={<FavoriteIcon />}
      />
      <BottomNavigationAction
        label="Help"
        value="help"
        icon={<LocationOnIcon />}
      />
      
    </BottomNavigation>
  );
}