package hashx

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// NewSalt 生成指定长度的随机盐值
func NewSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// SHA256 使用 SHA256 算法对密码进行加盐哈希
func SHA256(password string, salt string) (string, error) {
	if salt == "" {
		return "", errors.New("empty salt")
	}
	hashes := sha256.New()
	hashes.Write([]byte(password))
	hashes.Write([]byte(salt))
	hash := base64.StdEncoding.EncodeToString(hashes.Sum(nil))
	return hash, nil
}

// SHA256Check 检查给定密码是否与哈希值匹配
func SHA256Check(password string, hashedPassword string, salt string) bool {
	hash, _ := SHA256(password, salt)
	return hash == hashedPassword
}
