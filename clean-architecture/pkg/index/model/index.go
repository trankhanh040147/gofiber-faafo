package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Index struct {
	Name       string         `json:"name" bson:"name"`
	Keys       map[string]int `json:"keys" bson:"keys" binding:"required"`
	Collection string         `json:"-" bson:"collection" binding:"required"`
	IsUnique   bool           `json:"is_unique" bson:"is_unique"`
}

func ConvertFromMongoIndex(index primitive.D) *Index {
	indexModel := &Index{}
	//indexModel.Keys = make(map[string]interface{})
	//for _, element := range index {
	//	switch element.Key {
	//	case "name":
	//		indexModel.Name, _ = element.Value.(string)
	//	case "key":
	//		for _, val := range element.Value.(primitive.D) {
	//			indexModel.Keys[val.Key] = val.Value
	//		}
	//	case "collection":
	//		indexModel.Collection, _ = element.Value.(string)
	//	case "is_unique":
	//		indexModel.IsUnique, _ = element.Value.(bool)
	//	}
	//}
	return indexModel
}
