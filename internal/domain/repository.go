// interface para o repositório de contas
package domain

type AccountRepository interface {
	Save(account *Account)
	FindByAPIKey(apiKey string) (*Account, error)
	FindById(id string) (*Account, error)
	Update(account *Account) error
}
