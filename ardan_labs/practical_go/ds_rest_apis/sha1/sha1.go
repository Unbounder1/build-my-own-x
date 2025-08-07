package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(SHA1Sig("http.log.gz"))
}

func SHA1Sig(fileName string) (string, error) {
	if !strings.HasSuffix(fileName, ".gz") {
		return "", nil
	}
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		return "", fmt.Errorf("%q -gzip: %w", fileName, err)
	}
	defer r.Close()

	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q -gzip: %w", fileName, err)
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
