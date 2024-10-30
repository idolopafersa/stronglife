package controllers

import (
	driver "backendgo/Driver"
	"fmt"
	"log"
	"net/http"
	"os"
)

func UploadMealImage(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	/*_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}
	*/
	erro := r.ParseMultipartForm(10 << 20) // Limit Photo to 10mb
	if erro != nil {
		http.Error(w, "File too big", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		log.Printf("Error uploading Meal image: %s", err)
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	filepath := fmt.Sprintf("./assets/meals/meal-%s", id)

	erre := driver.UploadMealImage(id, filepath, file)
	if erre != nil {
		log.Printf("Error Uploading Meal image driver: %s /n", erre)
		http.Error(w, "couldnt upload image", http.StatusNotFound)
		return
	}

}

func GetMealImage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "meal ID is required", http.StatusBadRequest)
		return
	}
	path := fmt.Sprintf("./assets/meals/meal-%s", id)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Error getting Meal Image: %s", err)
		http.Error(w, "image not found", http.StatusNotFound)
		return
	}

	// Serve the file as an image
	w.Header().Set("Content-Type", "image/png") // or "image/png" based on your files
	http.ServeFile(w, r, path)
}

func UploadExerciseImage(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	/*_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}
	*/
	erro := r.ParseMultipartForm(10 << 20) // Limit Photo to 10mb
	if erro != nil {
		http.Error(w, "File too big", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		log.Printf("Error uploading Exercise image: %s", err)
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	filepath := fmt.Sprintf("./assets/exercises/exercise-%s", id)

	erre := driver.UploadExerciseImage(id, filepath, file)
	if erre != nil {
		log.Printf("Error Uploading Exercise image driver: %s /n", erre)
		http.Error(w, "couldnt upload image", http.StatusNotFound)
		return
	}

}

func GetExerciseImage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, " ID is required", http.StatusBadRequest)
		return
	}
	path := fmt.Sprintf("./assets/exercises/exercise-%s", id)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Error getting Exercise Image: %s", err)
		http.Error(w, "image not found", http.StatusNotFound)
		return
	}

	// Serve the file as an image
	w.Header().Set("Content-Type", "image/png") // or "image/png" based on your files
	http.ServeFile(w, r, path)
}

func UploadRoutineImage(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	/*_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}
	*/
	erro := r.ParseMultipartForm(10 << 20) // Limit Photo to 10mb
	if erro != nil {
		http.Error(w, "File too big", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		log.Printf("Error uploading Routine image: %s", err)
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	filepath := fmt.Sprintf("./assets/routines/routine-%s", id)

	erre := driver.UploadRoutineImage(id, filepath, file)
	if erre != nil {
		log.Printf("Error Uploading Routine image driver: %s /n", erre)
		http.Error(w, "couldnt upload image", http.StatusNotFound)
		return
	}

}

func GetRoutineImage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, " ID is required", http.StatusBadRequest)
		return
	}
	path := fmt.Sprintf("./assets/routines/routine-%s", id)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Error getting Routine Image: %s", err)
		http.Error(w, "image not found", http.StatusNotFound)
		return
	}

	// Serve the file as an image
	w.Header().Set("Content-Type", "image/png") // or "image/png" based on your files
	http.ServeFile(w, r, path)
}
