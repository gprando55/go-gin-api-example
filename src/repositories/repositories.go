package repositories

import user_repository "golang-gin-rest-api/src/repositories/user"

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	User user_repository.UserRepository
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	// Log        logger.Logger
}

// New cria uma nova instancia dos repositórios
func New(opts Options) *Container {
	return &Container{
		User: user_repository.NewMemoryRepository(),
	}
}
