package parse

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ExtractStaticKeyV1(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		b               []byte
		staticKeyV1Data string
		err             error
	}{
		"bad input": {
			b:   []byte{1, 2, 3},
			err: errors.New("cannot extract PEM data: cannot decode PEM encoded block"),
		},
		"valid key": {
			b:               []byte(validStaticKeyV1PEM),
			staticKeyV1Data: validStaticKeyV1Data,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			staticKeyV1Data, err := ExtractStaticKeyV1(testCase.b)

			if testCase.err != nil {
				require.Error(t, err)
				assert.Equal(t, testCase.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, testCase.staticKeyV1Data, staticKeyV1Data)
		})
	}
}

func Test_ExtractStaticKeyV1FromConfig(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		config          []byte
		staticKeyV1Data string
		err             error
	}{
		"bad config": {
			err: errors.New("cannot extract relevant block: start string not found: <tls-auth>"),
		},
		"bad static key v1": {
			config: []byte("<tls-auth>bad</tls-auth>"),
			err:    errors.New("cannot extract PEM data: cannot decode PEM encoded block"),
		},
		"valid key": {
			config:          []byte("<tls-auth>" + validStaticKeyV1PEM + "</tls-auth>"),
			staticKeyV1Data: validStaticKeyV1Data,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			staticKeyV1Data, err := ExtractStaticKeyV1FromConfig(testCase.config)

			if testCase.err != nil {
				require.Error(t, err)
				assert.Equal(t, testCase.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, testCase.staticKeyV1Data, staticKeyV1Data)
		})
	}
}

const validStaticKeyV1PEM = `
-----BEGIN OpenVPN Static key V1-----
8639991ad6c846ca4c0e8bef909d6acb
ab79cc6e243c93298bb63fff4040661d
b8ac0affcfc1b077d002046c9f7ed813
034787768002d610f122155782b903d6
95ef30ee8640dfc380ac556326b2504c
64d36d594482e7673348eb1921bc1de4
ba6742ac5d85742158e194d03ca8ffeb
8773ca5da548791f19ec3ffb0d2de7b7
ab426d1110743f26d37d26f4f960ae51
606fdc41a38aeff5d8e836313a3338d6
4a0217487234bec25932930f898a69d8
1f698e9b9e40e7611c0aecbc7383fe1c
0ca3128a33d7f427c8a5ff649399d9dd
e339e922dcfbfe5b11cf97c2f0d37c88
5745b7dbfe754e50c509c6d64bead9a3
e1152ee143d4dc70a0186deef93a19f8
-----END OpenVPN Static key V1-----
`
const validStaticKeyV1Data = "8639991ad6c846ca4c0e8bef909d6acbab79cc6e243c93298bb63fff4040661db8ac0affcfc1b077d002046c9f7ed813034787768002d610f122155782b903d695ef30ee8640dfc380ac556326b2504c64d36d594482e7673348eb1921bc1de4ba6742ac5d85742158e194d03ca8ffeb8773ca5da548791f19ec3ffb0d2de7b7ab426d1110743f26d37d26f4f960ae51606fdc41a38aeff5d8e836313a3338d64a0217487234bec25932930f898a69d81f698e9b9e40e7611c0aecbc7383fe1c0ca3128a33d7f427c8a5ff649399d9dde339e922dcfbfe5b11cf97c2f0d37c885745b7dbfe754e50c509c6d64bead9a3e1152ee143d4dc70a0186deef93a19f8" //nolint:lll
