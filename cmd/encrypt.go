package cmd

import (
	"log"
	"os"

	password "github.com/GabrielNexT/next-encrypt/internal"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/chacha20"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand(cmd, args)
	},
}

func runCommand(cmd *cobra.Command, args []string) {
	filePath, _ := cmd.Flags().GetString("file")

	src, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err.Error())
	}

	key, nonce := password.GetKeyAndNonceFromPassword()

	cipher, _ := chacha20.NewUnauthenticatedCipher(key, nonce)

	dst := make([]byte, len(src))

	cipher.XORKeyStream(dst, src)

	f, err := os.Create(filePath + ".nc")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	f.Write(dst)
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	encryptCmd.Flags().String("file", "", "File for encrypt")
}
