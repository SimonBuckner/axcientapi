package axcient

import (
	"fmt"

	"github.com/simonbuckner/axcient/apihelper"
)

type D2CAgentQuery struct {
	api   *AxcientApi
	query *apihelper.ApiQuery

	// URL Path fields
	clientId *int64
	vaultId  *int64

	// URL query fields
}

func newD2CAgentQuery(api *AxcientApi) *D2CAgentQuery {
	return &D2CAgentQuery{
		api: api,
	}
}
func (q *D2CAgentQuery) SelectByClientId(clientId int64) *D2CAgentQuery {
	q.clientId = &clientId
	return q
}

func (q *D2CAgentQuery) SelectByVaultId(vaultId int64) *D2CAgentQuery {
	q.vaultId = &vaultId
	return q
}

func (q *D2CAgentQuery) Build() (*D2CAgentQuery, error) {

	if q.clientId == nil {
		return nil, fmt.Errorf("client_id required")
	}

	if q.vaultId == nil {
		return nil, fmt.Errorf("vault_id required")
	}

	endpoint := fmt.Sprintf("client/%d/vault/%d/d2c_agent", *q.clientId, *q.vaultId)

	query := q.api.NewPostQuery(endpoint).
		SetDumpRequest(true).
		SetDumpResponse(true).
		SetDumpResponseBody(true)

	q.query = query
	return q, nil
}

func (q *D2CAgentQuery) get() (*apihelper.ApiQuery, error) {
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
func (q *D2CAgentQuery) GetSingle() (*ClientLevelDirect2CloudAgentToken, error) {
	query, err := q.get()
	if err != nil {
		return nil, err
	}

	var out ClientLevelDirect2CloudAgentToken

	if err := query.DecodeJsonBody(&out); err != nil {
		return nil, err
	}

	return &out, nil
}
