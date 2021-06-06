package models

import (
	"config"
	"entities"
	"gopkg.in/mgo.v2/bson"
)

type ProductModel struct {
}

func (this ProductModel) FindNameStartsWith(keyword string) ([]entities.Product, error) {
	db, err := config.GetMongoDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		err2 := db.C("product").Find(bson.M{
			"name": bson.RegEx{
				Pattern: "^" + keyword,
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

func (this ProductModel) FindNameEndsWith(keyword string) ([]entities.Product, error) {
	db, err := config.GetMongoDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		err2 := db.C("product").Find(bson.M{
			"name": bson.RegEx{
				Pattern: keyword + "$",
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