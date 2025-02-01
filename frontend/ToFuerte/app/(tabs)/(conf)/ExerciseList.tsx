import React, { useState, useEffect } from 'react';
import { View, Text, TextInput, FlatList, TouchableOpacity, Modal, Button, StyleSheet } from 'react-native';
import { getExercises, updateExercise } from '@/services/Exercises/Exercise';
import { Exercise } from '@/services/types/Exercise';

export default function ExerciseList() {
  const [exercises, setExercises] = useState<Exercise[]>([]);
  const [filteredExercises, setFilteredExercises] = useState<Exercise[]>([]);
  const [search, setSearch] = useState('');
  const [selectedExercise, setSelectedExercise] = useState<Exercise | null>(null);
  const [modalVisible, setModalVisible] = useState(false);

  useEffect(() => {
    fetchExercises();
  }, []);

  const fetchExercises = async () => {
    const data = await getExercises();
    setExercises(data);
    setFilteredExercises(data);
  };

  const handleSearch = (text: string) => {
    setSearch(text);
    const filtered = exercises.filter(exercise => exercise.name.toLowerCase().includes(text.toLowerCase()));
    setFilteredExercises(filtered);
  };

  const handleExercisePress = (exercise: Exercise) => {
    setSelectedExercise(exercise);
    setModalVisible(true);
  };

  const handleSave = async () => {
    if (selectedExercise) {
      await updateExercise(selectedExercise);
      fetchExercises();
      setModalVisible(false);
    }
  };

  return (
    <View style={styles.container}>
      <TextInput
        style={styles.searchBar}
        placeholder="Search Exercises"
        value={search}
        onChangeText={handleSearch}
      />
      <FlatList
        data={filteredExercises}
        keyExtractor={(item) => item.id.toString()}
        renderItem={({ item }) => (
          <TouchableOpacity onPress={() => handleExercisePress(item)}>
            <View style={styles.listItem}>
              <Text>{item.name}</Text>
            </View>
          </TouchableOpacity>
        )}
      />
      <Modal
        visible={modalVisible}
        animationType="slide"
        onRequestClose={() => setModalVisible(false)}
      >
        <View style={styles.modalContainer}>
          <Text>Edit Exercise</Text>
          {selectedExercise && (
            <>
              <TextInput
                style={styles.input}
                value={selectedExercise.name}
                onChangeText={(text) => setSelectedExercise({ ...selectedExercise, name: text })}
              />
              <TextInput
                style={styles.input}
                value={selectedExercise.description}
                onChangeText={(text) => setSelectedExercise({ ...selectedExercise, description: text })}
              />
              <Button title="Listo" onPress={handleSave} />
            </>
          )}
        </View>
      </Modal>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 16,
  },
  searchBar: {
    height: 40,
    borderColor: 'gray',
    borderWidth: 1,
    marginBottom: 12,
    paddingHorizontal: 8,
  },
  listItem: {
    padding: 16,
    borderBottomWidth: 1,
    borderBottomColor: 'gray',
  },
  modalContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 16,
  },
  input: {
    height: 40,
    borderColor: 'gray',
    borderWidth: 1,
    marginBottom: 12,
    paddingHorizontal: 8,
    width: '100%',
  },
});