package api

type technology struct {
	Identifier string
	Pivot      struct {
		Status string
	}
	Metadata []struct {
		Name  string
		Value string
	}
}

func (s *server) isOnline(id string) bool {
	if s.Technologies == nil || len(s.Technologies) == 0 {
		return false
	}
	for _, technology := range s.Technologies {
		if technology.Identifier == id {
			return technology.Pivot.Status == "online"
		}
	}
	return false
}

func (s *server) tcp() bool {
	return s.isOnline("openvpn_tcp")
}

func (s *server) udp() bool {
	return s.isOnline("openvpn_udp")
}

func (s *server) wireguard() bool {
	return s.isOnline("wireguard_udp")
}

func (s *server) wgPub() string {
	if s.Technologies == nil || len(s.Technologies) == 0 {
		return ""
	}
	for _, technology := range s.Technologies {
		if technology.Metadata == nil || technology.Identifier != "wireguard_udp" {
			continue
		}
		for _, metadata := range technology.Metadata {
			if metadata.Name == "public_key" {
				return metadata.Value
			}
		}
	}
	return ""
}
