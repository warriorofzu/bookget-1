package zhash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash/crc32"
)

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	re := h.Sum(nil)
	o := fmt.Sprintf("%x", re)
	return o
}

//生成sha1
func SHA1(str string) string {
	c := sha1.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

func CRC32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
