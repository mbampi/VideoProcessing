package main

import (
	"image"
	"image/color"
	"log"
	. "videoprocessing/effect"

	"gocv.io/x/gocv"
)

func main() {
	deviceID := 0
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		log.Println("Error opening camera", err.Error())
		return
	}
	defer webcam.Close()

	windowOriginal := gocv.NewWindow("Original")
	defer windowOriginal.Close()

	window := gocv.NewWindow("Result")
	defer window.Close()

	trackbar := window.CreateTrackbar("Level", 100)
	trackbar.SetMin(0)
	trackbar.SetPos(50)

	img := gocv.NewMat()
	defer img.Close()

	resultImg := gocv.NewMat()
	defer resultImg.Close()

	recording := false
	writer, err := gocv.VideoWriterFile("video.mov", "jpeg", 24, 1280, 720, true)
	if err != nil {
		log.Println("Error opening video file to write", err.Error())
		return
	}
	defer writer.Close()

	var effect EffectType = Original
	log.Printf("Reading camera device: %v\n", deviceID)

	for {
		// Grab image from webcam
		if ok := webcam.Read(&img); !ok {
			log.Printf("Cannot read device %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// Apply effect using trackbar level
		level := trackbar.GetPos()
		ApplyEffect(img, &resultImg, effect, level)
		window.SetWindowTitle(effect.String())

		// Save frame to video
		if recording && effect.AllowsRecording() {
			err := writer.Write(resultImg)
			if err != nil {
				log.Println("Error recording", err.Error())
			}

			gocv.Circle(&resultImg, image.Pt(15, 15), 5, color.RGBA{255, 0, 0, 255}, -1)
		}

		// Show live video
		windowOriginal.IMShow(img)
		window.IMShow(resultImg)

		// Check keyboard input
		key := window.WaitKey(1)
		if key == 27 {
			log.Println("Stop")
			return
		} else if key == 114 { // R
			recording = !recording
			log.Println("Recording: ", recording)
		} else if key != -1 {
			effect = EffectFromKey(key)
		}
	}
}
