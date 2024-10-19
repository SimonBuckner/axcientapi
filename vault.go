package axcient

import (
	"fmt"

	"github.com/simonbuckner/axcient/apihelper"
)

type VaultsQuery struct {
	api   *AxcientApi
	query *apihelper.ApiQuery

	// URL Path fields
	vaultId *int64

	// URL query fields
	vaultType      *string // Private / Cloud
	active         *bool
	withUrl        *bool
	limit          *int64
	includeDevices *bool
}

func newVaultsQuery(api *AxcientApi) *VaultsQuery {
	return &VaultsQuery{
		api: api,
	}
}
func (q *VaultsQuery) SelectByVaultId(vaultId int64) *VaultsQuery {
	q.vaultId = &vaultId
	return q
}

func (q *VaultsQuery) SetPublicVault() *VaultsQuery {
	vt := "Public"
	q.vaultType = &vt
	return q
}

func (q *VaultsQuery) SetCloudVault() *VaultsQuery {
	vt := "Cloud"
	q.vaultType = &vt
	return q
}

func (q *VaultsQuery) SetActiveState(state bool) *VaultsQuery {
	q.active = &state
	return q
}

func (q *VaultsQuery) SetWithUrlPresence(state bool) *VaultsQuery {
	q.withUrl = &state
	return q
}

func (q *VaultsQuery) SetVaultLimit(limit int64) *VaultsQuery {
	q.limit = &limit
	return q
}

func (q *VaultsQuery) SetIncludeDevices(state bool) *VaultsQuery {
	q.includeDevices = &state
	return q
}

func (q *VaultsQuery) Build() (*VaultsQuery, error) {

	endpoint := "vault"
	if q.vaultId != nil {
		endpoint = fmt.Sprintf("vault/%d", *q.vaultId)
	}

	query := q.api.NewGetQuery(endpoint).
		SetDumpRequest(false).
		SetDumpResponse(false)

	if q.vaultType != nil {
		query.AddUrlQuery("vault_type", *q.vaultType)
	}
	if q.active != nil {
		query.AddUrlQuery("active", apihelper.BooltoString(*q.active))
	}
	if q.withUrl != nil {
		query.AddUrlQuery("with_url", apihelper.BooltoString(*q.withUrl))
	}
	if q.limit != nil {
		query.AddUrlQuery("limit", fmt.Sprintf("%d", *q.limit))
	}
	if q.includeDevices != nil {
		query.AddUrlQuery("include_devices", apihelper.BooltoString(*q.includeDevices))
	}

	q.query = query
	return q, nil
}

func (q *VaultsQuery) get() (*apihelper.ApiQuery, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}

	query, err := q.query.Call()
	if !query.ResponsOK() {
		return nil, err
	}
	return query, nil
}

func (q *VaultsQuery) GetAll() ([]OrgLevelVault, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out []OrgLevelVault
	err := q.query.Get(&out)
	return out, err
}

func (q *VaultsQuery) GetSingle() (*OrgLevelVault, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out OrgLevelVault
	err := q.query.Get(&out)
	return &out, err
}

type VaultThresholdConnectivityQuery struct {
	api   *AxcientApi
	query *apihelper.ApiQuery

	// URL Path fields
	vaultId *int64
}

func newVaultThresholdConnectivityQuery(api *AxcientApi, vaultId int64) *VaultThresholdConnectivityQuery {
	return &VaultThresholdConnectivityQuery{
		api:     api,
		vaultId: &vaultId,
	}
}

func (q *VaultThresholdConnectivityQuery) Build() (*VaultThresholdConnectivityQuery, error) {

	if q.vaultId == nil {
		return nil, fmt.Errorf("vault_id required to query thresholds")
	}
	endpoint := fmt.Sprintf("vault/%d/threshold/connectivity", *q.vaultId)

	query := q.api.NewGetQuery(endpoint).
		SetDumpRequest(false).
		SetDumpResponse(true)

	q.query = query
	return q, nil
}

func (q *VaultThresholdConnectivityQuery) GetSingle() (*VaultThresholdBody, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out VaultThresholdBody
	err := q.query.Get(&out)
	return &out, err
}
