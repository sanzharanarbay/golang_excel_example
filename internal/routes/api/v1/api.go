package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzharanarbay/golang_excel_example/internal/configs/mongo"
	"github.com/sanzharanarbay/golang_excel_example/internal/controllers"
	"github.com/sanzharanarbay/golang_excel_example/internal/repositories"
	"github.com/sanzharanarbay/golang_excel_example/internal/services"
)

func ApiRoutes(prefix string, router *gin.Engine) {
	mongoDB := mongo.NewMongoDB()
	apiGroup := router.Group(prefix)
	{
		foodRepo := repositories.NewFoodRepository(mongoDB)
		foodService := services.NewFoodService(foodRepo)
		foodController := controllers.NewFoodController(foodService)

		foodsGroup := apiGroup.Group("/foods")
		{
			foodsGroup.GET("/:id", foodController.GetFood)
			foodsGroup.GET("/all", foodController.GetFoodsList)
		}

		fileService := services.NewFileService()
		fileController := controllers.NewFileController(fileService, foodService)

		filesGroup := apiGroup.Group("/files")
		{
			filesGroup.POST("/upload", fileController.UploadFile)
		}
	}
}

