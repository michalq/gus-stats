package model

type SubjectsResponse struct {
	Id        string                     `json:"id"`
	Name      string                     `json:"name"`
	Ancestors []SubjectsResponseAncestor `json:"ancestors"`
	Children  []SubjectsResponseChild    `json:"children"`
	Links     SubjectsResponseLinks      `json:"links"`
}

type SubjectsResponseLinks struct {
	Parent    *string `json:"$parent,omitempty"`
	Variables *string `json:"$variables,omitempty"`
}

type SubjectsResponseChild struct {
	Id          string                     `json:"id"`
	Name        string                     `json:"name"`
	ChildrenQty int                        `json:"children_qty"`
	Links       SubjectsResponseChildLinks `json:"links"`
}

type SubjectsResponseAncestor struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SubjectsResponseChildLinks struct {
	Self string `json:"$self"`
}
