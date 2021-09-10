package parse

import (
	"fmt"
)

func ExtractStaticKeyV1(b []byte) (staticKeyV1Data string, err error) {
	staticKeyV1Data, err = extractPEM(b, "OpenVPN Static key V1")
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrExtractPEM, err)
	}

	return staticKeyV1Data, nil
}

func ExtractStaticKeyV1FromConfig(config []byte) (certData string, err error) {
	block, err := extractBlock(config, "tls-auth")
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrExtractBlock, err)
	}

	return ExtractStaticKeyV1(block)
}
