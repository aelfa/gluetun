package configuration

import (
	"fmt"

	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/golibs/params"
)

func (settings *Provider) readVPNUnlimited(r reader) (err error) {
	settings.Name = constants.VPNUnlimited
	servers := r.servers.GetVPNUnlimited()

	settings.ServerSelection.TargetIP, err = readTargetIP(r.env)
	if err != nil {
		return err
	}

	settings.ServerSelection.Countries, err = r.env.CSVInside("COUNTRY", constants.VPNUnlimitedCountryChoices(servers))
	if err != nil {
		return fmt.Errorf("environment variable COUNTRY: %w", err)
	}

	settings.ServerSelection.Cities, err = r.env.CSVInside("CITY", constants.VPNUnlimitedCityChoices(servers))
	if err != nil {
		return fmt.Errorf("environment variable CITY: %w", err)
	}

	settings.ServerSelection.Hostnames, err = r.env.CSVInside("SERVER_HOSTNAME",
		constants.VPNUnlimitedHostnameChoices(servers))
	if err != nil {
		return fmt.Errorf("environment variable SERVER_HOSTNAME: %w", err)
	}

	settings.ServerSelection.FreeOnly, err = r.env.YesNo("FREE_ONLY", params.Default("no"))
	if err != nil {
		return fmt.Errorf("environment variable FREE_ONLY: %w", err)
	}

	settings.ServerSelection.StreamOnly, err = r.env.YesNo("STREAM_ONLY", params.Default("no"))
	if err != nil {
		return fmt.Errorf("environment variable STREAM_ONLY: %w", err)
	}

	return settings.ServerSelection.OpenVPN.readProtocolOnly(r)
}

func (settings *OpenVPN) readVPNUnlimited(r reader) (err error) {
	settings.ClientKey, err = readClientKey(r)
	if err != nil {
		return fmt.Errorf("%w: %s", errClientKey, err)
	}

	settings.ClientCrt, err = readClientCertificate(r)
	if err != nil {
		return fmt.Errorf("%w: %s", errClientCert, err)
	}

	return nil
}
