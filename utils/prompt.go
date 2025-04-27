package utils

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

// PromptPassword prompts the user for a password input (hidden)
func PromptPassword(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Mask:  '*',
	}

	password, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return password, nil
}

// PromptTextWithDefault prompts the user for a text input with a default value
func PromptTextWithDefault(label string, defaultValue string) string {
	prompt := promptui.Prompt{
		Label:   fmt.Sprintf("%s (default: %s)", label, defaultValue),
		Default: defaultValue,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
