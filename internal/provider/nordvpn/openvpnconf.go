package nordvpn

import (
	"strconv"

	"github.com/qdm12/gluetun/internal/configuration"
	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/utils"
)

func (n *Nordvpn) BuildConf(connection models.Connection,
	settings configuration.OpenVPN) (lines []string, err error) {
	if settings.Cipher == "" {
		settings.Cipher = constants.AES256cbc
	}

	if settings.Auth == "" {
		settings.Auth = constants.SHA512
	}

	if settings.MSSFix == 0 {
		settings.MSSFix = 1450
	}

	lines = []string{
		"client",
		"nobind",
		"tls-exit",
		"dev " + settings.Interface,
		"verb " + strconv.Itoa(settings.Verbosity),

		// Nordvpn specific
		"tun-mtu-extra 32",
		"mssfix " + strconv.Itoa(int(settings.MSSFix)),
		"ping 15",
		"remote-cert-tls server",
		"reneg-sec 0",
		"key-direction 1",
		"auth-user-pass " + constants.OpenVPNAuthConf,
		"auth " + settings.Auth,

		// Added constant values
		"auth-nocache",
		"mute-replay-warnings",
		"pull-filter ignore \"auth-token\"", // prevent auth failed loops
		"auth-retry nointeract",
		"suppress-timestamps",

		// Connection variables
		connection.OpenVPNProtoLine(),
		connection.OpenVPNRemoteLine(),
	}

	lines = append(lines, utils.CipherLines(settings.Cipher, settings.Version)...)

	if connection.Protocol == constants.UDP {
		lines = append(lines, "fast-io")
		lines = append(lines, "explicit-exit-notify")
	}

	if !settings.Root {
		lines = append(lines, "user "+settings.ProcUser)
		lines = append(lines, "persist-tun")
		lines = append(lines, "persist-key")
	}

	if !settings.IPv6 {
		lines = append(lines, `pull-filter ignore "route-ipv6"`)
		lines = append(lines, `pull-filter ignore "ifconfig-ipv6"`)
	}

	lines = append(lines, utils.WrapOpenvpnCA(
		constants.NordvpnCertificate)...)
	lines = append(lines, utils.WrapOpenvpnTLSAuth(
		constants.NordvpnOpenvpnStaticKeyV1)...)

	lines = append(lines, "")

	return lines, nil
}
