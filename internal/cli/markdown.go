package cli

import (
	"errors"
	"flag"
	"fmt"

	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/storage"
)

type MarkdownFormatter interface {
	FormatToMarkdown(args []string) error
}

var (
	ErrProviderUnspecified = errors.New("VPN provider to format was not specified")
)

func (c *CLI) FormatToMarkdown(args []string) error {
	var cyberghost, fastestvpn, hideMyAss, ipvanish, ivpn, mullvad,
		nordvpn, pia, privado, privatevpn, protonvpn, purevpn, surfshark,
		torguard, vpnUnlimited, vyprvpn, windscribe bool
	flagSet := flag.NewFlagSet("markdown", flag.ExitOnError)
	flagSet.BoolVar(&cyberghost, "cyberghost", false, "Format Cyberghost servers")
	flagSet.BoolVar(&fastestvpn, "fastestvpn", false, "Format FastestVPN servers")
	flagSet.BoolVar(&hideMyAss, "hidemyass", false, "Format HideMyAss servers")
	flagSet.BoolVar(&ipvanish, "ipvanish", false, "Format IpVanish servers")
	flagSet.BoolVar(&ivpn, "ivpn", false, "Format IVPN servers")
	flagSet.BoolVar(&mullvad, "mullvad", false, "Format Mullvad servers")
	flagSet.BoolVar(&nordvpn, "nordvpn", false, "Format Nordvpn servers")
	flagSet.BoolVar(&pia, "pia", false, "Format Private Internet Access servers")
	flagSet.BoolVar(&privado, "privado", false, "Format Privado servers")
	flagSet.BoolVar(&privatevpn, "privatevpn", false, "Format Private VPN servers")
	flagSet.BoolVar(&protonvpn, "protonvpn", false, "Format Protonvpn servers")
	flagSet.BoolVar(&purevpn, "purevpn", false, "Format Purevpn servers")
	flagSet.BoolVar(&surfshark, "surfshark", false, "Format Surfshark servers")
	flagSet.BoolVar(&torguard, "torguard", false, "Format Torguard servers")
	flagSet.BoolVar(&vpnUnlimited, "vpnunlimited", false, "Format VPN Unlimited servers")
	flagSet.BoolVar(&vyprvpn, "vyprvpn", false, "Format Vyprvpn servers")
	flagSet.BoolVar(&windscribe, "windscribe", false, "Format Windscribe servers")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	logger := newNoopLogger()
	storage, err := storage.New(logger, constants.ServersData)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrNewStorage, err)
	}
	currentServers := storage.GetServers()

	var markdown string
	switch {
	case cyberghost:
		markdown = currentServers.Cyberghost.ToMarkdown()
	case fastestvpn:
		markdown = currentServers.Fastestvpn.ToMarkdown()
	case hideMyAss:
		markdown = currentServers.HideMyAss.ToMarkdown()
	case ipvanish:
		markdown = currentServers.Ipvanish.ToMarkdown()
	case ivpn:
		markdown = currentServers.Ivpn.ToMarkdown()
	case mullvad:
		markdown = currentServers.Mullvad.ToMarkdown()
	case nordvpn:
		markdown = currentServers.Nordvpn.ToMarkdown()
	case pia:
		markdown = currentServers.Pia.ToMarkdown()
	case privado:
		markdown = currentServers.Privado.ToMarkdown()
	case privatevpn:
		markdown = currentServers.Privatevpn.ToMarkdown()
	case protonvpn:
		markdown = currentServers.Protonvpn.ToMarkdown()
	case purevpn:
		markdown = currentServers.Purevpn.ToMarkdown()
	case surfshark:
		markdown = currentServers.Surfshark.ToMarkdown()
	case torguard:
		markdown = currentServers.Torguard.ToMarkdown()
	case vpnUnlimited:
		markdown = currentServers.VPNUnlimited.ToMarkdown()
	case vyprvpn:
		markdown = currentServers.Vyprvpn.ToMarkdown()
	case windscribe:
		markdown = currentServers.Windscribe.ToMarkdown()
	default:
		return ErrProviderUnspecified
	}

	fmt.Println(markdown)

	return nil
}
