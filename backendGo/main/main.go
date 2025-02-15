package main

import (
	controllers "backendgo/Controllers"
	driver "backendgo/Driver"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lpernett/godotenv"
)

func main() {

	//Intenta pillar las variables del .env
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	//Inicicializar la base de datos
	driver.OpenDB()

	//Creamos el router para empezar las peticiones
	r := mux.NewRouter()

	//Para hacer login y register
	r.HandleFunc("/api/login", controllers.Login).Methods("POST")
	r.HandleFunc("/api/logout", controllers.Logout).Methods("GET")
	r.HandleFunc("/api/user/register", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/user/uploadimage", controllers.UploadUserImage).Methods("POST")
	r.HandleFunc("/api/user/getimage", controllers.GetUserImage).Methods("GET")
	r.HandleFunc("/api/cookie", controllers.CheckCookie).Methods("GET")

	//Para las comidas, deberemos hacer un get post  put y remove
	r.HandleFunc("/api/meal/getall", controllers.AllMeals).Methods("GET")
	r.HandleFunc("/api/meal/get", controllers.GetMeal).Methods("GET")
	r.HandleFunc("/api/meal/post", controllers.PostMeal).Methods("POST")
	r.HandleFunc("/api/meal/del", controllers.DelMeal).Methods("DELETE")
	r.HandleFunc("/api/meal/put", controllers.PutMeal).Methods("PUT")
	r.HandleFunc("/api/meal/uploadimage", controllers.UploadMealImage).Methods("POST")
	r.HandleFunc("/api/meal/getimage", controllers.GetMealImage).Methods("GET")

	//Para las rutinas, get post put
	r.HandleFunc("/api/routine/get", controllers.GetRoutine).Methods("GET")
	r.HandleFunc("/api/routine/post", controllers.PostRoutine).Methods("POST")
	r.HandleFunc("/api/routine/del", controllers.DelRoutine).Methods("DELETE")
	r.HandleFunc("/api/routine/put", controllers.PutRoutine).Methods("PUT")
	r.HandleFunc("/api/routine/getall", controllers.AllRoutines).Methods("GET")
	r.HandleFunc("/api/routine/uploadimage", controllers.UploadRoutineImage).Methods("POST")
	r.HandleFunc("/api/routine/getimage", controllers.GetRoutineImage).Methods("GET")

	//Para los ejercicios, get post put
	r.HandleFunc("/api/exercise/get", controllers.GetExercise).Methods("GET")
	r.HandleFunc("/api/exercise/post", controllers.PostExercise).Methods("POST")
	r.HandleFunc("/api/exercise/del", controllers.DelExercise).Methods("DELETE")
	r.HandleFunc("/api/exercise/put", controllers.PutExercise).Methods("PUT")
	r.HandleFunc("/api/exercise/getall", controllers.GetAlExercises).Methods("GET")
	r.HandleFunc("/api/exercise/uploadimage", controllers.UploadExerciseImage).Methods("POST")
	r.HandleFunc("/api/routine/getimage", controllers.GetExerciseImage).Methods("GET")

	//add,remove and get exercises of a routine
	r.HandleFunc("/api/exercises/routines/get", controllers.GetrExercises).Methods("GET")
	r.HandleFunc("/api/exercises/routines/post", controllers.AddrExercise).Methods("POST")
	r.HandleFunc("/api/exercises/routines/del", controllers.DelrExercise).Methods("DELETE")

	//sets handling
	r.HandleFunc("/api/sets/get", controllers.GetSet).Methods("GET")
	r.HandleFunc("/api/sets/getall", controllers.GetAlSet).Methods("GET")
	r.HandleFunc("/api/sets/getallroutine", controllers.GetAlSetRoutine).Methods("GET")
	r.HandleFunc("/api/sets/post", controllers.PostSet).Methods("POST")
	r.HandleFunc("/api/sets/put", controllers.PutSet).Methods("PUT")
	r.HandleFunc("/api/sets/del", controllers.DelSet).Methods("DELETE")

	//days handling
	r.HandleFunc("/api/day/get", controllers.Getday).Methods("GET")
	r.HandleFunc("/api/day/routine/post", controllers.AddRoutineToDay).Methods("POST")
	r.HandleFunc("/api/day/meal/post", controllers.AddMealToDay).Methods("POST")
	r.HandleFunc("/api/day/meal/del", controllers.DelMealToDay).Methods("DELETE")
	log.Println(http.ListenAndServe("127.0.0.1:7777", cors(r)))

}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET,  DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// Handle preflight requests
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
