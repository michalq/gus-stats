package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/michalq/gus-stats/internal/api"
	"github.com/michalq/gus-stats/internal/subject"
	"github.com/michalq/gus-stats/internal/variable"
)

func main() {
	variables, err := loadVariables()
	if err != nil {
		panic(err)
	}
	variablesBySubject := transformBySubject(variables)
	subjectsTree, err := loadSubjects()
	if err != nil {
		panic(err)
	}
	subjectsFlatArray := transformSubjectsToFlatArray(subjectsTree)
	subjectsMap := transformSubjectsToMap(subjectsFlatArray)

	r := gin.Default()
	r.GET("/subjects", func(c *gin.Context) {
		subjectsHandler(c, "", subjectsFlatArray[0], subjectsMap)
	})
	r.GET("/subjects/:subjectId", func(c *gin.Context) {
		subjectsHandler(c, c.Param("subjectId"), subjectsFlatArray[0], subjectsMap)
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
					Subject:  createApiUrl("/subjects/%s", subjectId),
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
					Variables: createApiUrl("/subjects/%s/variables", subjectId),
				},
			},
		})
	})
	r.Run(":3000")
}

func createApiUrl(endpoint string, params ...any) string {
	return "http://localhost:3000" + fmt.Sprintf(endpoint, params...)
}

func loadSubjects() (*subject.Subject, error) {
	subjectRaw, err := os.ReadFile("data/subjects.json")
	if err != nil {
		return nil, err
	}
	var subjects subject.Subject
	if err := json.Unmarshal(subjectRaw, &subjects); err != nil {
		return nil, err
	}
	return &subjects, nil
}

func transformSubjectsToMap(subjects []*subject.Subject) map[string]*subject.Subject {
	subjectsMap := make(map[string]*subject.Subject)
	for _, sbj := range subjects {
		subjectsMap[sbj.ID] = sbj
	}
	return subjectsMap
}
func transformSubjectsToFlatArray(gusSubject *subject.Subject) []*subject.Subject {
	subjects := make([]*subject.Subject, 0)
	subjects = append(subjects, gusSubject)
	for _, child := range gusSubject.Children {
		subjects = append(subjects, child)
		child.Parent = gusSubject
		if len(child.Children) > 0 {
			subjects = append(subjects, transformSubjectsToFlatArray(child)...)
		}
	}
	return subjects
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

func subjectsHandler(c *gin.Context, subjectId string, root *subject.Subject, subjectsMap map[string]*subject.Subject) {
	var sbj *subject.Subject
	if subjectId != "" {
		sbj = subjectsMap[subjectId]
	} else {
		sbj = root
	}
	sbjChildren := make([]api.SubjectsResponseChild, 0)
	for _, sbjChild := range sbj.Children {
		sbjChildren = append(sbjChildren, api.SubjectsResponseChild{
			Id:   sbjChild.ID,
			Name: sbjChild.Name,
			Links: api.SubjectsResponseChildLinks{
				Self: createApiUrl("/subjects/%s", sbjChild.ID),
			},
		})
	}
	var variablesLink *string
	if sbj.Variables {
		variablesLinkStr := createApiUrl("/subjects/%s/variables", sbj.ID)
		variablesLink = &variablesLinkStr
	}
	var parentLink string
	if sbj.Parent != nil {
		parentLink = createApiUrl("/subjects/%s", sbj.Parent.ID)
	} else {
		parentLink = createApiUrl("/subjects")
	}

	c.JSON(http.StatusOK, &api.ApiReponse[api.SubjectsResponse]{
		Data: api.SubjectsResponse{
			Id:   subjectId,
			Name: sbj.Name,
			Links: api.SubjectsResponseLinks{
				Parent:    parentLink,
				Variables: variablesLink,
			},
			Children: sbjChildren,
		},
	})
}
