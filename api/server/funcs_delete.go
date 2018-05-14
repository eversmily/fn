package server

import (
	"net/http"

	"github.com/fnproject/fn/api"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleFuncsDelete(c *gin.Context) {
	ctx := c.Request.Context()

	err := s.datastore.RemoveFunc(ctx, c.MustGet(api.Func).(string))
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted func"})
}
