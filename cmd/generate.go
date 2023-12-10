/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var length int
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Run:   generatePass,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntVarP(&length, "length", "l", 10, "Password length")
}

func generatePass(_ *cobra.Command, _ []string) {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var password string
	for i := 0; i < length; i++ {
		password += string(characters[rand.Intn(len(characters))])
	}
	fmt.Println(password)
}
