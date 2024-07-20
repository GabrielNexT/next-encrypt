package internal

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"golang.org/x/crypto/chacha20"
)

func EncryptManyFiles(filePaths []string, cipher *chacha20.Cipher) {
	var wg sync.WaitGroup

	for _, filePath := range filePaths {
		wg.Add(1)
		go EncryptFile(filePath, cipher, &wg)
	}

	wg.Wait()
}

func EncryptFile(filePath string, cipher *chacha20.Cipher, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	src, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Printf("Failed to read file %s: %s\n", filePath, err.Error())
		return
	}

	dst := make([]byte, len(src))

	cipher.XORKeyStream(dst, src)

	f, err := os.Create(getNewFilename(filePath))

	if err != nil {
		fmt.Printf("Failed to create file %s: %s", getNewFilename(filePath), err.Error())
		return
	}

	defer f.Close()

	f.Write(dst)
}

func getNewFilename(filePath string) string {
	hasSuffix := strings.HasSuffix(filePath, ".nc")

	if hasSuffix {
		filePath, _ = strings.CutSuffix(filePath, ".nc")
	} else {
		filePath = filePath + ".nc"
	}

	return filePath
}
