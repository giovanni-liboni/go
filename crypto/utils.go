package crypto

import (
	"math/rand"
	"strings"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func RandStringRunes() string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 50)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		if unicode.IsSymbol(r) {
			return -1
		}
		if strings.IndexRune(".!$", r) == 0 {
			return -1
		}
		return r
	}, str)
}

func GetHash(text string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(RandStringRunes()), bcrypt.DefaultCost)
	temp := strings.Replace(string(hashedPassword), "/", "", -1)
	return SpaceMap(temp)
}
