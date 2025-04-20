package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type InputSetup struct {
	SystemName string
	Expiration int
}

type OutputSetup struct {
	SecretKey  string
	PublicKey  string
	SystemName string
	Expiration int
}

func GenerateSecretKey(size int) (string, error) {
	key := make([]byte, size)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

func GenerateFixedJWT(input InputSetup, secretKey string) (OutputSetup, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": input.SystemName,
		"exp": time.Now().Add(time.Hour * 24 * time.Duration(input.Expiration)).Unix(),
	})

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return OutputSetup{}, err
	}

	return OutputSetup{
		SecretKey:  secretKey,
		PublicKey:  signedToken,
		SystemName: input.SystemName,
		Expiration: input.Expiration,
	}, nil
}

func main() {
	if len(os.Args) < 3 {
		panic("Usage: go run main.go <nome do sistema> <tempo de expiração em dias>")
	}

	systemName := os.Args[1]
	expiration := os.Args[2]

	expirationTime, err := strconv.Atoi(expiration)
	if err != nil {
		panic(err)
	}

	// secretKey, err := GenerateSecretKey(32)
	// if err != nil {
	// 	panic(err)
	// }
	secretKey := "ZCDIEsRYs29XY/gd92GAord2suVn2GI/cYuHgYayHEY="

	input := InputSetup{
		SystemName: systemName,
		Expiration: expirationTime,
	}

	jwt, err := GenerateFixedJWT(input, secretKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("===================================== Output =====================================")
	fmt.Println("Chave secreta:                ", secretKey)
	fmt.Println("Chave pública:                ", jwt.PublicKey)
	fmt.Println("Nome do sistema:              ", systemName)
	fmt.Println("Tempo de expiração (em dias): ", expirationTime)
	fmt.Println("==================================================================================")
}
