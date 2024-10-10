package security

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 2).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
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

func ExtractUser(tokenString string) (string, error) {

	token, _, _ := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if username, ok := claims["username"].(string); ok {
			return username, nil
		}
		return "", fmt.Errorf("username not found in token ")
	}

	return "", fmt.Errorf("invalid token claims")
}

func VerifyCookie(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("jwt_token")
	if err != nil {
		http.Error(w, "Unauthorized: No valid cookie", http.StatusUnauthorized)
		return
	}

	jwtToken := cookie.Value

	if err := VerifyToken(jwtToken); err != nil {
		http.Error(w, "invalid token, good try hacker", http.StatusUnauthorized)
	}
	w.WriteHeader(http.StatusOK)
}
