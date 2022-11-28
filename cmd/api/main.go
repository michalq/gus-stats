package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/michalq/gus-stats/internal/api/handler"
	"github.com/michalq/gus-stats/internal/config"
	"github.com/michalq/gus-stats/internal/gus"
	"github.com/michalq/gus-stats/internal/subject"
	"github.com/michalq/gus-stats/internal/variable"
)

func main() {
	cfg := config.LoadConfig()
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

	gusCli := gus.NewClient(cfg)

	r := gin.Default()
	r.GET("/subjects", func(c *gin.Context) {
		handler.SubjectsHandler(c, "", subjectsFlatArray[0], subjectsMap)
	})
	r.GET("/subjects/:subjectId", func(c *gin.Context) {
		handler.SubjectsHandler(c, c.Param("subjectId"), subjectsFlatArray[0], subjectsMap)
	})
	r.GET("/subjects/:subjectId/variables", func(c *gin.Context) {
		handler.VariablesHandler(c, variablesBySubject)
	})
	r.GET("/subjects/:subjectId/variables/:variableId/data", func(c *gin.Context) {
		handler.DataHandler(c, c.Param("subjectId"), c.Param("variableId"), gusCli.DataApi)
	})
	r.Run(":3000")
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
