package controller

import (
	"net/http"
)

const imagesDir = "modelsDir"

func Face(w http.ResponseWriter, r *http.Request) {
	// rec, err := face.NewRecognizer(imagesDir)
	// if err != nil {
	// 	log.Fatalf("Can't init face recognizer: %v", err)
	// }
	// // Free the resources when you're finished.
	// defer rec.Close()

	// testImagePristin := filepath.Join(imagesDir, "lalala.jpg")
	// // Recognize faces on that image.
	// faces, err := rec.RecognizeFile(testImagePristin)
	// if err != nil {
	// 	log.Fatalf("Can't init face sasasasa: %v", err)
	// }

	// fmt.Println("ini wajaaahh ===", len(faces))
}
