package smTest

import (
	"github.com/tjfoc/gmsm/sm3"
)

func SM3(key string) []byte {
	return sm3.Sm3Sum([]byte(key))
}

// func
