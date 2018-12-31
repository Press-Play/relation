package utils

import (
    "fmt"
    "crypto/rand"
)

func GenerateToken() string {
    b := make([]byte, 24)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}
