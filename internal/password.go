package internal

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// 生成哈希，cost 参数控制计算成本（越高越安全，但越慢）
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(password, hash string) bool {
	// 验证密码是否匹配
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
