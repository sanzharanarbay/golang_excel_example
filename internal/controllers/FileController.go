package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sanzharanarbay/golang_excel_example/internal/services"
	"github.com/sanzharanarbay/golang_excel_example/internal/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type FileController struct {
	fileService *services.FileService
	foodService *services.FoodService
}

func NewFileController(fileService *services.FileService, foodService *services.FoodService) *FileController {
	return &FileController{
		fileService: fileService,
		foodService: foodService,
	}
}

func (fi *FileController) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	allowedFileExtensions := []string{".xlsx", ".xls"}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	extension := filepath.Ext(file.Filename)
	checkFormat := utils.StringContains(allowedFileExtensions, extension)
	if !checkFormat {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "File Extension not satisfied !!!",
		})
		return
	}

	cwd, _ := os.Getwd()
	_, err = os.Open("uploads")
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("uploads", os.ModePerm) // you might want different file access, this suffice for this example
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("Created %s at %s\n", "new-dir", cwd)
			}
		}
	}
	newFileName := uuid.New().String() + extension
	path := filepath.Join(cwd, "uploads", newFileName)
	newFilePath := filepath.FromSlash(path)

	if err := ctx.SaveUploadedFile(file, newFilePath); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
			"error":   err.Error(),
		})
		return
	}

	data, err := fi.fileService.ReadExcel(newFilePath)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
			"error":   err.Error(),
		})
		return
	}

	foods, err := fi.foodService.Save(data)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file content",
			"error":   err.Error(),
		})
		return
	}

	if err == nil {
		log.Println("the content of the file was successfully inserted!!!")
	}

	ctx.JSON(http.StatusOK, foods)
}
