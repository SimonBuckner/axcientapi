package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gookit/goutil/dump"
	"github.com/joho/godotenv"
	"github.com/simonbuckner/axcient"
	"github.com/simonbuckner/axcient/apihelper"
)

func main() {

	// Load .env files
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	vaultId := apihelper.StringtoI64(os.Getenv("VAULT_ID"))
	clientId := apihelper.StringtoI64(os.Getenv("CLIENT_ID"))
	deviceId := apihelper.StringtoI64(os.Getenv("DEVICE_ID"))

	axcient := axcient.NewAxcientApi(baseUrl, true, true)
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

	singleClient, err := axcient.GetClients().
		SetIncludeAppliances(true).
		SelectByClientId(clientId).
		GetSingle()
	if err != nil {
		panic(err)
	}
	dump.Print(singleClient)

	printBanner("Get Specific Client Devices")

	clientDevices, err := axcient.GetClientDevices().
		SelectByClientId(clientId).
		SetDirect2Cloud(true).
		GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(clientDevices)

	printBanner("Get All Devices")

	devices, err := axcient.GetDevices().GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(devices)

	printBanner("Get Single Device")

	device, err := axcient.GetDevices().SelectByDeviceId(deviceId).GetSingle()
	if err != nil {
		panic(err)
	}
	dump.Print(device)

	printBanner("Get Single Device Autoverify Details")

	autoverify, err := axcient.GetDeviceAutoverifyDetails().SelectByDeviceId(deviceId).GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(autoverify)

	printBanner("Get Single Device Restore Points")

	restorePoints, err := axcient.GetDeviceRestorePoints().SelectByDeviceId(deviceId).GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(restorePoints)

	printBanner("Get Organisation Details")

	org, err := axcient.GetOrganisation().GetSingle()
	if err != nil {
		panic(err)
	}
	dump.Print(org)

	printBanner("Get Organisation Details")

	d2c, err := axcient.GetD2CAgent().
		SelectByClientId(clientId).
		SelectByVaultId(vaultId).
		GetSingle()

	if err != nil {
		panic(err)
	}
	dump.Print(d2c)

}

func printBanner(message string) {
	fmt.Println()
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("    " + message)
	fmt.Println("-------------------------------------------------------------")
	fmt.Println()
}

// func printSubBanner(message string) {
// 	fmt.Println("    " + message)
// 	fmt.Println("    ---------------------------------------------------------")
// 	fmt.Println()
// }

// func printStatus(status string) {
// 	fmt.Println("    Status: " + status)
// }
