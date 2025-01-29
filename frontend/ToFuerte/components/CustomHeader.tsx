import { useUser } from "@/app/context/UserContext";
import { StyleSheet, View, Text, Button, TouchableOpacity } from "react-native";
import { Image } from "react-native";
import { useState } from "react";
import { useRouter } from "expo-router";

export  const CustomHeader = () => {
    
    const currentUser = useUser(); 
    const router = useRouter();

    return (
        <View style={styles.headerContainer}>
          <Text style={styles.greetingText}>¿Listo para entrenar, {currentUser.username}?</Text>
           
            <TouchableOpacity onPress={() => { router.push("/User") }}>
              <Image
                source={{ uri: `http://stronglifeapi.fernandezpablo.es/api/user/getimage?user_id=${currentUser.id}` }}
                style={styles.userImage}
              />
            </TouchableOpacity>
          
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