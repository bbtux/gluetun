package nordvpn

import (
	"context"
	"fmt"
	"github.com/qdm12/gluetun/internal/models"
	"sort"

	"github.com/qdm12/gluetun/internal/provider/common"
)

func (p *Provider) FetchServers(_ context.Context, minServers int) (servers []models.Server, err error) {
	servers, err = p.api.GetServers(false)
	if err != nil {
		return nil, err
	}

	if len(servers) < minServers {
		return nil, fmt.Errorf("%w: %d and expected at least %d",
			common.ErrNotEnoughServers, len(servers), minServers)
	}

	sort.Sort(models.SortableServers(servers))

	return servers, nil
}
