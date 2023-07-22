package imagex

import (
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"os"

	_ "golang.org/x/image/webp"
)

func CutImage(inputPath, outputDir string) error {
	imgFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer func(imgFile *os.File) { _ = imgFile.Close() }(imgFile)

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}

	rect := img.Bounds()
	width := rect.Dx() / 2
	height := rect.Dy() / 2

	// 对图片进行 4 等分切割
	subImg := []*image.RGBA{
		image.NewRGBA(image.Rect(0, 0, width, height)),
		image.NewRGBA(image.Rect(0, 0, width, height)),
		image.NewRGBA(image.Rect(0, 0, width, height)),
		image.NewRGBA(image.Rect(0, 0, width, height)),
	}

	quarters := []struct {
		dx, dy int
	}{
		{0, 0},
		{width, 0},
		{0, height},
		{width, height},
	}

	for i, quarter := range quarters {
		bounds := subImg[i].Bounds()
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				subImg[i].Set(x-bounds.Min.X, y-bounds.Min.Y, img.At(x+quarter.dx, y+quarter.dy))
			}
		}

		// 输出切割后的图片
		outFile, err := os.Create(fmt.Sprintf("%s/%d.png", outputDir, i+1))
		if err != nil {
			return err
		}
		err = png.Encode(outFile, subImg[i])
		if err != nil {
			_ = outFile.Close()
			return err
		}
		_ = outFile.Close()
	}

	return nil
}
