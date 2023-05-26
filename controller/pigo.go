package controller

import (
	"go-markdown/serivce/response"
	"io/ioutil"
	"log"
	"net/http"

	pigo "github.com/esimov/pigo/core"
)

func Pigo(w http.ResponseWriter, r *http.Request) {
	cascadeFile, err := ioutil.ReadFile("./facefinder")
	if err != nil {
		log.Fatalf("Error reading the cascade file: %v", err)
	}

	src, err := pigo.GetImage("./images/photo_ktp.jpg")
	if err != nil {
		log.Fatalf("Cannot open the image file: %v", err)
	}

	pixels := pigo.RgbToGrayscale(src)
	cols, rows := src.Bounds().Max.X, src.Bounds().Max.Y

	cParams := pigo.CascadeParams{
		MinSize:     20,
		MaxSize:     1000,
		ShiftFactor: 0.1,
		ScaleFactor: 1.1,

		ImageParams: pigo.ImageParams{
			Pixels: pixels,
			Rows:   rows,
			Cols:   cols,
			Dim:    cols,
		},
	}

	pigo := pigo.NewPigo()
	// Unpack the binary file. This will return the number of cascade trees,
	// the tree depth, the threshold and the prediction from tree's leaf nodes.
	classifier, err := pigo.Unpack(cascadeFile)
	if err != nil {
		log.Fatalf("Error reading the cascade file: %s", err)
	}

	angle := 0.0 // cascade rotation angle. 0.0 is 0 radians and 1.0 is 2*pi radians

	// Run the classifier over the obtained leaf nodes and return the detection results.
	// The result contains quadruplets representing the row, column, scale and detection score.
	det := classifier.RunCascade(cParams, angle)

	// // Calculate the intersection over union (IoU) of two clusters.
	dets := classifier.ClusterDetections(det, 0.2)

	response.NewResponse().WithCode(http.StatusAccepted).WithData(map[string]interface{}{
		"html": dets,
	}).ParseResponse(w, r)
}
