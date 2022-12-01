package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michalq/gus-stats/internal/api/model"
	"github.com/michalq/gus-stats/internal/domain/subject"
)

func SubjectsTreeHandler(c *gin.Context, root *subject.Subject) {
	c.JSON(http.StatusOK, &model.ApiReponse[*subject.Subject]{
		Data: root,
	})
}
