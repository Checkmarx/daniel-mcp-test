package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	tokenString := "dnskdnmdlsms"
	secretKey := []byte("secret")

	// Secure verification - explicitly checks signing method and uses proper key handling
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify that the signing method is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the key for verification
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Type-safe claims extraction
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Verify required claims
		if err := claims.ValidateWithLeeway(jwt.Expected{
			Time: jwt.NewNumericDate(jwt.TimeFunc()),
		}, 0); err != nil {
			fmt.Println("Claims validation error:", err)
			return
		}

		// Access claims safely
		if name, ok := claims["name"].(string); ok {
			fmt.Println("Name:", name)
		}
		if sub, ok := claims["sub"].(string); ok {
			fmt.Println("Subject:", sub)
		}
	}
}