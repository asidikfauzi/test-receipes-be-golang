package category

import (
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/utils"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func (m *MasterCategory) GetAllCategories(c *gin.Context) {
	var requestPaginate models.Pagination

	page, limit, offset, err := utils.Pagination(
		c,
		requestPaginate.Page,
		requestPaginate.Limit,
	)
	if err != nil {
		return
	}

	data, totalData, err := m.CategoryDatabase.GetCategories(offset, limit)
	if err != nil {
		utils.FailOnError(err, "Error database GetCategories")
	}

	totalPages := int(math.Ceil(float64(totalData) / float64(limit)))

	responsePaginate := models.Pagination{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalData:  int(totalData),
	}

	response := models.ResponseWithPagination{
		Code:       200,
		Message:    "Get All Categories Successfully",
		Data:       data,
		Pagination: responsePaginate,
	}

	c.JSON(http.StatusOK, response)
}
