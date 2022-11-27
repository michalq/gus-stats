package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/michalq/gus-stats/internal/api"
	"github.com/michalq/gus-stats/internal/variable"
)

func main() {
	variables, err := loadVariables()
	if err != nil {
		panic(err)
	}
	variablesBySubject := transformBySubject(variables)

	r := gin.Default()
	r.GET("/subjects", func(c *gin.Context) {
		c.JSON(http.StatusOK, &api.ApiReponse[string]{Data: "ok"})
	})
	r.GET("/subjects/:subjectId/variables", func(c *gin.Context) {
		subjectId := c.Param("subjectId")
		variables := make([]api.VariablesResponseVariables, 0)
		for _, gusVar := range variablesBySubject[subjectId] {
			varId := strconv.Itoa(int(gusVar.Id))
			variables = append(variables, api.VariablesResponseVariables{
				Id:   varId,
				Name: gusVar.Name,
				Links: api.VariablesResponseVariablesLinks{
					Subjects: createApiUrl("/subjects"),
					Data:     createApiUrl("/subjects/%s/variables/%s/data", subjectId, varId),
				},
			})
		}
		c.JSON(http.StatusOK, &api.ApiReponse[api.VariablesResponse]{Data: variables})
	})
	r.GET("/subjects/:subjectId/variables/:variableId/data", func(c *gin.Context) {
		subjectId := c.Param("subjectId")
		c.JSON(http.StatusOK, &api.ApiReponse[api.DataResponse]{
			Data: api.DataResponse{
				Links: api.DataResponseLinks{
					Subject: createApiUrl("/subjects/%s/variables", subjectId),
				},
			},
		})
	})
	r.Run(":3000")
}

func createApiUrl(endpoint string, params ...any) string {
	return "http://localhost:3000" + fmt.Sprintf(endpoint, params...)
}

func loadVariables() ([]variable.Variable, error) {
	variablesRaw, err := os.ReadFile("data/variables.json")
	if err != nil {
		return nil, err
	}
	var variables []variable.Variable
	if err := json.Unmarshal(variablesRaw, &variables); err != nil {
		return nil, err
	}
	return variables, nil
}

func transformBySubject(variables []variable.Variable) map[string][]variable.Variable {
	variablesBySubject := make(map[string][]variable.Variable)
	for _, gusVar := range variables {
		variablesBySubject[gusVar.SubjectId] = append(variablesBySubject[gusVar.SubjectId], gusVar)
	}
	return variablesBySubject
}
