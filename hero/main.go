package main

import (
	"fmt"
	"heromod/models"
)



func main() {
	var productModel models.ProductModel

	fmt.Println("Find the product with the name starting with lap")
	products, err := productModel.FindNameStartsWith("lap")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, product := range products {
			fmt.Println(product.ToString())
			fmt.Println("---------------------------")
		}
	}

	fmt.Println("Find the product with the name ends with le 2")
	products, err = productModel.FindNameEndsWith("le 2")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, product := range products {
			fmt.Println(product.ToString())
			fmt.Println("---------------------------")
		}
	}

	fmt.Println("Find the product with the name contains vi")
	products, err = productModel.FindNameContains("vi")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, product := range products {
			fmt.Println(product.ToString())
			fmt.Println("---------------------------")
		}
	}
}