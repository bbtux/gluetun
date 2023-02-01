package nordvpn

import (
	"github.com/qdm12/gluetun/internal/provider/nordvpn/api"
	"math/rand"
	"net/http"

	"github.com/qdm12/gluetun/internal/constants/providers"
	"github.com/qdm12/gluetun/internal/provider/common"
	"github.com/qdm12/gluetun/internal/provider/utils"
)

type Provider struct {
	storage    common.Storage
	randSource rand.Source
	utils.NoPortForwarder
	api    api.NordvpnApi
	warner common.Warner
}

func New(storage common.Storage, randSource rand.Source, client *http.Client, updaterWarner common.Warner) *Provider {
	return &Provider{
		storage:         storage,
		randSource:      randSource,
		NoPortForwarder: utils.NewNoPortForwarding(providers.Nordvpn),
		api:             api.New(client),
		warner:          updaterWarner,
	}
}

func (p *Provider) Name() string {
	return providers.Nordvpn
}
