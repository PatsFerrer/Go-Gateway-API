// interface para o repositório de contas
package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByAPIKey(apiKey string) (*Account, error)
	FindById(id string) (*Account, error)
	UpdateBalance(account *Account) error
}

type InvoiceRepository interface {
	Save(invoice *Invoice) error
	FindById(id string) (*Invoice, error)
	FindByAccountId(accountID string) ([]*Invoice, error)
	UpdateStatus(invoice *Invoice) error
}
