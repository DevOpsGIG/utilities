package utilities

import (
	cryptorand "crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var errRandomZero = errors.New("Random: function expects int greater than zero as input")

// RandomInt return a pseudo-random between zero and max
func RandomInt(max int) (int, error) {
	if max <= 0 {
		return 0, errRandomZero
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max), nil
}

// UUID generates a random UUID according to RFC 4122
func UUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(cryptorand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
