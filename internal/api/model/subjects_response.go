package model

type SubjectsResponse struct {
	Id       string                  `json:"id"`
	Name     string                  `json:"name"`
	Children []SubjectsResponseChild `json:"children"`
	Links    SubjectsResponseLinks   `json:"links"`
}

type SubjectsResponseLinks struct {
	Parent    *string `json:"$parent,omitempty"`
	Variables *string `json:"$variables,omitempty"`
}

type SubjectsResponseChild struct {
	Id    string                     `json:"id"`
	Name  string                     `json:"name"`
	Links SubjectsResponseChildLinks `json:"links"`
}

type SubjectsResponseChildLinks struct {
	Self string `json:"$self"`
}
