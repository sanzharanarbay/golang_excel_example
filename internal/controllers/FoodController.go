package controllers

import (
	"github.com/sanzharanarbay/golang_excel_example/internal/models"
	"github.com/sanzharanarbay/golang_excel_example/internal/services"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type FoodController struct {
	foodService *services.FoodService
}

func NewFoodController(foodService *services.FoodService) *FoodController {
	return &FoodController{
		foodService: foodService,
	}
}

func (f *FoodController) GetFood(ctx *gin.Context) {
	var foodObj *models.Food
	var err error
	ID := ctx.Param("id")
	foodObj, err = f.foodService.Get(ID)
	if err != nil {
		log.Printf("ERROR - %s", err)
		response := map[string]interface{}{
			"status":  false,
			"message": err.Error(),
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, foodObj)
}

func (f *FoodController) GetFoodsList(ctx *gin.Context) {
	foodsList, err := f.foodService.List()
	if err != nil {
		response := map[string]interface{}{
			"status":  false,
			"message": err.Error(),
		}
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	ctx.JSON(http.StatusOK, foodsList)
}
