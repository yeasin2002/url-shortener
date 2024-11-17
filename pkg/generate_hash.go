package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GenerateHash(originalUrl string) (string, error) {
	hashMaker := md5.New()
	_, err := hashMaker.Write([]byte(originalUrl))
	mainData := hashMaker.Sum(nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	hash := hex.EncodeToString(mainData)
	return hash[:8], nil
}
