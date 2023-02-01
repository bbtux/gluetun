package api

import "net"

type ip struct {
	Ip struct {
		Ip string
	}
}

func (s *server) ips() (nips []net.IP) {
	nips = make([]net.IP, 0, len(s.Ips))
	for _, ip := range s.Ips {
		nips = append(nips, net.ParseIP(ip.Ip.Ip))
	}
	return nips
}
