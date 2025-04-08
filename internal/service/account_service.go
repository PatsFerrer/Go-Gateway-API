// service permite que eu tenha os métodos para criar e pegar os dados das contas
package service

import (
	"github.com/patsferrer/go-gateway/internal/domain"
	"github.com/patsferrer/go-gateway/internal/dto"
)

type AccountService struct {
	// usa a interface para acessar o banco de dados
	repository domain.AccountRepository
}

// função construtora
func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

// função para criar uma conta
func (s *AccountService) CreateAccount(input *dto.CreateAccountInput) (*dto.AccountOutput, error) {

	// converte o dto em um domínio
	account := dto.ToAccount(input)

	// verifica se a conta já existe baseado na APIKey
	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)

	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}

	// se a conta já existe e não é um erro de conta não encontrada, retorna um erro
	if existingAccount != nil {
		return nil, domain.ErrDuplicateAPIKey
	}

	// salva a account no banco de dados
	err = s.repository.Save(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return output, nil
}

// função para atualizar o saldo da conta
func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	// verifica se a conta existe
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	// atualiza o saldo da conta
	account.AddBalance(amount)

	// salva o saldo atualizado
	err = s.repository.UpdateBalance(account)

	if err != nil {
		return nil, err
	}

	// converte o domínio em um dto
	output := dto.FromAccount(account)
	return output, nil
}
