package settings

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/NY-Daystar/corpos-christie/utils"
)

// GetDefaultSmtp get value of default smtp connection
func GetDefaultSmtp() *Smtp {
	key := []byte("JCXGgBf8vp32esHCJVj9kAH6V5TmjP2f")

	decryptHost, _ := decrypt(key, cypherHost)
	decryptPort, _ := decrypt(key, cypherPort)
	decryptUser, _ := decrypt(key, cypherUser)
	decryptPassword, _ := decrypt(key, cypherPassword)

	portInt, _ := utils.ConvertStringToInt(decryptPort)

	return &Smtp{
		Host:     decryptHost,
		Port:     portInt,
		User:     decryptUser,
		Password: decryptPassword,
	}
}

// encrypt to cypher data with AES engine
func encrypt(key []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// GCM est un mode d'opération pour AES qui offre à la fois confidentialité et intégrité
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext), nil
}

// decrypt to read cypher data with AEG engine
func decrypt(key []byte, ciphertext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	data, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := data[:nonceSize], string(data[nonceSize:])

	cypherByte := []byte(ciphertext)

	plaintext, err := aesGCM.Open(nil, nonce, cypherByte, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

const cypherHost = "7f8a3083559845873cbdbc976f58f4bb29cb91a7ca55f1d29c7660bf45c50ed952b3d8888c58470a65"
const cypherPort = "7e0edf48d899e3a000b993c48ea6a0ecb1d4fbe47cf0189fc7c85f521f0269"
const cypherUser = "0d59859dd37e9f456ff9cb5d51ce19c8319da0133e086866994daeb1e116490e29b3d0b7fae502233369218340c14edfca3d0b744c2a5cfcc086"
const cypherPassword = "0cfb8df5392422e768e6cf37589a0d187e67e603715078d821d89c091d9bb84c242cdc0d7bbb2a9f9d42233f"
