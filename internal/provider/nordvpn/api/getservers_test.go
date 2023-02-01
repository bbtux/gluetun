package api

import (
	"github.com/qdm12/gluetun/internal/constants/vpn"
	"github.com/qdm12/gluetun/internal/models"
	"io"
	"net"
	"reflect"
	"strings"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		rdr io.Reader
	}
	tests := []struct {
		name        string
		args        args
		wantServers []models.Server
		wantErr     bool
	}{
		{"parse json", struct{ rdr io.Reader }{rdr: strings.NewReader(`
[
	{
	  "cpt": 7,
	  "created_at": "2019-12-03 07:26:00",
	  "groups": [
		{
		  "created_at": "2017-06-13 13:43:00",
		  "id": 11,
		  "identifier": "legacy_standard",
		  "title": "Standard VPN servers",
		  "type": {
			"created_at": "2017-06-13 13:40:17",
			"id": 3,
			"identifier": "legacy_group_category",
			"title": "Legacy category",
			"updated_at": "2017-06-13 13:40:23"
		  },
		  "updated_at": "2017-06-13 13:43:00"
		},
		{
		  "created_at": "2017-06-13 13:43:38",
		  "id": 15,
		  "identifier": "legacy_p2p",
		  "title": "P2P",
		  "type": {
			"created_at": "2017-06-13 13:40:17",
			"id": 3,
			"identifier": "legacy_group_category",
			"title": "Legacy category",
			"updated_at": "2017-06-13 13:40:23"
		  },
		  "updated_at": "2017-06-13 13:43:38"
		},
		{
		  "created_at": "2017-10-27 14:17:17",
		  "id": 19,
		  "identifier": "europe",
		  "title": "Europe",
		  "type": {
			"created_at": "2017-10-27 14:16:30",
			"id": 5,
			"identifier": "regions",
			"title": "Regions",
			"updated_at": "2017-10-27 14:16:30"
		  },
		  "updated_at": "2017-10-27 14:17:17"
		}
	  ],
	  "hostname": "es114.nordvpn.com",
	  "id": 950346,
	  "ips": [
		{
		  "created_at": "2022-11-17 11:34:41",
		  "id": 723986,
		  "ip": {
			"id": 85356,
			"ip": "37.120.199.243",
			"version": 4
		  },
		  "ip_id": 85356,
		  "server_id": 950346,
		  "type": "entry",
		  "updated_at": "2022-11-17 11:34:41"
		}
	  ],
	  "ipv6_station": "",
	  "load": 11,
	  "locations": [
		{
		  "country": {
			"city": {
			  "dns_name": "madrid",
			  "hub_score": 0,
			  "id": 2619989,
			  "latitude": 40.408566,
			  "longitude": -3.69222,
			  "name": "Madrid"
			},
			"code": "ES",
			"id": 202,
			"name": "Spain"
		  },
		  "created_at": "2017-06-15 14:06:47",
		  "id": 113,
		  "latitude": 40.408566,
		  "longitude": -3.69222,
		  "updated_at": "2017-06-15 14:06:47"
		}
	  ],
	  "name": "Spain #114",
	  "services": [
		{
		  "created_at": "2017-03-21 12:00:45",
		  "id": 1,
		  "identifier": "vpn",
		  "name": "VPN",
		  "updated_at": "2017-05-25 13:12:31"
		},
		{
		  "created_at": "2017-05-29 19:38:30",
		  "id": 5,
		  "identifier": "proxy",
		  "name": "Proxy",
		  "updated_at": "2017-05-29 19:38:30"
		}
	  ],
	  "specifications": [
		{
		  "id": 8,
		  "identifier": "version",
		  "title": "Version",
		  "values": [
			{
			  "id": 257,
			  "value": "2.1.0"
			}
		  ]
		}
	  ],
	  "station": "37.120.199.243",
	  "status": "online",
	  "technologies": [
		{
		  "created_at": "2017-03-21 12:00:24",
		  "id": 1,
		  "identifier": "ikev2",
		  "metadata": [],
		  "name": "IKEv2/IPSec",
		  "pivot": {
			"server_id": 950346,
			"status": "online",
			"technology_id": 1
		  },
		  "updated_at": "2017-09-05 14:20:16"
		},
		{
		  "created_at": "2017-05-04 08:03:24",
		  "id": 3,
		  "identifier": "openvpn_udp",
		  "metadata": [],
		  "name": "OpenVPN UDP",
		  "pivot": {
			"server_id": 950346,
			"status": "online",
			"technology_id": 3
		  },
		  "updated_at": "2017-05-09 19:27:37"
		},
		{
		  "created_at": "2017-05-09 19:28:14",
		  "id": 5,
		  "identifier": "openvpn_tcp",
		  "metadata": [],
		  "name": "OpenVPN TCP",
		  "pivot": {
			"server_id": 950346,
			"status": "online",
			"technology_id": 5
		  },
		  "updated_at": "2017-05-09 19:28:14"
		},
		{
		  "created_at": "2017-10-02 12:45:14",
		  "id": 21,
		  "identifier": "proxy_ssl",
		  "metadata": [],
		  "name": "HTTP Proxy (SSL)",
		  "pivot": {
			"server_id": 950346,
			"status": "online",
			"technology_id": 21
		  },
		  "updated_at": "2017-10-02 12:45:14"
		},
		{
		  "created_at": "2017-10-02 12:50:49",
		  "id": 23,
		  "identifier": "proxy_ssl_cybersec",
		  "metadata": [],
		  "name": "HTTP CyberSec Proxy (SSL)",
		  "pivot": {
			"server_id": 950346,
			"status": "online",
			"technology_id": 23
		  },
		  "updated_at": "2017-10-02 12:50:49"
		},
		{
		  "created_at": "2019-02-14 14:08:43",
		  "id": 35,
		  "identifier": "wireguard_udp",
		  "metadata": [
			{
			  "name": "public_key",
			  "value": "IF1FGVSzrUznFVZ+dymIz+6bdlCgsuiT/d6cyapN8lw="
			}
		  ],
		  "name": "Wireguard",
		  "pivot": {
			"server_id": 950346,
			"status": "online",
			"technology_id": 35
		  },
		  "updated_at": "2019-02-14 14:08:43"
		}
	  ],
	  "updated_at": "2023-01-21 04:04:22"
	}
]			`)}, []models.Server{
			{VPN: vpn.OpenVPN, Country: "Spain", Region: "Europe", City: "Madrid", Number: 950346, ServerName: "Spain #114", Hostname: "es114.nordvpn.com", TCP: true, UDP: true, WgPubKey: "", IPs: []net.IP{net.ParseIP("37.120.199.243")}},
			{VPN: vpn.Wireguard, Country: "Spain", Region: "Europe", City: "Madrid", Number: 950346, ServerName: "Spain #114", Hostname: "es114.nordvpn.com", TCP: false, UDP: false, WgPubKey: "IF1FGVSzrUznFVZ+dymIz+6bdlCgsuiT/d6cyapN8lw=", IPs: []net.IP{net.ParseIP("37.120.199.243")}},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotServers, err := parse(tt.args.rdr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotServers, tt.wantServers) {
				t.Errorf("parse() gotServers = %v, want %v", gotServers, tt.wantServers)
			}
		})
	}
}
