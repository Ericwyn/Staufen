package gen

import (
	"crypto/rand"
	"math/big"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const letterUuidBytes = "abcdefghijklmnopqrstuvwxyz01234567890123456789"

func GeneralUuid() string {
	str := ""
	for i := 0; i < 32; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(32)))
		index64 := index.Int64()
		str += letterUuidBytes[int(index64) : int(index64)+1]
	}
	return str
}

func GeneralRandomStr(length int) string {
	str := ""
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
		index64 := index.Int64()
		str += letterBytes[int(index64) : int(index64)+1]
	}
	return str
}

//
//func GeneralUUID() string {
//	// 生成 32 位 UUID
//
//}
