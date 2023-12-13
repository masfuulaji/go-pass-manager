/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run:   insertPass,
}

func init() {
	rootCmd.AddCommand(insertCmd)
}

func insertPass(_ *cobra.Command, args []string) {
	destination := args[0]
	fmt.Print("Enter password: ")

	password, err := readInput()
	if err != nil {
		fmt.Println("Error reading password :", err)
		return
	}

	err = savePassword(destination, password)
	if err != nil {
		fmt.Println("Error saving password :", err)
		return
	}

	fmt.Println("\nPassword inserted:", string(password))
}

func readInput() (string, error) {
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}

	return string(password), nil
}

func savePassword(filename, password string) error {
	stringFile := fmt.Sprintf("password/%s.txt", filename)
	err := os.MkdirAll(filepath.Dir(stringFile), 0755)
	if err != nil {
		return err
	}
	file, err := os.Create(stringFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(password)
	if err != nil {
		return err
	}
	return nil
}
