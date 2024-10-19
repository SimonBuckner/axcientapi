package axcient

import "github.com/simonbuckner/axcient/apihelper"

type AxcientApi struct {
	*apihelper.ApiHelper
}

func NewAxcientApi(baseUrl string, dumpRequest, dumpResponse bool) *AxcientApi {
	Api := apihelper.NewApiHelper(baseUrl).
		SetDefaultHeader("accept", "application/json")

	return &AxcientApi{
		ApiHelper: Api,
	}
}

func (api *AxcientApi) Authenticate(apiKey string) error {
	api.SetAuthHeader("X-API-Key", apiKey)
	return nil
}

// Vault Queries

func (api *AxcientApi) GetVaults() *VaultsQuery {
	return newVaultsQuery(api)
}

func (api *AxcientApi) GetVaultVaultThresholdConnectivity(vaultId int64) *VaultThresholdConnectivityQuery {
	return newVaultThresholdConnectivityQuery(api, vaultId)
}

// Device Queries
func (api *AxcientApi) GetDevices() *DeviceQuery {
	return newDeviceQuery(api)
}

// func (api *AxcientApiewOrganisationQuery() *OrganisationQuery {
// 	return newOrganisationQuery(api)
// }

// Job Queries

// Client Queries

func (api *AxcientApi) GetClients() *ClientQuery {
	return newClientQuery(api)
}

func (api *AxcientApi) GetClientDevices() *ClientDeviceQuery {
	return newClientDeviceQuery(api)
}

func (api *AxcientApi) GetDeviceAutoverifyDetails() *DeviceAutoverifyQuery {
	return newDeviceAutoverifyQuery(api)
}

func (api *AxcientApi) GetDeviceRestorePoints() *DeviceRestorePointQuery {
	return newDeviceRestorePointQuery(api)
}

// Appliance Queries

//Organisation Queries

// D21C Agent Queries
