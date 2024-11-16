package controllers

import (
	driver "backendgo/Driver"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

// CheckSize limita el tamaño del archivo a 100 KB
func CheckSize(w http.ResponseWriter, r *http.Request) error {
	r.Body = http.MaxBytesReader(w, r.Body, 100*1024)

	err := r.ParseMultipartForm(100 * 1024)
	if err != nil {
		return err
	}
	return nil
}

// CheckImageType verifica que el archivo sea de tipo imagen
func CheckImageType(file multipart.File) error {
	mtype, err := mimetype.DetectReader(file)
	if err != nil {
		return fmt.Errorf("error detecting file type: %w", err)
	}

	if !strings.HasPrefix(mtype.String(), "image/") {
		return fmt.Errorf("file type not supported: %s", mtype.String())
	}

	return nil
}

func UploadMealImage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	err := CheckSize(w, r)
	if err != nil {
		log.Printf("Error Uploading image, max size: %s", err)
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

	if err := CheckImageType(file); err != nil {
		log.Printf("Error uploading Meal image: %s", err)
		http.Error(w, "Invalid image type", http.StatusUnsupportedMediaType)
		return
	}

	filepath := fmt.Sprintf("./assets/meals/meal-%s.jpg", id)

	if err := driver.UploadImage(id, filepath, file); err != nil {
		log.Printf("Error Uploading Meal image driver: %s", err)
		http.Error(w, "couldn't upload image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image uploaded and converted to JPEG successfully"))
}

// GetMealImage obtiene la imagen de la comida
func GetMealImage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "meal ID is required", http.StatusBadRequest)
		return
	}
	path := fmt.Sprintf("./assets/meals/meal-%s.jpg", id) // Asegúrate de que el sufijo sea .jpg

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Error getting Meal Image: %s", err)
		http.Error(w, "image not found", http.StatusNotFound)
		return
	}

	// Servir el archivo como una imagen
	w.Header().Set("Content-Type", "image/jpeg") // Cambiar a "image/jpeg"
	http.ServeFile(w, r, path)
}

// UploadExerciseImage carga la imagen de ejercicio
func UploadExerciseImage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	err := CheckSize(w, r)
	if err != nil {
		log.Printf("Error Uploading image, max size: %s", err)
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

	if err := CheckImageType(file); err != nil {
		log.Printf("Error uploading Exercise image: %s", err)
		http.Error(w, "Invalid image type", http.StatusUnsupportedMediaType)
		return
	}

	filepath := fmt.Sprintf("./assets/exercises/exercise-%s.jpg", id)

	if err := driver.UploadImage(id, filepath, file); err != nil {
		log.Printf("Error Uploading Exercise image driver: %s", err)
		http.Error(w, "couldn't upload image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Exercise image uploaded successfully"))
}

// GetExerciseImage obtiene la imagen de ejercicio
func GetExerciseImage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	path := fmt.Sprintf("./assets/exercises/exercise-%s.jpg", id)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Error getting Exercise Image: %s", err)
		http.Error(w, "image not found", http.StatusNotFound)
		return
	}

	// Servir el archivo como una imagen
	w.Header().Set("Content-Type", "image/jpeg") // Cambiar a "image/jpeg"
	http.ServeFile(w, r, path)
}

// UploadRoutineImage carga la imagen de rutina
func UploadRoutineImage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	err := CheckSize(w, r)
	if err != nil {
		log.Printf("Error Uploading image, max size: %s", err)
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

	if err := CheckImageType(file); err != nil {
		log.Printf("Error uploading Routine image: %s", err)
		http.Error(w, "Invalid image type", http.StatusUnsupportedMediaType)
		return
	}

	filepath := fmt.Sprintf("./assets/routines/routine-%s.jpg", id)

	if err := driver.UploadImage(id, filepath, file); err != nil {
		log.Printf("Error Uploading Routine image driver: %s", err)
		http.Error(w, "couldn't upload image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Routine image uploaded successfully"))
}

// GetRoutineImage obtiene la imagen de rutina
func GetRoutineImage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	path := fmt.Sprintf("./assets/routines/routine-%s.jpg", id)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Error getting Routine Image: %s", err)
		http.Error(w, "image not found", http.StatusNotFound)
		return
	}

	// Servir el archivo como una imagen
	w.Header().Set("Content-Type", "image/jpeg") // Cambiar a "image/jpeg"
	http.ServeFile(w, r, path)
}
