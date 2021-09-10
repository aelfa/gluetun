package parse

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ExtractPrivateKey(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		b       []byte
		keyData string
		err     error
	}{
		"no input": {
			err: errors.New("cannot extract PEM data: cannot decode PEM encoded block"),
		},
		"bad input": {
			b:   []byte{1, 2, 3},
			err: errors.New("cannot extract PEM data: cannot decode PEM encoded block"),
		},
		"valid key": {
			b:       []byte(validPrivateKeyPEM),
			keyData: validPrivateKeyData,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			keyData, err := ExtractPrivateKey(testCase.b)

			if testCase.err != nil {
				require.Error(t, err)
				assert.Equal(t, testCase.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, testCase.keyData, keyData)
		})
	}
}

func Test_ExtractPrivateKeyFromConfig(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		config  []byte
		keyData string
		err     error
	}{
		"bad config": {
			err: errors.New("cannot extract relevant block: start string not found: <key>"),
		},
		"bad private key": {
			config: []byte("<key>bad</key>"),
			err:    errors.New("cannot extract PEM data: cannot decode PEM encoded block"),
		},
		"valid key": {
			config:  []byte("<key>" + validPrivateKeyPEM + "</key>"),
			keyData: validPrivateKeyData,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			keyData, err := ExtractPrivateKeyFromConfig(testCase.config)

			if testCase.err != nil {
				require.Error(t, err)
				assert.Equal(t, testCase.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, testCase.keyData, keyData)
		})
	}
}

const validPrivateKeyPEM = `
-----BEGIN PRIVATE KEY-----
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
zZjrL52saevO25cigVl+hxcnY8DTpbk=
-----END PRIVATE KEY-----
`
const validPrivateKeyData = "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" //nolint:lll
