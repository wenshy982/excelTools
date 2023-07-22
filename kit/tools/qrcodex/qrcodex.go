package qrcodex

import (
	"encoding/base64"
	"os"

	"github.com/skip2/go-qrcode"
)

// Base64CodeUrl 微信支付链接转二维码并base64编码
func Base64CodeUrl(codeUrl string) string {
	data, _ := QRCode(codeUrl)
	encodedUrl := base64.StdEncoding.EncodeToString(data)
	return encodedUrl
}

// QRCode 将内容生成二维码 https://github.com/skip2/go-qrcode
func QRCode(content string) ([]byte, error) {
	return qrcode.Encode(content, qrcode.Medium, 256)
}

// QRCodeToFile 将内容生成二维码并输出到文件
func QRCodeToFile(content, filename string) error {
	png, err := QRCode(content)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, png, os.FileMode(0644))
}
