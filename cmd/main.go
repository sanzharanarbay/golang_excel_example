package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	v1 "github.com/sanzharanarbay/golang_excel_example/internal/routes/api/v1"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")

	prefix := os.Getenv("ROUTE_PREFIX")
	mode := os.Getenv("GIN_MODE")

	fmt.Println("Server started at " + port + "...")
	gin.SetMode(mode)
	router := gin.New()
	router.MaxMultipartMemory = 8 << 20
	v1.ApiRoutes(prefix, router)
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}

}
