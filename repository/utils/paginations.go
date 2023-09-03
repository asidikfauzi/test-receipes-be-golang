package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func Pagination(c *gin.Context, page, limit int) (int, int, int, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	paginationPage := strconv.Itoa(page)
	paginationLimit := strconv.Itoa(limit)

	page, err := strconv.Atoi(c.DefaultQuery("page", paginationPage))
	if err != nil {
		BadRequest(c, "Invalid page number format")
		return 0, 0, 0, err
	}

	limit, err = strconv.Atoi(c.DefaultQuery("limit", paginationLimit))
	if err != nil {
		BadRequest(c, "Invalid limit number format")
		return 0, 0, 0, err
	}

	offset := (page - 1) * limit

	return page, limit, offset, nil
}
