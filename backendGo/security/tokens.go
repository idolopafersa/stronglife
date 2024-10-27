package security

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

// CreateToken genera un JWT con un ID y una expiración de 2 horas.
func CreateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken valida el JWT recibido.
func VerifyToken(tokenString string) error {
	// Eliminar el prefijo "Bearer " si existe
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

// ExtractID extrae el ID del usuario desde el JWT sin verificar su validez.
func ExtractID(tokenString string) (int, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return 0, fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if id, ok := claims["id"].(float64); ok {
			return int(id), nil
		}
		return 0, fmt.Errorf("ID not found in token")
	}

	return 0, fmt.Errorf("invalid token claims")
}

// VerifyCookie verifica el JWT almacenado en una cookie llamada "jwt_token" y extrae el ID.
func VerifyCookie(r *http.Request) (int, error) {
	cookie, err := r.Cookie("jwt_token")
	if err != nil {
		log.Printf("Error: Cookie not valid: %v\n", err)
		return 0, err
	}

	jwtToken := cookie.Value

	if err := VerifyToken(jwtToken); err != nil {
		log.Printf("Error: Token invalid: %v\n", err)
		return 0, err
	}

	// Extraer el ID del JWT después de la verificación
	id, err := ExtractID(jwtToken)
	if err != nil {
		return 0, err
	}

	return id, nil
}
