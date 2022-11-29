package variable

import (
	"context"
	"fmt"
	"time"

	gus "github.com/michalq/gus-stats/pkg/client_gus"
	"github.com/michalq/gus-stats/pkg/limiter"
)

const pageSize = 100

type Finder struct {
	variablesApi gus.VariablesApi
	limiter      limiter.Limiter
}

func NewFinder(variablesApi gus.VariablesApi, limiter limiter.Limiter) *Finder {
	return &Finder{variablesApi, limiter}
}

func (f *Finder) FindAll(ctx context.Context) ([]*Variable, error) {
	variables := make([]*Variable, 0)
	var i int32 = 0
	for {
		f.limiter.Wait(ctx)
		variablesRequest := f.variablesApi.
			VariablesGet(ctx).
			Format("json").
			PageSize(pageSize).
			Page(i)
		list, _, err := variablesRequest.Execute()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(list.Results) == 0 {
			break
		}
		fmt.Printf("[%s] Page %d, results %d\n", time.Now(), i, len(list.Results))
		for _, apiVariable := range list.Results {
			var name string = *apiVariable.N1
			if apiVariable.N2 != nil {
				name += " " + *apiVariable.N2
			}
			if apiVariable.N3 != nil {
				name += " " + *apiVariable.N2
			}
			if apiVariable.N4 != nil {
				name += " " + *apiVariable.N2
			}
			if apiVariable.N5 != nil {
				name += " " + *apiVariable.N2
			}
			variables = append(variables, &Variable{
				Id:              *apiVariable.Id,
				SubjectId:       *apiVariable.SubjectId,
				Name:            name,
				Level:           int(*apiVariable.Level),
				MeasureUnitId:   int(*apiVariable.MeasureUnitId),
				MeasureUnitName: *apiVariable.MeasureUnitName,
			})
		}
		i++
	}

	return variables, nil
}
