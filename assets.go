package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}
	return nil
}

func getAssetPath(mediaType string) string {
	parts := strings.Split(mediaType, "/")
	var extension string
	if len(parts) != 2 {
		extension = "bin"
	} else {
		extension = parts[1]
	}
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic("failed to generate random bytes")
	}
	id := base64.RawURLEncoding.EncodeToString(b)
	return fmt.Sprintf("%s.%s", id, extension)
}
