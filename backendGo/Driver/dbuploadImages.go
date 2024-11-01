package driver

import (
	"fmt"
	"image"

	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

// convertImage converts an image to JPG format and saves it to the specified path
func convertImage(path string, file multipart.File) error {
	// Create the directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	// Move the file pointer to the beginning, SI NO SE HACE, LA PROXIMA VEZ QUE SE LEA EL ARCHIVO, SE LEERA DESDE EL FINAL
	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("error seeking to beginning of file: %w", err)
	}

	// Decode the image
	img, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("error decoding image: %w", err)
	}
	log.Printf("Image format: %s", format)

	// Create the destination file
	dstFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer dstFile.Close()

	// Encode the image to JPG format
	if err := jpeg.Encode(dstFile, img, nil); err != nil {
		return fmt.Errorf("error encoding image to JPG: %w", err)
	}

	return nil
}

func UploadMealImage(id string, path string, file multipart.File) error {

	if err := convertImage(path, file); err != nil {
		return fmt.Errorf("error converting image: %w", err)
	}

	query := "UPDATE Meals SET photo_url = ? WHERE id = ?"
	if _, err := db.Exec(query, path, id); err != nil {
		log.Printf("error saving meal path in db: %s", err)
		return fmt.Errorf("error updating database: %w", err)
	}

	return nil
}

func UploadExerciseImage(id string, path string, file multipart.File) error {
	if err := convertImage(path, file); err != nil {
		return fmt.Errorf("error converting image: %w", err)
	}

	query := "UPDATE Exercises SET photo_url = ? WHERE id = ?"
	_, erro := db.Exec(query, path, id)
	if erro != nil {
		log.Printf("error saving exercise path in db : %s", erro)
		return erro
	}
	return nil
}

func UploadRoutineImage(id string, path string, file multipart.File) error {

	if err := convertImage(path, file); err != nil {
		return fmt.Errorf("error converting image: %w", err)
	}

	query := "UPDATE Routines SET photo_url = ? WHERE id = ?"
	_, erro := db.Exec(query, path, id)
	if erro != nil {
		log.Printf("error saving Routine path in db : %s", erro)
		return erro
	}
	return nil
}
