import { useUser } from "@/app/context/UserContext";
import { ThemedText } from "@/components/ThemedText";
import { ThemedView } from "@/components/ThemedView";
import { useRouter } from "expo-router";
import { StyleSheet , Text} from "react-native";



export default function Params() {

    const currentUser = useUser();
     const router = useRouter();
   
     
   
     
     return (
       <ThemedView style={styles.container}>
         <ThemedText type="title">Welcome to the Home Screen! {currentUser.username}</ThemedText>
         <Text>
           This is where you can find various components and features of the app .
         </Text>
         {/* Additional components and content can be added here */}
       </ThemedView>
     );
   }
   
   const styles = StyleSheet.create({
     container: {
       flex: 1,
       justifyContent: 'center',
       alignItems: 'center',
       padding: 16,
     },
   });