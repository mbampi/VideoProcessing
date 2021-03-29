package effect

import (
	"image"

	"gocv.io/x/gocv"
)

type EffectType int

const (
	Original EffectType = iota
	GaussianBlur
	Canny
	Sobel
	RotateClockwise
	RotateCounterClockwise
	Rotate180
	Grayscale
	VerticalFlip
	HorizontalFlip
	Brightness
	Contrast
	Resize
)

var EffectToString = map[EffectType]string{
	GaussianBlur:           "Gaussian Blur",
	Canny:                  "Canny",
	Sobel:                  "Sobel",
	Grayscale:              "Grayscale",
	Brightness:             "Brightness",
	Contrast:               "Contrast",
	Resize:                 "Resize",
	Rotate180:              "Rotate 180",
	RotateClockwise:        "Rotate Clockwise",
	RotateCounterClockwise: "Rotate CounterClockwise",
	VerticalFlip:           "Vertical Flip",
	HorizontalFlip:         "Horizontal Flip",
	Original:               "Original",
}

func (e EffectType) String() string {
	effectName, ok := EffectToString[e]
	if !ok {
		effectName = "Original"
	}
	return effectName
}

func ApplyEffect(img gocv.Mat, res *gocv.Mat, effect EffectType, level int) {
	switch effect {

	case GaussianBlur:
		kSize := level*2 + 1
		kernelSize := image.Pt(kSize, kSize)
		gocv.GaussianBlur(img, res, kernelSize, 10, 10, gocv.BorderDefault)
		break
	case Canny:
		gocv.Canny(img, res, float32(level)*3, 30)
		gocv.Merge([]gocv.Mat{*res, *res, *res}, res)
		break
	case Sobel:
		gocv.Sobel(img, res, gocv.MatTypeCV32F, 1, 1, 5, 1, 0, gocv.BorderDefault)
		res.ConvertTo(res, gocv.MatTypeCV8UC3)
		break
	case RotateClockwise:
		gocv.Rotate(img, res, gocv.Rotate90Clockwise)
		break
	case RotateCounterClockwise:
		gocv.Rotate(img, res, gocv.Rotate90CounterClockwise)
		break
	case Rotate180:
		gocv.Rotate(img, res, gocv.Rotate180Clockwise)
		break
	case Grayscale:
		gocv.CvtColor(img, res, gocv.ColorRGBToGray)
		gocv.Merge([]gocv.Mat{*res, *res, *res}, res)
		break
	case HorizontalFlip:
		gocv.Flip(img, res, 1)
		break
	case VerticalFlip:
		gocv.Flip(img, res, 0)
		break
	case Brightness:
		img.ConvertToWithParams(res, -1, 1, float32(level-50))
		break
	case Contrast:
		img.ConvertToWithParams(res, -1, float32(level)/10, 0)
		break
	case Resize:
		width := img.Cols() / (1 + level/4)
		height := img.Rows() / (1 + level/4)
		size := image.Point{X: width, Y: height}
		gocv.Resize(img, res, size, 0, 0, gocv.InterpolationDefault)
		break

	default:
		*res = img.Clone()
	}
}

func (e EffectType) AllowsRecording() bool {
	switch e {
	case
		RotateClockwise,
		RotateCounterClockwise,
		Resize:
		return false
	}
	return true
}

var KeyToMap = map[int]EffectType{
	49:  GaussianBlur,           // 1
	50:  Canny,                  // 2
	51:  Sobel,                  // 3
	52:  Grayscale,              // 4
	53:  Brightness,             // 5
	54:  Contrast,               // 6
	55:  Resize,                 // 7
	0:   Rotate180,              // UP
	2:   RotateClockwise,        // RIGHT
	3:   RotateCounterClockwise, // LEFT
	104: VerticalFlip,           // V
	118: HorizontalFlip,         // H
}

func EffectFromKey(key int) EffectType {
	effect, ok := KeyToMap[key]
	if !ok {
		effect = Original
	}
	return effect
}
