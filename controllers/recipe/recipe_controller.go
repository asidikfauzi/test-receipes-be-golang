package recipe

import (
	"errors"
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (m *MasterRecipe) CreateRecipe(c *gin.Context) {
	var (
		requestRecipe models.RecipeRequest
		err           error
	)

	if err = c.ShouldBindJSON(&requestRecipe); err != nil {
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

	err = m.RecipeDatabase.CheckExists(requestRecipe.RecipeName)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	err = m.RecipeDatabase.InsertRecipe(requestRecipe, requestRecipe.Ingredients)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	response := models.Response{
		Code:    http.StatusCreated,
		Message: "Successfully Add Recipe!",
		Data:    requestRecipe,
	}

	c.JSON(http.StatusCreated, response)
}

func (m *MasterRecipe) UpdateRecipe(c *gin.Context) {
	var (
		requestRecipe models.RecipeRequest
		err           error
	)

	id := c.Param("id")
	_, err = m.RecipeDatabase.GetRecipeById(id)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err = c.ShouldBindJSON(&requestRecipe); err != nil {
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

	err = m.RecipeDatabase.CheckExistsById(id, requestRecipe.RecipeName)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	err = m.RecipeDatabase.UpdateRecipe(id, requestRecipe, requestRecipe.Ingredients)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	response := models.Response{
		Code:    http.StatusCreated,
		Message: "Successfully Update Recipe!",
		Data:    requestRecipe,
	}

	c.JSON(http.StatusCreated, response)
}

func (m *MasterRecipe) DeleteRecipe(c *gin.Context) {
	var err error

	id := c.Param("id")
	_, err = m.RecipeDatabase.GetRecipeById(id)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	err = m.RecipeDatabase.DeleteRecipe(id)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	response := models.Response{
		Code:    http.StatusCreated,
		Message: "Successfully Delete Recipe!",
		Data:    id,
	}

	c.JSON(http.StatusCreated, response)
}
