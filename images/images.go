package images

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/Kamaalio/kamaalgo/files"
	"golang.org/x/image/draw"
)

func ResizeImage(inputImage image.Image, width int, height int, outputImage *os.File) *os.File {
	outputImageSpec := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.NearestNeighbor.Scale(outputImageSpec, outputImageSpec.Rect, inputImage, inputImage.Bounds(), draw.Over, nil)
	png.Encode(outputImage, outputImageSpec)
	return outputImage
}

func OpenAndDecodeImage(filePath string) (image.Image, error) {
	fileContent, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fileContent.Close()

	fileExtension := files.ExtractFileExtension(filePath)
	switch fileExtension {
	case "jpeg", "jpg":
		return jpeg.Decode(fileContent)
	case "png":
		return png.Decode(fileContent)
	default:
		return nil, fmt.Errorf("%s file extension are not supported", fileExtension)
	}
}
