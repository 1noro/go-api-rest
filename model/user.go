package model

import (
    "crypto/sha1"
    "io"
    "encoding/hex"
)

// User representación de un usuario en la DB
type User struct {
    Username string
    ConcatenatedPasswordSha string
    Salt string
}

// CheckPassword comprueba si la contraseña introducida es correcta
func (user User) CheckPassword(passwordSha string) bool {
    if user.calcConcatenatedPasswordSha(passwordSha) == user.ConcatenatedPasswordSha {
        return true
    }
    return false
}


// calcConcatenatedPasswordSha calcula el SHA-1 de la concatenación de la Salt mas la passwordSha
func (user User) calcConcatenatedPasswordSha(passwordSha string) string {
    calculator := sha1.New()
    io.WriteString(calculator, user.Salt + passwordSha)
    hashByteArray := calculator.Sum(nil)
    newConcatenatedPasswordSha := hex.EncodeToString(hashByteArray)
    return newConcatenatedPasswordSha
}