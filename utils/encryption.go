package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func M_base64(s string) string{

	md := md5.New()
	md.Write([]byte(s))
	mdSum := md.Sum(nil)

	hexString := hex.EncodeToString(mdSum)
	base64String := base64.StdEncoding.EncodeToString(mdSum)

	fmt.Println("  hexString:", hexString)

	fmt.Println("base64String:", base64String)

	return base64String
}
