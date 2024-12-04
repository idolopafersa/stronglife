import React from "react";
import { View, Text, TouchableOpacity, SafeAreaView } from "react-native";

export default props => {
  const {
    state,
    descriptors,
    navigation,
    activeBackgroundColor,
    activeTintColor,
    inactiveBackgroundColor,
    inactiveTintColor
    ,
  } = props;
  
  const { routes } = state;
  
  return (
    <SafeAreaView>
      <View
        style={{
          flexDirection: "row",
          height: 64,
          alignItems: "center",
          justifyContent: "space-between",
          backgroundColor: "#1C1C1E"
        }}
      >
        {routes.map((route, index) => {
          const { options } = descriptors[route.key];

          const label =
            options.tabBarLabel !== undefined
              ? options.tabBarLabel
              : options.title !== undefined
              ? options.title
              : route.name;

          const isFocused = state.index === index;

          const tintColor = isFocused ? activeTintColor : inactiveTintColor;
          const backgroundColor = isFocused
            ? activeBackgroundColor
            : inactiveBackgroundColor;

          const onPress = () => {
            navigation.navigate(route.name);
            
          };

          return (
            <TouchableOpacity
              key={index}
              style={{
                backgroundColor: '#1C1C1E',
                flexDirection: "row",
                margin: 4,
                height: 50,
                padding: 8,
                paddingRight: 16,
                paddingLeft: 16,
                justifyContent: "center",
                alignItems: "center"
              }}
              onPress={onPress}
            >
              {options.tabBarIcon !== undefined &&
                options.tabBarIcon({ color: tintColor, size: 28 })}
              {isFocused && (
                <Text
                  style={{
                    marginLeft: 8,
                    color: 'white',
                    fontWeight: "bold",
                  }}
                >
                  {label}
                </Text>
              )}
            </TouchableOpacity>
          );
        })}
      </View>
    </SafeAreaView>
  );
};