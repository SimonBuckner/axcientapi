package axcient

import (
	"fmt"
	"net/http"
)

type VaultsQuery struct {
	*axcientQuery
	vaultId        *int64
	vaultType      *string // Private / Cloud
	active         *bool
	withUrl        *bool
	limit          *int64
	includeDevices *bool
}

func newVaultsQuery(ac *Axcient) *VaultsQuery {
	return &VaultsQuery{
		axcientQuery: newAxcientQuery(ac),
	}
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

func (q *VaultsQuery) SelectByVaultId(vaultId int64) *VaultsQuery {
	q.vaultId = &vaultId
	return q
}

func (q *VaultsQuery) Build() (*VaultsQuery, error) {

	endpoint := "vault"
	if q.vaultId != nil {
		endpoint = fmt.Sprintf("vault/%d", *q.vaultId)
	}

	if err := q.NewGetQuery(endpoint); err != nil {
		return nil, err
	}

	if q.vaultType != nil {
		q.AddUrlQuery("vault_type", *q.vaultType)
	}
	if q.active != nil {
		if *q.active {
			q.AddUrlQuery("active", "true")
		} else {
			q.AddUrlQuery("active", "false")
		}
	}
	if q.withUrl != nil {
		if *q.withUrl {
			q.AddUrlQuery("with_url", "true")
		} else {
			q.AddUrlQuery("with_url", "false")
		}
	}
	if q.limit != nil {
		q.AddUrlQuery("limit", fmt.Sprintf("%d", *q.limit))
	}
	if q.includeDevices != nil {
		if *q.includeDevices {
			q.AddUrlQuery("include_devices", "true")
		} else {
			q.AddUrlQuery("include_devices", "false")
		}
	}

	return q, nil
}

func (q *VaultsQuery) GetAll() ([]OrgLevelVault, error) {
	if q.request == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}

	res, err := q.ac.ApiHelper.Call(q.request)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, err
	}

	var out []OrgLevelVault

	if err := q.ac.DecodeJsonBody(res.Body, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (q *VaultsQuery) GetSingle() (*OrgLevelVault, error) {
	if q.request == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}

	res, err := q.ac.ApiHelper.Call(q.request)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, err
	}

	var out OrgLevelVault

	if err := q.ac.DecodeJsonBody(res.Body, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

type VaultThresholdConnectivityQuery struct {
	*axcientQuery
	vaultId *int64
}

func newVaultThresholdConnectivityQuery(ac *Axcient, vaultId int64) *VaultThresholdConnectivityQuery {
	return &VaultThresholdConnectivityQuery{
		axcientQuery: newAxcientQuery(ac),
		vaultId:      &vaultId,
	}
}

func (q *VaultThresholdConnectivityQuery) Build() (*VaultThresholdConnectivityQuery, error) {

	if q.vaultId == nil {
		return nil, fmt.Errorf("vault_id required to query thresholds")
	}
	endpoint := fmt.Sprintf("vault/%d/threshold/connectivity", *q.vaultId)

	if err := q.NewGetQuery(endpoint); err != nil {
		return nil, err
	}

	return q, nil
}

func (q *VaultThresholdConnectivityQuery) GetSingle() (*VaultThresholdBody, error) {
	if q.request == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}

	res, err := q.ac.ApiHelper.Call(q.request)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, err
	}

	var out VaultThresholdBody

	if err := q.ac.DecodeJsonBody(res.Body, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
