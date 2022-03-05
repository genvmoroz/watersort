package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strings"
	"watersort/drawable"
	"watersort/model"
	"watersort/resources"
)

func main() {
	args := os.Args

	out, err := readArg("out", args)
	if err != nil {
		log.Panicf("failed to read args: %s", err.Error())
	}

	img := image.NewRGBA(image.Rect(0, 0, 400, 400))
	black := color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xff}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: black}, image.Point{}, draw.Src)

	resMap := map[model.Colour]string{
		model.Red:    "./asserts/red.png",
		model.Pink:   "./asserts/pink.png",
		model.Purple: "./asserts/purple.png",
	}
	res, err := resources.BuildResources(resMap)

	mFlask := model.NewFlask()
	mFlask.Put(model.Red)
	mFlask.Put(model.Pink)
	mFlask.Put(model.Purple)
	dFlask := drawable.NewFlask(&mFlask, res)
	dFlask.Draw(img, image.Rect(200, 50, 251, 351))

	file, err := os.Create(out)
	if err != nil {
		log.Panicf("failed to create file: %s", err.Error())
	}

	if err = png.Encode(file, img); err != nil {
		log.Panicf("failed to write png to file: %s", err.Error())
	}

	log.Println("\nDone. Exit")
}

func readArg(key string, args []string) (string, error) {
	for _, arg := range args {
		parts := strings.Split(strings.TrimSpace(arg), "=")
		if len(parts) == 2 && parts[0] == key {
			return parts[1], nil
		}
	}

	return "", fmt.Errorf("no arg value for key: %s", key)
}
