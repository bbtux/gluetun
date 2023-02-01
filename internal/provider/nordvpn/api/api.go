package api

import (
	"github.com/qdm12/gluetun/internal/models"
	"net/http"
)

type NordvpnApi interface {
	GetServers(recommended bool) (servers []models.Server, err error)
}

const baseUrl = "https://api.nordvpn.com/v1/"

type api struct {
	client *http.Client
}

func New(client *http.Client) NordvpnApi {
	return &api{client: client}
}
