package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	tokenString := "dnskdnmdlsms"

	// Secure verification - explicitly checks signing method and uses proper key handling
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the key for verification
		return []byte("secret"), nil
	})


	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Explicitly check required claims
		if err := claims.ValidateWithLeeway(jwt.Expected{
			Time: true,
		}, 0); err != nil {
			fmt.Println("Claims validation error:", err)
			return
		}

		fmt.Println("Name:", claims["name"])
		fmt.Println("Subject:", claims["sub"])
	}
}