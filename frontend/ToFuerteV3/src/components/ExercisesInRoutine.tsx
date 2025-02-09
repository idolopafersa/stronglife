import { useEffect, useState } from 'react'
import { Exercise } from '../interfaces/Exercise'
import ExerciseList from '../navigation/screens/Exercise/ExerciseList'
import { GetRoutineExercises } from '../services/ExercisesRoutines'

const ExercisesInRoutine = ({ route, navigation }: any) => {
  const [exercises, setExercises] = useState<Exercise[]>([])
  const { routine } = route.params

  const fetchExercises = async () => {
    setExercises(await GetRoutineExercises(routine.id))
  }
}
