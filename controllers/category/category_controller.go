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

func (m *MasterCategory) GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	data, err := m.CategoryDatabase.GetCategoryById(id)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	response := models.Response{
		Code:    200,
		Message: "Get Category By ID Successfully",
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}
