package dto

import (
	"time"

	"github.com/patsferrer/go-gateway/internal/domain"
)

// dto que vou receber da requisição
type CreateAccountInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// dto que vou enviar como resposta
// se quiser omitir o campo APIKey na resposta, adicione o tag 'omitempty'
// ex: APIKey    string    `json:"api_key,omitempty"`
type AccountOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	APIKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// função para converter o dto (que é o JSON que vou receber da requisição) em um domínio (o domínio é o modelo de dados que vou usar no banco de dados)
func ToAccount(input *CreateAccountInput) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

// função para converter o domínio em um dto
func FromAccount(account *domain.Account) *AccountOutput {
	return &AccountOutput{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		Balance:   account.Balance,
		APIKey:    account.APIKey,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
