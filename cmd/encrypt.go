package cmd

import (
	"github.com/GabrielNexT/next-encrypt/internal"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/chacha20"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt or decrypt a file using a password",
	Long: `Encrypt or decrypt a file using a password.
	encrypt --file LICENSE
	encrypt --file LICENSE --file README.md
	encrypt --file "LICENSE,README.md"
	encrypt --file "LICENSE.nc,README.md.nc"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand(cmd, args)
	},
}

func runCommand(cmd *cobra.Command, args []string) {
	filePaths, _ := cmd.Flags().GetStringSlice("file")

	if len(filePaths) == 0 {
		return
	}

	key, nonce := internal.GetKeyAndNonceFromPassword()

	cipher, _ := chacha20.NewUnauthenticatedCipher(key, nonce)

	internal.EncryptManyFiles(filePaths, cipher)
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	encryptCmd.Flags().StringSlice("file", []string{}, "File for encrypt")
	encryptCmd.MarkFlagRequired("file")
}
