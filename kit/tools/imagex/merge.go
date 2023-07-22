package imagex

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

// MergeImages 四个较小的图像重新组合为一个较大的图像
func MergeImages(imgPath1, imgPath2, imgPath3, imgPath4, outputDir string) error {
	// 从文件中读取图像
	img1, err := loadImage(imgPath1)
	if err != nil {
		return err
	}

	img2, err := loadImage(imgPath2)
	if err != nil {
		return err
	}

	img3, err := loadImage(imgPath3)
	if err != nil {
		return err
	}

	img4, err := loadImage(imgPath4)
	if err != nil {
		return err
	}

	// 创建新图像
	bounds := img1.Bounds()
	width := bounds.Max.X * 2
	height := bounds.Max.Y * 2
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// 将四个图像画到新图像上
	draw.Draw(newImg, img1.Bounds(), img1, bounds.Min, draw.Src)
	draw.Draw(newImg, image.Rect(bounds.Max.X, bounds.Min.Y, width, bounds.Max.Y), img2, bounds.Min, draw.Src)
	draw.Draw(newImg, image.Rect(bounds.Min.X, bounds.Max.Y, bounds.Max.X, height), img3, bounds.Min, draw.Src)
	draw.Draw(newImg, image.Rect(bounds.Max.X, bounds.Max.Y, width, height), img4, bounds.Min, draw.Src)

	// 将新图像写入文件
	outputFile, err := os.Create(fmt.Sprintf("%s/merge.png", outputDir))
	if err != nil {
		return err
	}

	err = png.Encode(outputFile, newImg)
	if err != nil {
		return err
	}

	return outputFile.Close()
}

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) { _ = file.Close() }(file)

	img, _, err := image.Decode(file)
	return img, err
}
