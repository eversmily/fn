package server

import (
	"net/http"

	"github.com/fnproject/fn/api"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleFuncsGet(c *gin.Context) {
	ctx := c.Request.Context()

	fnc := c.MustGet(api.Func).(string)
	f, err := s.datastore.GetFunc(ctx, fnc)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, funcResponse{"Successfully loaded func", f})
}
