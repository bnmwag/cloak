package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/bnmwag/cloak/crypto"
	"github.com/bnmwag/cloak/utils"
)

var DecryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt a file with a password",
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")

		// Prompt for missing input/output
		if input == "" {
			input = utils.PromptTextWithDefault("Input file", ".env.enc")
		}
		if output == "" {
			output = utils.PromptTextWithDefault("Output file", ".env")
		}

		password, err := utils.PromptPassword("Enter password for decryption:")
		if err != nil {
			fmt.Println("❌ Failed to read password:", err)
			os.Exit(1)
		}

		err = crypto.DecryptFile(input, output, password)
		if err != nil {
			fmt.Println("❌ Decryption failed:", err)
			os.Exit(1)
		}

		fmt.Println("✅ File decrypted successfully:", output)
	},
}

func init() {
	DecryptCmd.Flags().StringP("input", "i", "", "Path to input file")
	DecryptCmd.Flags().StringP("output", "o", "", "Path to output file")
}
