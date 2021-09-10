package parse

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_extractBlock(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		config []byte
		key    string
		block  []byte
		err    error
	}{
		"no input": {
			err: errors.New("start string not found: <>"),
		},
		"start not found": {
			key:    "key",
			config: []byte(`<other>value</key>`),
			err:    errors.New("start string not found: <key>"),
		},
		"end not found": {
			key:    "key",
			config: []byte(`<key>value</other>`),
			err:    errors.New("end string not found: </key>"),
		},
		"success": {
			key:    "key",
			config: []byte(`other<key>value 1</key>other`),
			block:  []byte(`value 1`),
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			block, err := extractBlock(testCase.config, testCase.key)

			if testCase.err != nil {
				require.Error(t, err)
				assert.Equal(t, testCase.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, testCase.block, block)
		})
	}
}
