package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	mu        sync.RWMutex // mutex para evitar que o saldo seja atualizado simultaneamente, para não dar erro de calculo (concorrência)
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Gera uma API Key aleatória
func generateAPIKey() string {
	sliceDeBytes := make([]byte, 16)
	rand.Read(sliceDeBytes)
	return hex.EncodeToString(sliceDeBytes)
}

// Como se fosse a função construtora
func NewAccount(name, email string) *Account {
	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Balance:   0,
		APIKey:    generateAPIKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

// método para adicionar saldo à conta, se chama de método porque é uma função que pertence a um tipo (Account)
func (a *Account) AddBalance(amount float64) {
	a.mu.Lock()         // bloqueia o acesso ao saldo
	defer a.mu.Unlock() // libera o acesso ao saldo
	a.Balance += amount
	a.UpdatedAt = time.Now()
}
