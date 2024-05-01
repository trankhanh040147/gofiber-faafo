package model

type Database struct {
	Uri    string `json:"uri" bson:"uri" binding:"required"`
	Name   string `json:"name" bson:"name"`
	DBName string `json:"db_name" bson:"db_name"`
}
