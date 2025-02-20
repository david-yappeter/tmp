package seeder

import (
	"context"

	"myapp/model"
	"myapp/repository"
)

var (
	UserJohnDoe = model.User{
		Id:       "442e75b9-dd93-4c02-ba8f-26e9e6c00b6c",
		Name:     "John Doe",
		Email:    "email@gmail.com",
		Password: "$2a$10$UfN/zrbj5M5u1A5zwEeA6ufWFZTY9Gl5UUzeNvwU67xXYpa0KMASy", // 123456
	}
)

func UserSeeder(repositoryManager repository.RepositoryManager) {
	userRepository := repositoryManager.UserRepository()

	count, err := userRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := userRepository.InsertMany(context.Background(), getUserData()); err != nil {
		panic(err)
	}
}

func getUserData() []model.User {
	return []model.User{
		UserJohnDoe,
	}
}
