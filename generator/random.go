package generator

import "math/rand"

var alphabet = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890_")
const keyLength = 10

func GetRandomKey() string {
	idBytes := make([]byte, keyLength)
	for i := 0; i < len(idBytes); i++ {
		idBytes[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(idBytes)
}
