package resources

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"watersort/model"
)

type Resources struct {
	images map[model.Colour]image.Image
}

func BuildResources(res map[model.Colour]string) (Resources, error) {
	r := Resources{images: map[model.Colour]image.Image{}}

	for key, val := range res {
		img, err := decode(val)
		if err != nil {
			return Resources{}, fmt.Errorf("failed to decode image with key [%s], value [%s]: %w", key, val, err)
		}

		r.images[key] = img
	}

	return r, nil
}

func (r Resources) Get(key model.Colour) (image.Image, error) {
	if img, ok := r.images[key]; ok {
		return img, nil
	}

	return nil, fmt.Errorf("no image for key [%s]", key)
}

func decode(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s: %w", path, err)
	}

	return png.Decode(f)
}
