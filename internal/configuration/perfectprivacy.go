package configuration

import (
	"fmt"

	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/golibs/params"
)

func (settings *Provider) readPerfectPrivacy(r reader) (err error) {
	settings.Name = constants.Perfectprivacy
	servers := r.servers.GetPerfectprivacy()

	settings.ServerSelection.TargetIP, err = readTargetIP(r.env)
	if err != nil {
		return err
	}

	settings.ServerSelection.Cities, err = r.env.CSVInside("CITY", constants.PerfectprivacyCityChoices(servers))
	if err != nil {
		return fmt.Errorf("environment variable CITY: %w", err)
	}

	return settings.ServerSelection.OpenVPN.readPerfectPrivacy(r.env)
}

func (settings *OpenVPNSelection) readPerfectPrivacy(env params.Interface) (err error) {
	settings.TCP, err = readProtocol(env)
	if err != nil {
		return err
	}

	settings.CustomPort, err = readOpenVPNCustomPort(env, settings.TCP,
		[]uint16{142, 152, 300, 301, 1142, 1152},
		[]uint16{148, 149, 150, 151, 1148, 1149, 1150, 1151})
	if err != nil {
		return err
	}

	return nil
}
