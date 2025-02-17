package constants

import (
	"sort"

	"github.com/qdm12/gluetun/internal/models"
)

//nolint:lll
const (
	CyberghostCertificate = "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"
)

func CyberghostRegionChoices(servers []models.CyberghostServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Region
	}
	return makeUnique(choices)
}

func CyberghostGroupChoices(servers []models.CyberghostServer) (choices []string) {
	uniqueChoices := map[string]struct{}{}
	for _, server := range servers {
		uniqueChoices[server.Group] = struct{}{}
	}

	choices = make([]string, 0, len(uniqueChoices))
	for choice := range uniqueChoices {
		choices = append(choices, choice)
	}

	sortable := sort.StringSlice(choices)
	sortable.Sort()

	return sortable
}

func CyberghostHostnameChoices(servers []models.CyberghostServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Hostname
	}
	return makeUnique(choices)
}
