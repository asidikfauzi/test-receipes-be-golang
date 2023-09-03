package category

import (
	"errors"
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (m *MasterCategory) CreateCategory(c *gin.Context) {
	var (
		requestCategory models.CategoryRequest
		err             error
	)

	if err = c.ShouldBindJSON(&requestCategory); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]models.ErrorMessageEmpty, len(ve))
			for i, fe := range ve {
				out[i] = models.ErrorMessageEmpty{Field: fe.Field(), Message: utils.GetErrorMessageEmpty(fe)}
			}
			utils.BadRequest(c, out)
		}
		return
	}

	err = m.CategoryDatabase.CheckExists(requestCategory.CategoryName)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	err = m.CategoryDatabase.InsertCategory(requestCategory)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	response := models.Response{
		Code:    201,
		Message: "Successfully Add Category!",
	}

	c.JSON(http.StatusOK, response)
}
