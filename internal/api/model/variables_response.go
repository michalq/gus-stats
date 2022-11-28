package model

type VariablesResponse struct {
	Links     VariablesResponseLinks       `json:"links"`
	Variables []VariablesResponseVariables `json:"variables"`
}

type VariablesResponseLinks struct {
	Subject string `json:"$subject"`
}

type VariablesResponseVariables struct {
	Id    string                          `json:"id"`
	Links VariablesResponseVariablesLinks `json:"links"`
	Name  string                          `json:"name"`
}

type VariablesResponseVariablesLinks struct {
	Data string `json:"$data"`
}
