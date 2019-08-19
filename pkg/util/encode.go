package util

import (
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
)

func EncodePsw(value string) string {
	dk, _:= scrypt.Key([]byte(value), []byte("salt"), 16384, 8, 1, 32)
	return hex.EncodeToString(dk)
}
