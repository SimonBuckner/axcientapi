package axcient

import (
	"github.com/simonbuckner/backupcheck/backupcheck/apihelper"
)

const (
	BASE_URL = "https://axapi.axcient.com/x360recover/"
)

type Axcient struct {
	*apihelper.ApiHelper
}

func NewAxcientApi(dumpRequests bool) *Axcient {
	Api := apihelper.NewApiHelper(BASE_URL, dumpRequests)
	Api.SetDefaultHeader("accept", "application/json")

	return &Axcient{
		ApiHelper: Api,
	}
}

func (ac *Axcient) Authenticate(apiKey string) error {
	ac.SetAuthHeader("X-API-Key", apiKey)
	return nil
}

// Vault Queries

func (ac *Axcient) GetVaults() *VaultsQuery {
	return newVaultsQuery(ac)
}

func (ac *Axcient) GetVaultVaultThresholdConnectivity(vaultId int64) *VaultThresholdConnectivityQuery {
	return newVaultThresholdConnectivityQuery(ac, vaultId)
}

// func (ac *Axcient) NewVaultQuery(vaultId int64) *VaultQuery {
// 	return newVaultQuery(ac, vaultId)
// }

// Device Queries

// func (ac *Axcient) NewOrganisationQuery() *OrganisationQuery {
// 	return newOrganisationQuery(ac)
// }

// Job Queries

// Client Queries

func (ac *Axcient) GetClients() *ClientsQuery {
	return newClientsQuery(ac)
}

// Appliance Queries

//Organisation Queries

// D21C Agent Queries
