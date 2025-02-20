package seeder

import (
	"context"

	"myapp/model"
	"myapp/repository"
)

var (
	Product1 = model.Product{
		Id:    "be5f5a85-74c4-4961-8f66-e731fa7f8248",
		Name:  "Product A",
		Price: 1000,
	}
	Product2 = model.Product{
		Id:    "41f929e6-7312-47cb-aa3b-719a7e757c73",
		Name:  "Product B",
		Price: 1500,
	}
	Product3 = model.Product{
		Id:    "41662ae0-b676-4e81-87ff-4f374ad6a5d9",
		Name:  "Product C",
		Price: 2500,
	}
	Product4 = model.Product{
		Id:    "c9acbc76-9349-4478-b6ec-b01f2dcc6e59",
		Name:  "Product D",
		Price: 2000,
	}
	Product5 = model.Product{
		Id:    "c75510a7-f960-45d0-bcca-c254ea2f1654",
		Name:  "Product E",
		Price: 1000,
	}
	Product6 = model.Product{
		Id:    "86d76f5f-d672-4456-9211-5e2eaaae44a3",
		Name:  "Product F",
		Price: 1188,
	}
	Product7 = model.Product{
		Id:    "2713af6e-fd26-4820-a02e-eaeedbf239ce",
		Name:  "Product G",
		Price: 1500,
	}
	Product8 = model.Product{
		Id:    "ac0d5aea-69f1-4ff4-87ba-1549d5061d13",
		Name:  "Product H",
		Price: 500,
	}
	Product9 = model.Product{
		Id:    "6e6d5997-c18a-4496-ba58-27308c520a93",
		Name:  "Product I",
		Price: 1000,
	}
	Product10 = model.Product{
		Id:    "6aceeb62-3fb9-484d-a481-1d1e83da8a51",
		Name:  "Product J",
		Price: 1250,
	}
)

func ProductSeeder(repositoryManager repository.RepositoryManager) {
	productRepository := repositoryManager.ProductRepository()

	count, err := productRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := productRepository.InsertMany(context.Background(), getProductData()); err != nil {
		panic(err)
	}
}

func getProductData() []model.Product {
	return []model.Product{
		Product1,
		Product2,
		Product3,
		Product4,
		Product5,
		Product6,
		Product7,
		Product8,
		Product9,
		Product10,
	}
}
