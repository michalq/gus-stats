package variable

import (
	"context"
	"fmt"
	"log"

	gus "github.com/michalq/gus-stats/pkg/client_gus"
)

type Variable struct {
	Id              int32  `json:"id"`
	SubjectId       string `json:"subjectId"`
	Name            string `json:"name"`
	Level           int    `json:"level"`
	MeasureUnitId   int    `json:"measureUnitId"`
	MeasureUnitName string `json:"measureUnitName"`
}

func PrintVariables(ctx context.Context, dataApi gus.DataApi) {
	variables, _, err := dataApi.DataByUnitGet(ctx, "unitId").Execute()
	if err != nil {
		log.Fatal(err)
	}
	for _, variable := range variables.Results {
		fmt.Printf("| %d | %+v\n", variable.Id, *variable.MeasureUnitId)
	}
}
