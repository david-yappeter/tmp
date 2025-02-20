package production_seeder

import (
	"fmt"

	"myapp/repository"
)

var Seeders = map[string]func(repositoryManager repository.RepositoryManager){}

func Seed(repositoryManager repository.RepositoryManager, tableName string) {
	if seed, exist := Seeders[tableName]; exist {
		seed(repositoryManager)
	} else {
		fmt.Printf("Seeder for table `%s` not found\n", tableName)
	}
}

func SeedAll(repositoryManager repository.RepositoryManager) {
	seedOrders := []string{}

	for _, tableName := range seedOrders {
		seed, ok := Seeders[tableName]
		if !ok {
			panic(fmt.Errorf("table name %s not found", tableName))
		}

		seed(repositoryManager)
	}
}
