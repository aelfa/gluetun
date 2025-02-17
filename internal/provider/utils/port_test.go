package utils

import (
	"testing"

	"github.com/qdm12/gluetun/internal/configuration"
	"github.com/qdm12/gluetun/internal/constants"
	"github.com/stretchr/testify/assert"
)

func Test_GetPort(t *testing.T) {
	t.Parallel()

	const (
		defaultOpenVPNTCP = 443
		defaultOpenVPNUDP = 1194
		defaultWireguard  = 51820
	)

	testCases := map[string]struct {
		selection configuration.ServerSelection
		port      uint16
	}{
		"default": {
			port: defaultOpenVPNUDP,
		},
		"OpenVPN UDP": {
			selection: configuration.ServerSelection{
				VPN: constants.OpenVPN,
			},
			port: defaultOpenVPNUDP,
		},
		"OpenVPN TCP": {
			selection: configuration.ServerSelection{
				VPN: constants.OpenVPN,
				OpenVPN: configuration.OpenVPNSelection{
					TCP: true,
				},
			},
			port: defaultOpenVPNTCP,
		},
		"OpenVPN custom port": {
			selection: configuration.ServerSelection{
				VPN: constants.OpenVPN,
				OpenVPN: configuration.OpenVPNSelection{
					CustomPort: 1234,
				},
			},
			port: 1234,
		},
		"Wireguard": {
			selection: configuration.ServerSelection{
				VPN: constants.Wireguard,
			},
			port: defaultWireguard,
		},
		"Wireguard custom port": {
			selection: configuration.ServerSelection{
				VPN: constants.Wireguard,
				Wireguard: configuration.WireguardSelection{
					EndpointPort: 1234,
				},
			},
			port: 1234,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			port := GetPort(testCase.selection,
				defaultOpenVPNTCP, defaultOpenVPNUDP, defaultWireguard)

			assert.Equal(t, testCase.port, port)
		})
	}
}
