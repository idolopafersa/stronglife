import { useUser } from "@/app/context/UserContext";
import { StyleSheet, View, Text } from "react-native";
import { Image } from "react-native";
import { useState } from "react";

export  const CustomHeader = () => {
    
    const currentUser = useUser(); 


    return (
        <View style={styles.headerContainer}>
          <Text style={styles.greetingText}>¿Listo para entrenar, {currentUser.username}?</Text>
           
            <Image
              source={{ uri: `http://stronglifeapi.fernandezpablo.es/api/user/getimage?user_id=${currentUser.id}` }}
              style={styles.userImage}
            />
          
        </View>
      );
    };
    
    const styles = StyleSheet.create({
      headerContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        justifyContent: 'space-between',
        padding: 20,
        backgroundColor: 'green', // Cambia este color según tus necesidades
      },
      greetingText: {
        color: '#fff',
        fontSize: 18,
      },
      userImage: {
        width: 40,
        height: 40,
        borderRadius: 20,
      },
    });