package api

import (
	"encoding/json"
	"fmt"
	"github.com/qdm12/gluetun/internal/constants/vpn"
	"github.com/qdm12/gluetun/internal/models"
	"io"
	"net/http"
)

type server struct {
	Id           uint32
	Groups       []group
	Hostname     string
	Ips          []ip
	Load         int
	Locations    []location
	Name         string
	Status       string
	Technologies []technology
}

func (api *api) GetServers(recommended bool) (servers []models.Server, err error) {
	var url string
	if recommended {
		url = baseUrl + "servers/recommendations?limit=10000"
	} else {
		url = baseUrl + "servers?limit=10000"
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status code not OK: %s", res.Status)
	}

	return parse(res.Body)
}

func parse(rdr io.Reader) (servers []models.Server, err error) {
	var jsonServers []server
	decoder := json.NewDecoder(rdr)
	if err := decoder.Decode(&jsonServers); err != nil {
		return nil, fmt.Errorf("failed unmarshaling resp body: %w", err)
	}

	servers = make([]models.Server, 0, len(jsonServers))
	for _, jsonServer := range jsonServers {
		server := models.Server{
			Number: jsonServer.Id, Hostname: jsonServer.Hostname, ServerName: jsonServer.Name,
			Region: jsonServer.region(), Country: jsonServer.country(), City: jsonServer.city(),
			IPs: jsonServer.ips(),
		}

		tcp := jsonServer.tcp()
		udp := jsonServer.udp()
		if tcp || udp {
			openVpn := server
			openVpn.VPN = vpn.OpenVPN
			openVpn.TCP = tcp
			openVpn.UDP = udp
			servers = append(servers, openVpn)
		}

		wgPub := jsonServer.wgPub()
		if jsonServer.wireguard() && len(wgPub) > 0 {
			wireguard := server
			wireguard.VPN = vpn.Wireguard
			wireguard.WgPubKey = wgPub
			servers = append(servers, wireguard)
		}

	}

	return servers, nil
}
