package services

import (
	"github.com/sanzharanarbay/golang_excel_example/internal/models"
	"github.com/sanzharanarbay/golang_excel_example/internal/repositories"
	"log"
	"strconv"
	"time"
)

type FoodService struct {
	foodRepository *repositories.FoodRepository
}

func NewFoodService(foodRepository *repositories.FoodRepository) *FoodService {
	return &FoodService{
		foodRepository: foodRepository,
	}
}

func (f *FoodService) Save(data [][]string) ([]interface{}, error) {
	var foods []interface{}
	for i := 0; i < len(data); i++ {
		if i == 0 {
			continue
		}
		var food models.Food
		for j := 0; j < len(data[i]); j++ {
			switch j {
			case 0:
				date, err := time.Parse("01-02-06", data[i][j])
				if err != nil {
					log.Println(err)
					panic(err)
				}
				food.OrderDate = date
			case 1:
				food.Region = data[i][j]
			case 2:
				food.City = data[i][j]
			case 3:
				food.Category = data[i][j]
			case 4:
				food.Product = data[i][j]
			case 5:
				quantity, err := strconv.ParseInt(data[i][j], 10, 64)
				if err != nil {
					panic(err)
				}
				food.Quantity = quantity
			case 6:
				unitPrice, err := strconv.ParseFloat(data[i][j], 64)
				if err != nil {
					panic(err)
				}
				food.UnitPrice = unitPrice
			case 7:
				totalPrice, err := strconv.ParseFloat(data[i][j], 64)
				if err != nil {
					panic(err)
				}
				food.TotalPrice = totalPrice
				food.CreatedAt = time.Now()
			}
		}
		foods = append(foods, food)
	}

	state, err := f.foodRepository.SaveFoods(foods)

	return state, err
}

func (f *FoodService) Get(ID string) (*models.Food, error) {
	state, err := f.foodRepository.GetFood(ID)
	return state, err
}

func (f *FoodService) List() (*[]models.Food, error) {
	state, err := f.foodRepository.GetAllFoods()
	return state, err
}
