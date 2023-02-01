package api

type location struct {
	Country struct {
		City struct {
			Name string
		}
		Code string
		Name string
	}
}

func (s *server) country() string {
	if s.Locations == nil || len(s.Locations) == 0 {
		return ""
	}
	for _, location := range s.Locations {
		if &location.Country != nil {
			return location.Country.Name
		}
	}
	return ""
}

func (s *server) city() string {
	if s.Locations == nil || len(s.Locations) == 0 {
		return ""
	}
	for _, location := range s.Locations {
		if &location.Country != nil && &location.Country.City != nil {
			return location.Country.City.Name
		}
	}
	return ""
}
