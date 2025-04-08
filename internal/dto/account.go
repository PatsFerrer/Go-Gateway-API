package dto

import "time"

// dto que vou receber da requisição
type CreateAccount struct {
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
