package controllers // ajuste aqui se o seu package tiver outro nome

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestCriarAdministrador_Sucesso faz um teste automatizado da rota CriarAdministrador,
// verificando apenas o status code de sucesso (201 Created).
func TestCriarAdministrador_Sucesso(t *testing.T) {
	// IMPORTANTE:
	// - Use um galpão que exista no banco (ex.: 1).
	// - Use um email que ainda não tenha sido usado na tabela de administradores.
	//   Se quiser, pode trocar o email manualmente antes de rodar o teste.

	bodyJSON := `{
		"nome": "Administrador Teste",
		"telefone": "53999999999",
		"email": "admin_teste_automatizado3@example.com",
		"senha": "senha123",
		"galpao": 1
	}`

	// Montando a requisição HTTP em memória (como se fosse um POST do cliente)
	req := httptest.NewRequest(http.MethodPost, "/administradores", strings.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")

	// Gravador de resposta (captura status e corpo)
	w := httptest.NewRecorder()

	// Chamando diretamente o handler
	CriarAdministrador(w, req)

	// Verificando o status de sucesso (201 Created)
	if w.Code != http.StatusCreated {
		t.Fatalf("esperava status %d (Created), mas veio %d. Corpo da resposta: %s",
			http.StatusCreated, w.Code, w.Body.String())
	}
}
