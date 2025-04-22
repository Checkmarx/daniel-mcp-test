package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	// This is vulnerable to the "none" algorithm attack
	tokenString := "dnskdnmdlsms"

	// Vulnerable verification - doesn't properly check signing method
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// WARNING: This is vulnerable because it doesn't verify the signing method
		// An attacker can use the "none" algorithm to bypass signature verification
		return []byte("secret"), nil
	})


	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Name:", claims["name"])
		fmt.Println("Subject:", claims["sub"])
	}
}