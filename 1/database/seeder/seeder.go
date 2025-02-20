package seeder

import (
	"fmt"

	"myapp/model"
	"myapp/repository"
)

var Seeders = map[string]func(repositoryManager repository.RepositoryManager){
	model.ProductTableName: ProductSeeder,
	model.UserTableName:    UserSeeder,
}

func Seed(repositoryManager repository.RepositoryManager, tableName string) {
	if seed, exist := Seeders[tableName]; exist {
		seed(repositoryManager)
	} else {
		fmt.Printf("Seeder for table `%s` not found\n", tableName)
	}
}

func SeedAll(repositoryManager repository.RepositoryManager) {
	seedOrders := []string{
		model.ProductTableName,
		model.UserTableName,
	}

	for _, tableName := range seedOrders {
		seed, ok := Seeders[tableName]
		if !ok {
			panic(fmt.Errorf("table name %s not found", tableName))
		}

		seed(repositoryManager)
	}
}
