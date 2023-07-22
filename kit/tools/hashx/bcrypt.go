package hashx

import (
	"golang.org/x/crypto/bcrypt"
)

const hashCost = 10 // 生成哈希值的成本

// Bcrypt 使用 bcrypt 算法生成哈希值
func Bcrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// BcryptCheck 检查给定密码是否与哈希值匹配
func BcryptCheck(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil // if No error
}
