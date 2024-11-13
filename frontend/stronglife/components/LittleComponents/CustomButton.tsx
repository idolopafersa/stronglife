import React from "react";
import { ActivityIndicator, Text, TouchableOpacity, TouchableOpacityProps, StyleSheet } from "react-native";

interface CustomButtonProps extends TouchableOpacityProps {
  title: string; // Texto a poner en el boton
  handlePress: () => void; // Que hacer cuando se pulsa
  containerStyles?: object; // estilos
  textStyles?: object; // estilos
  isLoading?: boolean; // si esta cargando
}

const CustomButton: React.FC<CustomButtonProps> = ({
  title,
  handlePress,
  containerStyles = {},
  textStyles = {},
  isLoading = false,
  ...props
}) => {
  return (
    <TouchableOpacity
      onPress={handlePress}
      activeOpacity={0.7}
      style={[
        styles.container,
        containerStyles,
        isLoading && styles.loading,
      ]}
      disabled={isLoading}
      {...props}
    >
      <Text style={[styles.text, textStyles]}>
        {title}
      </Text>

      {isLoading && (
        <ActivityIndicator
          animating={isLoading}
          color="#fff"
          size="small"
          style={styles.indicator}
        />
      )}
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  container: {
    
    backgroundColor: "#FFF50A",
    borderRadius: 15,
    padding: 15, 
    minHeight: 62,
    flexDirection: "row",
    justifyContent: "center",
    alignItems: "center",
  },
  loading: {
    opacity: 0.5,
  },
  text: {
    color: "#000000", // Ejemplo de color primario
    fontFamily: "Poppins-SemiBold",
    fontSize: 18,
  },
  indicator: {
    marginLeft: 8,
  },
});

export default CustomButton;
