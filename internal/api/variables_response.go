package api

type VariablesResponse []VariablesResponseVariables

type VariablesResponseVariables struct {
	Id   string `json:"id"`
	Data string `json:"$data"`
	Name string `json:"name"`
}
