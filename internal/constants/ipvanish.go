package constants

import (
	"github.com/qdm12/gluetun/internal/models"
)

//nolint:lll
const (
	IpvanishCA = "MIIErTCCA5WgAwIBAgIJAMYKzSS8uPKDMA0GCSqGSIb3DQEBDQUAMIGVMQswCQYDVQQGEwJVUzELMAkGA1UECBMCRkwxFDASBgNVBAcTC1dpbnRlciBQYXJrMREwDwYDVQQKEwhJUFZhbmlzaDEVMBMGA1UECxMMSVBWYW5pc2ggVlBOMRQwEgYDVQQDEwtJUFZhbmlzaCBDQTEjMCEGCSqGSIb3DQEJARYUc3VwcG9ydEBpcHZhbmlzaC5jb20wHhcNMTIwMTExMTkzMjIwWhcNMjgxMTAyMTkzMjIwWjCBlTELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkZMMRQwEgYDVQQHEwtXaW50ZXIgUGFyazERMA8GA1UEChMISVBWYW5pc2gxFTATBgNVBAsTDElQVmFuaXNoIFZQTjEUMBIGA1UEAxMLSVBWYW5pc2ggQ0ExIzAhBgkqhkiG9w0BCQEWFHN1cHBvcnRAaXB2YW5pc2guY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAt9DBWNr/IKOuY3TmDP5x7vYZR0DGxLbXU8TyAzBbjUtFFMbhxlHiXVQrZHmgzih94x7BgXM7tWpmMKYVb+gNaqMdWE680Qm3nOwmhy/dulXDkEHAwD05i/iTx4ZaUdtV2vsKBxRg1vdC4AEiwD7bqV4HOi13xcG971aQ55Mj1KeCdA0aNvpat1LWx2jjWxsfI8s2Lv5Fkoi1HO1+vTnnaEsJZrBgAkLXpItqP29Lik3/OBIvkBIxlKrhiVPixE5qNiD+eSPirsmROvsyIonoJtuY4Dw5K6pcNlKyYiwo1IOFYU3YxffwFJk+bSW4WVBhsdf5dGxq/uOHmuz5gdwxCwIDAQABo4H9MIH6MAwGA1UdEwQFMAMBAf8wHQYDVR0OBBYEFEv9FCWJHefBcIPX9p8RHCVOGe6uMIHKBgNVHSMEgcIwgb+AFEv9FCWJHefBcIPX9p8RHCVOGe6uoYGbpIGYMIGVMQswCQYDVQQGEwJVUzELMAkGA1UECBMCRkwxFDASBgNVBAcTC1dpbnRlciBQYXJrMREwDwYDVQQKEwhJUFZhbmlzaDEVMBMGA1UECxMMSVBWYW5pc2ggVlBOMRQwEgYDVQQDEwtJUFZhbmlzaCBDQTEjMCEGCSqGSIb3DQEJARYUc3VwcG9ydEBpcHZhbmlzaC5jb22CCQDGCs0kvLjygzANBgkqhkiG9w0BAQ0FAAOCAQEAI2dkh/43ksV2fdYpVGhYaFZPVqCJoToCez0IvOmLeLGzow+EOSrY508oyjYeNP4VJEjApqo0NrMbKl8g/8bpLBcotOCF1c1HZ+y9v7648uumh01SMjsbBeHOuQcLb+7gX6c0pEmxWv8qj5JiW3/1L1bktnjW5Yp5oFkFSMXjOnIoYKHyKLjN2jtwH6XowUNYpg4qVtKU0CXPdOznWcd9/zSfa393HwJPeeVLbKYaFMC4IEbIUmKYtWyoJ9pJ58smU3pWsHZUg9Zc0LZZNjkNlBdQSLmUHAJ33Bd7pJS0JQeiWviC+4UTmzEWRKa7pDGnYRYNu2cUo0/voStphv8EVA=="
)

func IpvanishCountryChoices(servers []models.IpvanishServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Country
	}
	return makeUnique(choices)
}

func IpvanishCityChoices(servers []models.IpvanishServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].City
	}
	return makeUnique(choices)
}

func IpvanishHostnameChoices(servers []models.IpvanishServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Hostname
	}
	return makeUnique(choices)
}
