package auth

import (
	"API/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GerarToken gera um token guardando id do usuário logado
func GerarToken(id int) (string, error) {
	claims := jwt.MapClaims{"id": id, "criacao": time.Now().UnixNano()}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.ChaveSecreta)
}

// ValidarToken valida o JWT recebido
func ValidarToken(tokenString string) error {
	_, erro := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica se o método de assinatura é HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metodo de assinatura invalido: %v", token.Header["alg"])
		}
		return config.ChaveSecreta, nil
	})

	if erro != nil {
		return erro
	}
	// Se o token for válido, retorna nil
	return nil
}

// Extrairid extrai id do usuario logado e valida o token
func Extrairid(tokenString string) (int, error) {
	// Parse o token para decodificá-lo e validá-lo
	token, erro := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica se o método de assinatura é HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metodo de assinatura invalido: %v", token.Header["alg"])
		}
		return config.ChaveSecreta, nil
	})

	if erro != nil {
		return 0, erro
	}

	// Verifica se o token é válido e acessa os claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extrai a matrícula do usuário dos claims
		if id, ok := claims["id"].(float64); ok { // JSON decodifica números como float64
			return int(id), nil
		}
		return 0, fmt.Errorf("campo 'id' não encontrado ou invalido")
	}

	return 0, fmt.Errorf("token invalido")
}

// ExtrairToken extrai o token do cabeçalho da requisição
func ExtrairToken(r *http.Request) (string, error) {
	// Obtém o valor do header Authorization
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return "", errors.New("header Authorization ausente")
	}

	// Verifica se o formato é "Bearer <token>"
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("formato do header Authorization invalido")
	}

	// Retorna o token
	return parts[1], nil
}
