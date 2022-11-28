package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/michalq/gus-stats/internal/api/model"
	gus "github.com/michalq/gus-stats/pkg/client_gus"
)

func DataHandler(c *gin.Context, subjectId, variableId string, dataApi gus.DataApi) {
	varId, _ := strconv.Atoi(variableId)
	// TODO Pagination!!
	data, _, err := dataApi.DataByVariableGet(context.TODO(), int32(varId)).PageSize(100).Execute()
	if err != nil {
		log.Println(err)
		return
	}
	buckets := make([]model.DataResponseBucket, 0, len(data.Results))
	for _, res := range data.Results {
		bucketData := make([]model.DataResponseData, 0, len(res.Values))
		for _, resData := range res.Values {
			bucketData = append(bucketData, model.DataResponseData{
				Arg:   *resData.Year,
				Value: *resData.Val,
			})
		}
		buckets = append(buckets, model.DataResponseBucket{
			Name: *res.Name,
			Data: bucketData,
		})
	}
	c.JSON(http.StatusOK, &model.ApiReponse[model.DataResponse]{
		Data: model.DataResponse{
			Buckets: buckets,
			Links: model.DataResponseLinks{
				Variables: createApiUrl("/subjects/%s/variables", subjectId),
			},
		},
	})
}
