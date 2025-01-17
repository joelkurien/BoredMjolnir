package main

type Ball struct {
	pivotX   float64
	pivotY   float64
	velocity float64
}

func Tesseract() *Ball {
	return &Ball{
		pivotX:   0,
		pivotY:   0,
		velocity: 0,
	}
}
