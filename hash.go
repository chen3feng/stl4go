package stl4go

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
)

func Md5(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

func Hash1(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

// Hash256 return the sha256 of the data passed
func Hash256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// Hash512 return the sha512 of the data passed
func Hash512(data []byte) []byte {
	h := sha512.New()
	h.Write(data)
	return h.Sum(nil)
}

// GenHashInts generate n hash values by the seed passed
func GenHashInts(seed []byte, n int) []uint64 {
	data := seed
	var hashInts []uint64
	for len(hashInts) < n {
		data := Hash512(data)
		temp := make([]uint64, len(data)/8)
		buf := bytes.NewBuffer(data)
		binary.Read(buf, binary.BigEndian, temp)
		hashInts = append(hashInts, temp...)
	}
	return hashInts[:n]
}
