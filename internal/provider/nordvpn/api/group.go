package api

type group struct {
	Title string
	Type  struct {
		Identifier string
	}
}

func (s *server) region() string {
	if s.Groups == nil || len(s.Groups) == 0 {
		return ""
	}
	for _, group := range s.Groups {
		if &group.Type != nil && group.Type.Identifier == "regions" {
			return group.Title
		}
	}
	return ""
}
