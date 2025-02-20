package manager

import (
	"myapp/global"
	"myapp/infrastructure"
	jwtInternal "myapp/internal/jwt"
	"myapp/repository"
	"myapp/use_case"
)

type Container struct {
	infrastructureManager infrastructure.InfrastructureManager
	jwt                   jwtInternal.Jwt
	repositoryManager     repository.RepositoryManager
	useCaseManager        use_case.UseCaseManager
}

func NewContainer() *Container {
	container := &Container{}

	container.infrastructureManager = infrastructure.NewInfrastructureManager(global.GetConfig())
	container.repositoryManager = repository.NewRepositoryManager(container.infrastructureManager)

	container.jwt = jwtInternal.NewJwt([]byte(global.GetJwtSecretKey()))

	container.useCaseManager = use_case.NewUseCaseManager(
		container.infrastructureManager,
		container.repositoryManager,
		container.jwt,
	)

	return container
}

func (c *Container) InfrastructureManager() infrastructure.InfrastructureManager {
	return c.infrastructureManager
}

func (c *Container) RepositoryManager() repository.RepositoryManager {
	return c.repositoryManager
}

func (c *Container) UseCaseManager() use_case.UseCaseManager {
	return c.useCaseManager
}

func (c *Container) MigrateDB(migrationDir string, isRollingBack bool, steps int, force *int) error {
	return c.infrastructureManager.MigrateDB(migrationDir, isRollingBack, steps, force)
}

func (c *Container) RefreshDB() error {
	if err := c.infrastructureManager.RefreshDB(); err != nil {
		return err
	}

	if err := c.Close(); err != nil {
		return err
	}

	*c = *NewContainer()

	return nil
}

func (c Container) Close() error {
	if err := c.infrastructureManager.CloseDB(); err != nil {
		return err
	}

	return nil
}
