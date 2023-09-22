package utils

import (
	"crypto"
	"encoding/hex"
)

func Hashing256(text string) (string, error) {
	// create hash
	hash := crypto.SHA256.New()
	// write text in hash
	_, err := hash.Write([]byte(text))
	if err != nil {
		return "", err
	}
	// calculate hash
	hashed := hash.Sum(nil)
	// convert hash to string and return
	return hex.EncodeToString(hashed), nil
}
