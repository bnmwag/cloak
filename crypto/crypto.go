package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

const (
	saltSize = 16
	keySize  = 32
	ivSize   = 12
	iter     = 100_000
)

func EncryptFile(inputPath, outputPath, password string) error {
	plaintext, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}

	key := pbkdf2.Key([]byte(password), salt, iter, keySize, sha256.New)

	iv := make([]byte, ivSize)
	if _, err := rand.Read(iv); err != nil {
		return fmt.Errorf("failed to generate IV: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %w", err)
	}

	ciphertext := aesgcm.Seal(nil, iv, plaintext, nil)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	if _, err := outFile.Write(salt); err != nil {
		return err
	}
	if _, err := outFile.Write(iv); err != nil {
		return err
	}
	if _, err := outFile.Write(ciphertext); err != nil {
		return err
	}

	return nil
}

func DecryptFile(inputPath, outputPath, password string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	if len(data) < saltSize+ivSize {
		return fmt.Errorf("file too small to contain salt, IV, and ciphertext")
	}

	salt := data[:saltSize]
	iv := data[saltSize : saltSize+ivSize]
	ciphertext := data[saltSize+ivSize:]

	key := pbkdf2.Key([]byte(password), salt, iter, keySize, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %w", err)
	}

	plaintext, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("decryption failed: %w", err)
	}

	err = os.WriteFile(outputPath, plaintext, 0644)
	if err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}
