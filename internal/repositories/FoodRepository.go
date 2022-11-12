package repositories

import (
	mongo_con "github.com/sanzharanarbay/golang_excel_example/internal/configs/mongo"
	"github.com/sanzharanarbay/golang_excel_example/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FoodRepository struct {
	MongoDB *mongo_con.MongoDB
}

func NewFoodRepository(MongoDB *mongo_con.MongoDB) *FoodRepository {
	return &FoodRepository{
		MongoDB: MongoDB,
	}
}

type FoodRepositoryInterface interface {
	SaveFoods(foods []interface{}) ([]interface{}, error)
	GetFood(ID string) (*models.Food, error)
	GetAllFoods() (*[]models.Food, error)
}

func (f *FoodRepository) SaveFoods(foods []interface{}) ([]interface{}, error) {

	result, err := f.MongoDB.FoodsCollection.InsertMany(f.MongoDB.Context, foods)
	if err != nil {
		return nil, err
	}
	return result.InsertedIDs, nil
}

func (f *FoodRepository) GetFood(ID string) (*models.Food, error) {
	var food models.Food
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	err = f.MongoDB.FoodsCollection.FindOne(f.MongoDB.Context, bson.D{{"_id", objectId}}).Decode(&food)
	if err != nil {
		return nil, err
	}
	return &food, nil
}

func (f *FoodRepository) GetAllFoods() (*[]models.Food, error) {
	var food models.Food
	var foods []models.Food

	cursor, err := f.MongoDB.FoodsCollection.Find(f.MongoDB.Context, bson.D{})
	if err != nil {
		defer cursor.Close(f.MongoDB.Context)
		return nil, err
	}

	for cursor.Next(f.MongoDB.Context) {
		err := cursor.Decode(&food)
		if err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}

	return &foods, nil
}

