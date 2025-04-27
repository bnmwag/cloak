package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/bnmwag/cloak/crypto"
	"github.com/bnmwag/cloak/utils"
)

var EncryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a file with a password",
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")

		// Prompt for missing input/output
		if input == "" {
			input = utils.PromptTextWithDefault("Input file", ".env")
		}
		if output == "" {
			output = utils.PromptTextWithDefault("Output file", ".env.enc")
		}

		password, err := utils.PromptPassword("Enter password for encryption:")
		if err != nil {
			fmt.Println("❌ Failed to read password:", err)
			os.Exit(1)
		}

		err = crypto.EncryptFile(input, output, password)
		if err != nil {
			fmt.Println("❌ Encryption failed:", err)
			os.Exit(1)
		}

		fmt.Println("✅ File encrypted successfully:", output)
	},
}

func init() {
	EncryptCmd.Flags().StringP("input", "i", "", "Path to input file")
	EncryptCmd.Flags().StringP("output", "o", "", "Path to output file")
}
