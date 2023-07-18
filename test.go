package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io/ioutil"
	"log"
)

func main() {
	tokenString, err := GenerateToken("policy.json", "trackingService-maximal", "purpose1", "key.pem", 2)
	if err != nil {
		log.Println(err)
	}
	// Read the ECDSA public key from file
	keyData, err := ioutil.ReadFile("public.pem") // Replace with the path to your ECDSA public key file
	if err != nil {
		log.Fatalf("Error reading public key: %v", err)
	}

	// Parse the ECDSA public key
	publicKey, err := jwt.ParseECPublicKeyFromPEM(keyData)
	if err != nil {
		log.Fatalf("Error parsing public key: %v", err)
	}

	// Verify the token signature using the ECDSA-256 public key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		log.Fatalf("Error verifying token: %v", err)
	}

	if token.Valid {
		fmt.Println("Token signature is valid")
	} else {
		fmt.Println("Token signature is invalid")
	}

	fmt.Println("Here is your token: ", tokenString)
}
