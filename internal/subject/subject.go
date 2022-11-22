package subject

type Subject struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Variables bool       `json:"variables"`
	Children  []*Subject `json:"children"`
}

func (s *Subject) Id() string {
	return s.ID
}
