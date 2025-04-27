// main.go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/bnmwag/cloak/cmd"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "cloak",
		Short: "Cloak - Securely encrypt and decrypt secret files",
		Long: `Cloak is a simple CLI tool for encrypting and decrypting secret files
like .env files, using password-based encryption (AES-256-GCM).`,
	}

	rootCmd.AddCommand(cmd.EncryptCmd)
	rootCmd.AddCommand(cmd.DecryptCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
