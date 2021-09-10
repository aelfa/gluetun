package perfectprivacy

import (
	"net"

	"github.com/qdm12/gluetun/internal/models"
)

type cityToServer map[string]models.PerfectprivacyServer

func (cts cityToServer) add(city, cert, key, staticKeyV1 string, ips []net.IP) {
	server, ok := cts[city]
	if !ok {
		server.City = city
		server.OpenVPNCert = cert
		server.OpenVPNKey = key
		server.OpenVPNTLSAuth = staticKeyV1
		server.TCP = true
		server.UDP = true
	}

	server.IPs = append(server.IPs, ips...)

	cts[city] = server
}

func (cts cityToServer) toServersSlice() (servers []models.PerfectprivacyServer) {
	servers = make([]models.PerfectprivacyServer, 0, len(cts))
	for _, server := range cts {
		servers = append(servers, server)
	}
	return servers
}
