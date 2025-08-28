package security

import "golang.org/x/crypto/bcrypt"

//GerarSenhaComHash gera uma senha com hash a partir de uma senha
func GerarSenhaComHash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

//VerificarSenha compara uma senha string com uma c hash e retorna se são iguais
func VerificarSenha(senhaHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
