package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michalq/gus-stats/internal/api/model"
	"github.com/michalq/gus-stats/internal/domain/subject"
)

func SubjectsHandler(c *gin.Context, subjectId string, root *subject.Subject, subjectsMap map[string]*subject.Subject) {
	var sbj *subject.Subject
	if subjectId != "" {
		sbj = subjectsMap[subjectId]
	} else {
		sbj = root
	}
	sbjChildren := make([]model.SubjectsResponseChild, 0)
	for _, sbjChild := range sbj.Children {
		sbjChildren = append(sbjChildren, model.SubjectsResponseChild{
			Id:   sbjChild.ID,
			Name: sbjChild.Name,
			Links: model.SubjectsResponseChildLinks{
				Self: createApiUrl("/subjects/%s", sbjChild.ID),
			},
		})
	}
	var variablesLink *string
	if sbj.Variables {
		variablesLinkStr := createApiUrl("/subjects/%s/variables", sbj.ID)
		variablesLink = &variablesLinkStr
	}
	var parentLink *string
	if sbj.Parent != nil {
		parentLinkStr := createApiUrl("/subjects/%s", sbj.Parent.ID)
		parentLink = &parentLinkStr
	} else if subjectId != "" {
		parentLinkStr := createApiUrl("/subjects")
		parentLink = &parentLinkStr
	}

	c.JSON(http.StatusOK, &model.ApiReponse[model.SubjectsResponse]{
		Data: model.SubjectsResponse{
			Id:   subjectId,
			Name: sbj.Name,
			Links: model.SubjectsResponseLinks{
				Parent:    parentLink,
				Variables: variablesLink,
			},
			Children: sbjChildren,
		},
	})
}
