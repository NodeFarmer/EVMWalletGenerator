package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"runtime"

	"github.com/ethereum/go-ethereum/crypto"
)

var (
	Reset     = ""
	Red       = ""
	Green     = ""
	Yellow    = ""
	Blue      = ""
	Cyan      = ""
	Bold      = ""
	Underline = ""
)

func init() {
	// Activer les couleurs uniquement sous les systèmes compatibles
	if runtime.GOOS != "windows" {
		Reset = "\033[0m"
		Red = "\033[31m"
		Green = "\033[32m"
		Yellow = "\033[33m"
		Blue = "\033[34m"
		Cyan = "\033[36m"
		Bold = "\033[1m"
		Underline = "\033[4m"
	}
}

func main() {

	fmt.Println("========================================================")
	fmt.Println("EVM Wallet Generator - v1.0 by NodeFarmer @_node_farmer_")
	fmt.Println("========================================================\n")

	privateKey, publicKey, address, err := generateEVMWallet()
	if err != nil {
		fmt.Printf("%sError Generating Wallet:%s %v\n", Red, Reset, err)

	} else {

		fmt.Printf("\n%sGenerated Wallet:%s\n", Bold+Blue, Reset)
		fmt.Printf("%sPrivate Key:%s %s\n", Cyan, Reset, privateKey)
		fmt.Printf("%sPublic Key:%s %s\n", Cyan, Reset, publicKey)
		fmt.Printf("%sAddress:%s %s\n", Cyan, Reset, address)
	}

	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln() // Attend une entrée utilisateur
}
func generateEVMWallet() (string, string, string, error) {
	// Generate private key
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to generate private key: %w", err)
	}

	// Convert private key to hex string
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)

	// Derive public key from private key
	publicKey := privateKey.PublicKey
	publicKeyBytes := crypto.FromECDSAPub(&publicKey)
	publicKeyHex := fmt.Sprintf("%x", publicKeyBytes)

	// Generate wallet address
	address := crypto.PubkeyToAddress(publicKey).Hex()

	return privateKeyHex, publicKeyHex, address, nil
}
