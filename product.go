package main

var products = []*product{}

type product struct {
	Id    string
	Name  string
	Price int32
}

func GetProducts() []*product {
	return products
}

func SelectProduct(id string) *product {
	for _, each := range products {
		if each.Id == id {
			return each
		}
	}

	return nil
}
