package perfectprivacy

import (
	"github.com/qdm12/gluetun/internal/configuration"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/utils"
)

func (p *Perfectprivacy) GetConnection(selection configuration.ServerSelection) (
	connection models.Connection, err error) {
	port := getPort(selection)
	protocol := utils.GetProtocol(selection)

	servers, err := p.filterServers(selection)
	if err != nil {
		return connection, err
	}

	var connections []models.Connection
	for _, server := range servers {
		for _, IP := range server.IPs {
			connection := models.Connection{
				Type:           selection.VPN,
				IP:             IP,
				Port:           port,
				Protocol:       protocol,
				OpenVPNCert:    server.OpenVPNCert,
				OpenVPNKey:     server.OpenVPNKey,
				OpenVPNTLSAuth: server.OpenVPNTLSAuth,
			}
			connections = append(connections, connection)
		}
	}

	return utils.PickConnection(connections, selection, p.randSource)
}

func getPort(selection configuration.ServerSelection) (port uint16) {
	const (
		defaultOpenVPNTCP = 1142
		defaultOpenVPNUDP = 1149
		defaultWireguard  = 51820
	)
	return utils.GetPort(selection, defaultOpenVPNTCP,
		defaultOpenVPNUDP, defaultWireguard)
}
