package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/michalq/gus-stats/internal/api/model"
	"github.com/michalq/gus-stats/internal/variable"
)

func VariablesHandler(c *gin.Context, variablesBySubject map[string][]variable.Variable) {
	subjectId := c.Param("subjectId")
	variables := make([]model.VariablesResponseVariables, 0)
	for _, gusVar := range variablesBySubject[subjectId] {
		varId := strconv.Itoa(int(gusVar.Id))
		variables = append(variables, model.VariablesResponseVariables{
			Id:   varId,
			Name: gusVar.Name,
			Links: model.VariablesResponseVariablesLinks{
				Data: createApiUrl("/subjects/%s/variables/%s/data", subjectId, varId),
			},
		})
	}
	c.JSON(http.StatusOK, &model.ApiReponse[model.VariablesResponse]{Data: model.VariablesResponse{
		Links: model.VariablesResponseLinks{
			Subject: createApiUrl("/subjects/%s", subjectId),
		},
		Variables: variables,
	}})
}
