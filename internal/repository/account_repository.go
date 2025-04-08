package repository

import (
	"database/sql"

	"github.com/patsferrer/go-gateway/internal/domain"
)

type AccountRepository struct {
	// conexão com banco de dados
	db *sql.DB
}

// função construtora
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// método para salvar uma conta no banco de dados
func (r *AccountRepository) Save(account *domain.Account) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// executa a query, retorna a quantidade de linhas afetadas e o erro
	// o _ é usado para ignorar o retorno da quantidade de linhas afetadas
	_, err = stmt.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.APIKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)
	// verifica se houve erro
	if err != nil {
		return err
	}
	// se não houve erro, retorna nil
	// o nil é um valor nulo em Go
	return nil
}
