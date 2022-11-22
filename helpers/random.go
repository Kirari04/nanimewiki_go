package helpers

import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const intsBytes = "1234567890"

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
func RandIntBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = intsBytes[rand.Int63()%int64(len(intsBytes))]
	}
	return string(b)
}
