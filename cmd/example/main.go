package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gookit/goutil/dump"
	"github.com/joho/godotenv"
	"github.com/simonbuckner/backupcheck/backupcheck/axcientclient"
)

func main() {

	// Load .env files
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	apiKey := os.Getenv("API_KEY")
	clientId := AtoI64(os.Getenv("CLIENT_ID"))
	vaultId := AtoI64(os.Getenv("VAULT_ID"))

	axcient := axcientclient.NewAxcientApi(true)
	axcient.Authenticate(apiKey)

	printBanner("Get All Vaults")

	allVaults, err := axcient.GetVaults().SetIncludeDevices(true).GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(allVaults)

	printBanner("Get Specific Vault")

	singleVault, err := axcient.GetVaults().
		SelectByVaultId(vaultId).
		GetSingle()
	if err != nil {
		panic(err)
	}
	dump.Print(singleVault)

	printBanner("Get Specific Vault Threshold Connectivity")

	singleVaultThreshold, err := axcient.GetVaultVaultThresholdConnectivity(vaultId).
		GetSingle()
	if err != nil {
		panic(err)
	}
	dump.Print(singleVaultThreshold)

	printBanner("Get All Clients")

	allClients, err := axcient.GetClients().SetIncludeAppliances(true).GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(allClients)

	printBanner("Get Specific Client")

	singlClient, err := axcient.GetClients().
		SetIncludeAppliances(true).
		SelectByClientId(clientId).
		GetSingle()
	if err != nil {
		panic(err)
	}
	dump.Print(singlClient)

}

func AtoI64(val string) int64 {
	i64, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(err)
	}
	return i64
}

func printBanner(message string) {
	fmt.Println()
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("    " + message)
	fmt.Println("-------------------------------------------------------------")
	fmt.Println()
}

func printSubBanner(message string) {
	fmt.Println("    " + message)
	fmt.Println("    ---------------------------------------------------------")
	fmt.Println()
}

func printStatus(status string) {
	fmt.Println("    Status: " + status)
}
