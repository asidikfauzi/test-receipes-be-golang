package ingredient

import (
	"errors"
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"math"
	"net/http"
)

func (m *MasterIngredient) GetAllIngredients(c *gin.Context) {
	var requestPaginate models.Pagination

	page, limit, offset, err := utils.Pagination(
		c,
		requestPaginate.Page,
		requestPaginate.Limit,
	)
	if err != nil {
		return
	}

	data, totalData, err := m.IngredientDatabase.GetIngredients(offset, limit)
	if err != nil {
		utils.FailOnError(err, "Error database GetIngredients")
	}

	totalPages := int(math.Ceil(float64(totalData) / float64(limit)))

	responsePaginate := models.Pagination{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalData:  int(totalData),
	}

	response := models.ResponseWithPagination{
		Code:       http.StatusOK,
		Message:    "Get All Ingredients Successfully",
		Data:       data,
		Pagination: responsePaginate,
	}

	c.JSON(http.StatusOK, response)
}

func (m *MasterIngredient) GetIngredientById(c *gin.Context) {
	id := c.Param("id")
	data, err := m.IngredientDatabase.GetIngredientById(id)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	response := models.Response{
		Code:    http.StatusOK,
		Message: "Get Ingredient By ID Successfully",
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}

func (m *MasterIngredient) CreateIngredient(c *gin.Context) {
	var (
		requestIngredient models.IngredientRequest
		err               error
	)

	if err = c.ShouldBindJSON(&requestIngredient); err != nil {
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

	err = m.IngredientDatabase.CheckExists(requestIngredient.IngredientName)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	err = m.IngredientDatabase.InsertIngredient(requestIngredient)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	response := models.Response{
		Code:    http.StatusCreated,
		Message: "Successfully Add Ingredient!",
	}

	c.JSON(http.StatusCreated, response)
}

func (m *MasterIngredient) UpdateIngredient(c *gin.Context) {
	var (
		requestIngredient models.IngredientRequest
		err               error
	)

	id := c.Param("id")
	_, err = m.IngredientDatabase.GetIngredientById(id)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err = c.ShouldBindJSON(&requestIngredient); err != nil {
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

	err = m.IngredientDatabase.CheckExistsById(id, requestIngredient.IngredientName)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	err = m.IngredientDatabase.UpdateIngredient(id, requestIngredient)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	response := models.Response{
		Code:    http.StatusOK,
		Message: "Successfully Update Ingredient!",
	}

	c.JSON(http.StatusOK, response)
}

func (m *MasterIngredient) DeleteIngredient(c *gin.Context) {
	var err error
	id := c.Param("id")

	_, err = m.IngredientDatabase.GetIngredientById(id)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	err = m.IngredientDatabase.DeleteIngredient(id)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	response := models.Response{
		Code:    http.StatusOK,
		Message: "Successfully Delete Ingredient!",
	}

	c.JSON(http.StatusOK, response)
}
