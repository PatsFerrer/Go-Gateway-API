package domain

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
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
