package utils

import (
	"image"
	"image/jpeg"
	"os"
)

func LoadImage(path string) (image.Image, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(fd)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func WriteImage(path string, img image.Image) error {
	fd, err := os.Create(path)
	if err != nil {
		return err
	}
	jpeg.Encode(fd, img, &jpeg.Options{Quality: 100})
	return nil
}
