package utilities

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocrud/config"
)

func GetPagination(c *gin.Context) (limit int, offset int) {
	lim, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		lim = config.GetConfig().Pagination.Limit
	}

	off, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		off = config.GetConfig().Pagination.Offset
	}

	return lim, off
}
