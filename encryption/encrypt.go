package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// NOTE: Key should be 16, 24 or 32 bytes long, resulting in AES-128, AES-192 or AES-256 encryption, respectively.
func Encrypt(text string, key []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func Decrypt(encrypted string, key []byte) string {
	block, _ := aes.NewCipher(key)
	decoded, _ := base64.StdEncoding.DecodeString(encrypted)
	iv := decoded[:aes.BlockSize]
	decoded = decoded[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decoded, decoded)
	return string(decoded)
}
