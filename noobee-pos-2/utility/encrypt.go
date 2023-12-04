package utility

import "golang.org/x/crypto/bcrypt"

func Hash(plain string) (hash string, err error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	hash = string(hashByte)
	return
}

func Verify(hash string, plain string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return
}
