package crypto

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func randomID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Intn(1_000_000))
}

func TestEncryptDecrypt(t *testing.T) {
	id := randomID()

	originalFile := fmt.Sprintf(".env.test-%s", id)
	encryptedFile := fmt.Sprintf(".env.enc.test-%s", id)
	decryptedFile := fmt.Sprintf(".env.decrypted-%s", id)

	password := "correct_password"
	wrongPassword := "wrong_password"

	testContent := "DATABASE_URL=postgres://user:pass@localhost:5432/db\nSECRET_KEY=mysecret"

	// Step 1: Create test file
	err := os.WriteFile(originalFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(originalFile) // clean up

	// Step 2: Encrypt
	err = EncryptFile(originalFile, encryptedFile, password)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}
	defer os.Remove(encryptedFile) // clean up

	// Step 3: Decrypt with correct password
	err = DecryptFile(encryptedFile, decryptedFile, password)
	if err != nil {
		t.Fatalf("Decryption with correct password failed: %v", err)
	}
	defer os.Remove(decryptedFile) // clean up

	// Step 4: Compare files
	originalData, err := os.ReadFile(originalFile)
	if err != nil {
		t.Fatalf("Failed to read original file: %v", err)
	}

	decryptedData, err := os.ReadFile(decryptedFile)
	if err != nil {
		t.Fatalf("Failed to read decrypted file: %v", err)
	}

	if string(originalData) != string(decryptedData) {
		t.Fatalf("Decrypted data does not match original!\nOriginal: %s\nDecrypted: %s", originalData, decryptedData)
	}

	// Step 5: Decrypt with wrong password (expect failure)
	err = DecryptFile(encryptedFile, decryptedFile, wrongPassword)
	if err == nil {
		t.Fatalf("Decryption should have failed with wrong password, but it succeeded")
	}
}
