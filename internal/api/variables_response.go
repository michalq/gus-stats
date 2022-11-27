package api

type VariablesResponse []VariablesResponseVariables

type VariablesResponseVariables struct {
	Id    string                          `json:"id"`
	Links VariablesResponseVariablesLinks `json:"links"`
	Name  string                          `json:"name"`
}

type VariablesResponseVariablesLinks struct {
	Subjects string `json:"$subjects"`
	Data     string `json:"$data"`
}
