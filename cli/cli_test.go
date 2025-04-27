package cli

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"testing"
	"time"
)

func randomID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Intn(1_000_000))
}

func TestCLIEncryptDecrypt(t *testing.T) {
	id := randomID()

	originalFile := fmt.Sprintf(".env.test-%s", id)
	encryptedFile := fmt.Sprintf(".env.enc.test-%s", id)
	decryptedFile := fmt.Sprintf(".env.decrypted-%s", id)

	password := "correct_password"

	// Step 1: Create a test .env file
	testContent := "DATABASE_URL=postgres://user:pass@localhost:5432/db\nSECRET_KEY=mysecret"
	err := os.WriteFile(originalFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(originalFile)
	defer os.Remove(encryptedFile)
	defer os.Remove(decryptedFile)

	// Step 2: Encrypt via CLI
	cmdEncrypt := exec.Command("../cloak", "encrypt", "--input", originalFile, "--output", encryptedFile)
	cmdEncrypt.Stdin = stringReader(password + "\n") // simulate password typing
	cmdEncrypt.Stderr = os.Stderr
	cmdEncrypt.Stdout = os.Stdout

	err = cmdEncrypt.Run()
	if err != nil {
		t.Fatalf("Encrypt CLI command failed: %v", err)
	}

	// Step 3: Decrypt via CLI
	cmdDecrypt := exec.Command("../cloak", "decrypt", "--input", encryptedFile, "--output", decryptedFile)
	cmdDecrypt.Stdin = stringReader(password + "\n")
	cmdDecrypt.Stderr = os.Stderr
	cmdDecrypt.Stdout = os.Stdout

	err = cmdDecrypt.Run()
	if err != nil {
		t.Fatalf("Decrypt CLI command failed: %v", err)
	}

	// Step 4: Compare original and decrypted files
	originalData, err := os.ReadFile(originalFile)
	if err != nil {
		t.Fatalf("Failed to read original file: %v", err)
	}

	decryptedData, err := os.ReadFile(decryptedFile)
	if err != nil {
		t.Fatalf("Failed to read decrypted file: %v", err)
	}

	if string(originalData) != string(decryptedData) {
		t.Fatalf("Mismatch between original and decrypted data!\nOriginal: %s\nDecrypted: %s", originalData, decryptedData)
	}
}

// stringReader simulates user input for exec.Command
func stringReader(s string) *os.File {
	r, w, _ := os.Pipe()
	go func() {
		w.Write([]byte(s))
		w.Close()
	}()
	return r
}
