package perfectprivacy

import (
	"strconv"

	"github.com/qdm12/gluetun/internal/configuration"
	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/utils"
)

func (p *Perfectprivacy) BuildConf(connection models.Connection,
	settings configuration.OpenVPN) (lines []string, err error) {
	if settings.Cipher == "" {
		settings.Cipher = constants.AES256cbc
	}

	if settings.Auth == "" {
		settings.Auth = constants.SHA512
	}

	lines = []string{
		"client",
		"dev " + settings.Interface,
		"nobind",
		"persist-key",
		"remote-cert-tls server",
		"ping-timer-rem",
		"tls-exit",

		// Perfect Privacy specific
		"tun-mtu 1500",
		"reneg-sec 3600",
		"comp-lzo",
		"key-direction 1",
		"ping 5",
		"ping-restart 120",
		"route-delay 2",
		"route-method exe",
		"hand-window 120",
		"inactive 604800",
		"tls-cipher TLS-DHE-RSA-WITH-AES-256-GCM-SHA384:TLS-DHE-RSA-WITH-AES-256-CBC-SHA256:TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA:TLS-DHE-RSA-WITH-AES-256-CBC-SHA:TLS-RSA-WITH-CAMELLIA-256-CBC-SHA:TLS-RSA-WITH-AES-256-CBC-SHA", //nolint:lll
		"tls-timeout 5",

		// Added constant values
		"auth-nocache",
		"mute-replay-warnings",
		"pull-filter ignore \"auth-token\"", // prevent auth failed loops
		"auth-retry nointeract",
		"suppress-timestamps",

		// Modified variables
		"verb " + strconv.Itoa(settings.Verbosity),
		"auth-user-pass " + constants.OpenVPNAuthConf,
		connection.OpenVPNProtoLine(),
		connection.OpenVPNRemoteLine(),
		"auth " + settings.Auth,
	}

	lines = append(lines, utils.CipherLines(settings.Cipher, settings.Version)...)

	if settings.MSSFix > 0 {
		mssFixLine := "mssfix " + strconv.Itoa(int(settings.MSSFix))
		lines = append(lines, mssFixLine)
	} else if connection.Protocol == constants.UDP {
		lines = append(lines, "mssfix") // perfect privacy specific
	}

	if !settings.Root {
		lines = append(lines, "user "+settings.ProcUser)
	}

	if settings.IPv6 {
		lines = append(lines, "tun-ipv6")
	} else {
		lines = append(lines, `pull-filter ignore "route-ipv6"`)
		lines = append(lines, `pull-filter ignore "ifconfig-ipv6"`)
	}

	lines = append(lines, utils.WrapOpenvpnCA(
		constants.PerfectprivacyCA)...)
	lines = append(lines, utils.WrapOpenvpnCert(
		connection.OpenVPNCert)...)
	lines = append(lines, utils.WrapOpenvpnKey(
		connection.OpenVPNKey)...)
	lines = append(lines, utils.WrapOpenvpnTLSAuth(
		connection.OpenVPNTLSAuth)...)

	lines = append(lines, "")

	return lines, nil
}
