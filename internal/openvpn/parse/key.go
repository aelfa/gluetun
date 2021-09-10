package parse

import (
	"fmt"
)

func ExtractPrivateKey(b []byte) (keyData string, err error) {
	keyData, err = extractPEM(b, "PRIVATE KEY")
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrExtractPEM, err)
	}

	return keyData, nil
}

func ExtractPrivateKeyFromConfig(config []byte) (keyData string, err error) {
	block, err := extractBlock(config, "key")
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrExtractBlock, err)
	}

	return ExtractPrivateKey(block)
}
