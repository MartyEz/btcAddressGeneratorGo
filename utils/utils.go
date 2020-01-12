package utils

import (
	"math/rand"
	"strings"
	"time"
)

func RndString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" + "abcdefghijklmnopqrstuvwxyzåäö" + "0123456789")
	length := 25
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
