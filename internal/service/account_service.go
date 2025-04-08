// service permite que eu tenha os métodos para criar e pegar os dados das contas
package service

import "github.com/patsferrer/go-gateway/internal/domain"

type AccountService struct {
	// usa a interface para acessar o banco de dados
	repository domain.AccountRepository
}

// função construtora
func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}




