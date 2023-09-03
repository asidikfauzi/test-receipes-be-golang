package recipe

import (
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/utils"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func (m *MasterRecipe) GetAllRecipes(c *gin.Context) {
	var requestPaginate models.Pagination

	page, limit, offset, err := utils.Pagination(
		c,
		requestPaginate.Page,
		requestPaginate.Limit,
	)
	if err != nil {
		return
	}

	data, totalData, err := m.RecipeDatabase.GetRecipes(offset, limit)
	if err != nil {
		utils.FailOnError(err, "Error database GetRecipes")
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
		Message:    "Get All Recipes Successfully",
		Data:       data,
		Pagination: responsePaginate,
	}

	c.JSON(http.StatusOK, response)
}

func (m *MasterRecipe) GetIngredientById(c *gin.Context) {
	id := c.Param("id")
	data, err := m.RecipeDatabase.GetRecipeById(id)
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
