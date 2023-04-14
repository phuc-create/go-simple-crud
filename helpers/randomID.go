package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"
)

func GenerateID() string {
	// Generate a new random
	b := make([]byte, 5)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	id := hex.EncodeToString(b)
	// Add timestamp
	t := time.Now().UnixNano()
	id = id + "-" + strconv.FormatInt(t, 10)
	return id
}
