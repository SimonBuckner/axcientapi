package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gookit/goutil/dump"
	"github.com/joho/godotenv"
	"github.com/simonbuckner/axcient"
	"github.com/simonbuckner/goquadac"
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
	vaultId := goquadac.StringtoI64(os.Getenv("VAULT_ID"))
	clientId := goquadac.StringtoI64(os.Getenv("CLIENT_ID"))
	deviceId := goquadac.StringtoI64(os.Getenv("DEVICE_ID"))
	jobId := goquadac.StringtoI64(os.Getenv("JOB_ID"))

	axcient := axcient.NewAxcientApi(baseUrl, true, true)
	axcient.Authenticate(apiKey)

	// -------------------------------------------------------------------------
	//   Endpoints starting /vault
	// -------------------------------------------------------------------------

	printBanner("GET:/vault")
	allVaults, err := axcient.GetVaults().SetIncludeDevices(true).GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(allVaults)

	printBanner("GET:/vault/{vault_id}")
	singleVault, err := axcient.GetVaults().
		SelectByVaultId(vaultId).
		GetSingle()
	if err != nil {
		panic(err)
	}
	dump.Print(singleVault)

	printBanner("GET:/vault/{vault_id}/threshold/connectivity")
	singleVaultThreshold, err := axcient.GetVaultVaultThresholdConnectivity(vaultId).
		GetSingle()
	if err != nil {
		panic(err)
	}
	dump.Print(singleVaultThreshold)

	// -------------------------------------------------------------------------
	//   Endpoints starting /device
	// -------------------------------------------------------------------------

	printBanner("GET:/device")
	devices, err := axcient.GetDevices().GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(devices)

	printBanner("GET:/device/{device_id}")
	device, err := axcient.GetDevices().SelectByDeviceId(deviceId).GetSingle()
	if err != nil {
		panic(err)
	}
	dump.Print(device)

	printBanner("GET:/device/{device_id}/autoverify")
	autoverify, err := axcient.GetDeviceAutoverifyDetails().SelectByDeviceId(deviceId).GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(autoverify)

	printBanner("GET:/device/{device_id}/restore_point")
	restorePoints, err := axcient.GetDeviceRestorePoints().SelectByDeviceId(deviceId).GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(restorePoints)

	// -------------------------------------------------------------------------
	//   Endpoints starting /client
	// -------------------------------------------------------------------------

	printBanner("GET:/client")
	allClients, err := axcient.GetClients().SetIncludeAppliances(true).GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(allClients)

	printBanner("GET:/client/{client_id}")
	singleClient, err := axcient.GetClients().
		SetIncludeAppliances(true).
		SelectByClientId(clientId).
		GetSingle()
	if err != nil {
		panic(err)
	}
	dump.Print(singleClient)

	printBanner("GET:/client/{client_id}/device")
	clientDevices, err := axcient.GetClientDevices().
		SelectByClientId(clientId).
		SetDirect2Cloud(true).
		GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(clientDevices)

	printBanner("GET:/client/{client_id}/device/{device_id}/job")
	deviceJobs, err := axcient.GetJobs().
		SelectByClientId(clientId).
		SelectByDeviceId(deviceId).
		GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(deviceJobs)

	printBanner("GET:/client/{client_id}/device/{device_id}/job/{job_id}")
	deviceJob, err := axcient.GetJobs().
		SelectByClientId(clientId).
		SelectByDeviceId(deviceId).
		SelectByJobId(jobId).
		GetSingle()
	if err != nil {
		panic(err)
	}
	dump.Print(deviceJob)

	printBanner("GET:/client/{client_id}/device/{device_id}/job/{job_id}/history")
	jobHistory, err := axcient.GetJobHistory().
		SelectByClientId(clientId).
		SelectByDeviceId(deviceId).
		SelectByJobId(jobId).
		GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(jobHistory)

	// -------------------------------------------------------------------------
	//   Endpoints starting /organization
	// -------------------------------------------------------------------------

	printBanner("Get Organisation Details")
	org, err := axcient.GetOrganisation().GetSingle()
	if err != nil {
		panic(err)
	}
	dump.Print(org)

	// -------------------------------------------------------------------------
	//   Endpoints starting /organization
	// -------------------------------------------------------------------------

	printBanner("POST:/client/{client_id}/vault/{vault_id}/d2c_agent")
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
