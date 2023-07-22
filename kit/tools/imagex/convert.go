package imagex

import (
	"image"
	"image/png"
	"os"

	_ "golang.org/x/image/webp"
)

// ConvertWebpToPng 接受一个webp格式图片文件的路径，并将其转化为png格式
func ConvertWebpToPng(webpPath string, pngPath string) error {
	// 打开webp文件
	webpFile, err := os.Open(webpPath)
	if err != nil {
		return err
	}
	defer func(webpFile *os.File) { _ = webpFile.Close() }(webpFile)

	// 使用x/image/webp解码
	img, _, err := image.Decode(webpFile)
	if err != nil {
		return err
	}

	// 创建输出的PNG文件
	outFile, err := os.Create(pngPath)
	if err != nil {
		return err
	}
	defer func(outFile *os.File) { _ = outFile.Close() }(outFile)

	// 将img编码为PNG格式并输出
	err = png.Encode(outFile, img)
	if err != nil {
		return err
	}

	return nil
}
