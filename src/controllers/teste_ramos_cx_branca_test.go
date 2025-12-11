package controllers

import (
	"API/src/config"
	"API/src/models"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// Carrega o .env antes de todos os testes
	config.Carregar()

	// Executa os testes
	exitCode := m.Run()

	os.Exit(exitCode)
}

// ---- apoio para o ramo 4: erro em io.ReadAll ----

type erroReader struct{}

func (e erroReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("erro proposital na leitura do corpo")
}

func (e erroReader) Close() error {
	return nil
}

// ===================================================
// RAMO 1 – SUCESSO → 201 Created
// ===================================================

func TestCriarAdministrador_Ramo1_Sucesso(t *testing.T) {
	// Aqui forçamos o caminho de sucesso
	// Caixa branca: sabemos que todas as validações precisam passar,
	// então usamos valores exatamente nos limites: nome e senha com 2 caracteres.

	// Guardando funções reais para restaurar depois
	origConectarDB := conectarDB
	origCriarAdministrador := criarAdministrador
	defer func() {
		conectarDB = origConectarDB
		criarAdministrador = origCriarAdministrador
	}()

	// Fakes para não depender de BD real:
	conectarDB = func() (*sql.DB, error) {
		return nil, nil // ✅ não precisamos de DB real neste teste
	}
	criarAdministrador = func(a *models.Administrador, db *sql.DB) error {
		// simulamos sucesso sem mexer em banco
		return nil
	}

	bodyJSON := `{
		"nome": "Aa",
		"telefone": "53999999999",
		"email": "sucesso_ramo1@example.com",
		"senha": "Bb",
		"galpao": 1
	}`

	req := httptest.NewRequest(http.MethodPost, "/administradores", strings.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CriarAdministrador(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("esperava status %d (Created), mas veio %d. body: %s",
			http.StatusCreated, w.Code, w.Body.String())
	}
}

// ===================================================
// RAMO 2 – JSON inválido → 400 Bad Request
// ===================================================

func TestCriarAdministrador_Ramo2_JSONInvalido(t *testing.T) {
	// Caixa branca: sabemos que json.Unmarshal é chamado logo após ReadAll,
	// então montamos um JSON mal formado para cair nesse if.

	bodyJSON := `{"nome": "Admin", "email": "admin@example.com" "senha": "123"}`

	req := httptest.NewRequest(http.MethodPost, "/administradores", strings.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CriarAdministrador(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperava status %d (Bad Request) para JSON inválido, mas veio %d. body: %s",
			http.StatusBadRequest, w.Code, w.Body.String())
	}
}

// ===================================================
// RAMO 3 – Validar() falha → 400 Bad Request
// ===================================================

func TestCriarAdministrador_Ramo3_ValidacaoFalha(t *testing.T) {
	// Caixa branca: olhamos o código de Validar()
	// if len(u.Nome) < 2 { ... }
	// Então passamos "A" para garantir que esse ramo seja executado.

	bodyJSON := `{
		"nome": "A",
		"telefone": "53999999999",
		"email": "admin_valido@example.com",
		"senha": "senha123",
		"galpao": 1
	}`

	req := httptest.NewRequest(http.MethodPost, "/administradores", strings.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CriarAdministrador(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperava status %d (Bad Request) para falha de validação, mas veio %d. body: %s",
			http.StatusBadRequest, w.Code, w.Body.String())
	}
}

// ===================================================
// RAMO 4 – io.ReadAll falha → 422 Unprocessable Entity
// ===================================================

func TestCriarAdministrador_Ramo4_ErroLeituraCorpo(t *testing.T) {
	// Caixa branca: sabemos que a função chama io.ReadAll(r.Body)
	// logo no início; trocamos o Body por um tipo que SEMPRE falha.

	req := httptest.NewRequest(http.MethodPost, "/administradores", nil)
	req.Body = erroReader{} // nossa implementação que sempre retorna erro

	w := httptest.NewRecorder()

	CriarAdministrador(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("esperava status %d (Unprocessable Entity) para erro de leitura, mas veio %d. body: %s",
			http.StatusUnprocessableEntity, w.Code, w.Body.String())
	}
}

// ===================================================
// RAMO 5 – ConectarDB falha → 500 Internal Server Error
// ===================================================

func TestCriarAdministrador_Ramo5_ErroConectarDB(t *testing.T) {
	// Caixa branca: ao analisar o código, vemos:
	// db, erro := database.ConectarDB()
	// if erro != nil { ... 500 }
	//
	// Usamos a função-injetável conectarDB para forçar esse erro.

	origConectarDB := conectarDB
	defer func() { conectarDB = origConectarDB }()

	conectarDB = func() (*sql.DB, error) {
		return nil, errors.New("erro proposital em ConectarDB")
	}

	bodyJSON := `{
		"nome": "Admin Ok",
		"telefone": "53999999999",
		"email": "admin_ok@example.com",
		"senha": "senha123",
		"galpao": 1
	}`

	req := httptest.NewRequest(http.MethodPost, "/administradores", strings.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CriarAdministrador(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("esperava status %d (Internal Server Error) para erro de conexão com DB, mas veio %d. body: %s",
			http.StatusInternalServerError, w.Code, w.Body.String())
	}
}

// ===================================================
// RAMO 6 – CriarAdministrador falha → 500 Internal Server Error
// ===================================================

func TestCriarAdministrador_Ramo6_ErroRepositorio(t *testing.T) {
	// Caixa branca: pelo código sabemos que, se repositories.CriarAdministrador
	// retornar erro, o handler devolve 500.
	// Usamos a função-injetável criarAdministrador para simular isso.

	origConectarDB := conectarDB
	origCriarAdministrador := criarAdministrador
	defer func() {
		conectarDB = origConectarDB
		criarAdministrador = origCriarAdministrador
	}()

	conectarDB = func() (*sql.DB, error) {
		return nil, nil // ✅ não precisamos de DB real neste teste
	}
	criarAdministrador = func(a *models.Administrador, db *sql.DB) error {
		return errors.New("erro proposital no repositório")
	}

	bodyJSON := `{
		"nome": "Admin Ok",
		"telefone": "53999999999",
		"email": "erro_repositorio@example.com",
		"senha": "senha123",
		"galpao": 1
	}`

	req := httptest.NewRequest(http.MethodPost, "/administradores", strings.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CriarAdministrador(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("esperava status %d (Internal Server Error) para erro do repositório, mas veio %d. body: %s",
			http.StatusInternalServerError, w.Code, w.Body.String())
	}
}
