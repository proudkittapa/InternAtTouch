package main

import (
	"go.mongodb.org/mongo-driver/bson"
	)

func (this ProductModel) FindNameContains(keyword string) ([]entities.Product, error) {
	db, err := config.GetMongoDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		err2 := db.C("product").Find(bson.M{
			"name": bson.RegEx{
				Pattern: keyword,
				Options: "i",
			},
		}).All(&products)
		if err2 != nil {
			return nil, err
		} else {
				return products, nil
		}
	}
}