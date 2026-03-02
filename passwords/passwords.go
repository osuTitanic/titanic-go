package passwords

import (
	"crypto/md5"

	"golang.org/x/crypto/bcrypt"
)

func CreatePasswordHash(password string) (string, error) {
	md5Hash := md5.Sum([]byte(password))
	return CreatePasswordHashFromMd5(string(md5Hash[:]))
}

func VerifyPasswordHash(password string, hash string) bool {
	md5Hash := md5.Sum([]byte(password))
	return VerifyPasswordHashFromMd5(string(md5Hash[:]), hash)
}

func CreatePasswordHashFromMd5(md5 string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(md5),
		bcrypt.DefaultCost,
	)
	return string(hashedBytes), err
}

func VerifyPasswordHashFromMd5(md5 string, hash string) bool {
	if cached, exists := passwordCache[hash]; exists {
		return cached
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(md5))
	if err != nil {
		passwordCache[hash] = false
		return false
	}

	passwordCache[hash] = true
	return true
}
