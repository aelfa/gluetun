package parse

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errStartNotFound = errors.New("start string not found")
	errEndNotFound   = errors.New("end string not found")
)

func extractBlock(config []byte, key string) (block []byte, err error) {
	blockString := string(config)
	start := "<" + key + ">"
	end := "</" + key + ">"

	startIndex := strings.Index(blockString, start)
	if startIndex == -1 {
		return nil, fmt.Errorf("%w: %s", errStartNotFound, start)
	}
	blockString = blockString[startIndex+len(start):]

	endIndex := strings.Index(blockString, end)
	if endIndex == -1 {
		return nil, fmt.Errorf("%w: %s", errEndNotFound, end)
	}
	blockString = blockString[:endIndex]

	return []byte(blockString), nil
}
