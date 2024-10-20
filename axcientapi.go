package axcient

import "github.com/simonbuckner/goquadac"

type AxcientApi struct {
	*goquadac.ApiHelper
}

func NewAxcientApi(baseUrl string) *AxcientApi {
	Api := goquadac.NewApiHelper(baseUrl).
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

func (api *AxcientApi) GetDeviceAutoverifyDetails() *DeviceAutoverifyQuery {
	return newDeviceAutoverifyQuery(api)
}

func (api *AxcientApi) GetDeviceRestorePoints() *DeviceRestorePointQuery {
	return newDeviceRestorePointQuery(api)
}

// Client Queries
func (api *AxcientApi) GetClients() *ClientQuery {
	return newClientQuery(api)
}

func (api *AxcientApi) GetClientDevices() *ClientDeviceQuery {
	return newClientDeviceQuery(api)
}

// Job Queries
func (api *AxcientApi) GetJobs() *JobQuery {
	return newJobQuery(api)
}

func (api *AxcientApi) GetJobHistory() *JobHistoryQuery {
	return newJobHistoryQuery(api)
}

// Organisation Queries
func (api *AxcientApi) GetOrganisation() *OrganisationQuery {
	return newOrganisationQuery(api)
}

// Appliance Queries
func (api *AxcientApi) GetAppliance() *ApplianceQuery {
	return newApplianceQuery(api)
}

// D2C Agent Queries
func (api *AxcientApi) GetD2CAgent() *D2CAgentQuery {
	return newD2CAgentQuery(api)
}
