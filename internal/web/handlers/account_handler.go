// package para lidar com as requisições e respostas do usuário, contendo os endpoints
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/patsferrer/go-gateway/internal/dto"
	"github.com/patsferrer/go-gateway/internal/service"
)

// estrutura para o handler
type AccountHandler struct {
	accountService *service.AccountService
}

// função construtora
func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

// o go usa w e r para lidar com as requisições e respostas
// função para criar uma conta
func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	// decodifica o corpo da requisição para o tipo CreateAccountInput
	var input dto.CreateAccountInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// cria a conta
	output, err := h.accountService.CreateAccount(&input)

	// se houver erro, retorna o erro
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// define o tipo de conteúdo da resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// retorna o status code 201 (Created) e a resposta
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
