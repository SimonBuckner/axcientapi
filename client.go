package axcientapi

import (
	"fmt"

	"github.com/simonbuckner/goquadac"
)

type ClientQuery struct {
	api   *AxcientApi
	query *goquadac.ApiQuery

	// URL Path fields
	clientId *int64

	// URL query fields
	includeAppliances *bool
}

func newClientQuery(api *AxcientApi) *ClientQuery {
	return &ClientQuery{
		api: api,
	}
}
func (q *ClientQuery) SelectByClientId(clientId int64) *ClientQuery {
	q.clientId = &clientId
	return q
}

func (q *ClientQuery) SetIncludeAppliances(state bool) *ClientQuery {
	q.includeAppliances = &state
	return q
}

func (q *ClientQuery) Build() (*ClientQuery, error) {

	endpoint := "client"
	if q.clientId != nil {
		endpoint = fmt.Sprintf("client/%d", *q.clientId)
	}

	query := q.api.NewGetQuery(endpoint)

	if q.includeAppliances != nil {
		query.AddUrlQuery("include_appliances", goquadac.BooltoString(*q.includeAppliances))
	}

	q.query = query
	return q, nil
}
func (q *ClientQuery) GetAll() ([]OrgLevelClient, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out []OrgLevelClient
	err := q.query.Get(&out)
	return out, err
}

func (q *ClientQuery) GetSingle() (*OrgLevelClient, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out OrgLevelClient
	err := q.query.Get(&out)
	return &out, err
}

type ClientDeviceQuery struct {
	api   *AxcientApi
	query *goquadac.ApiQuery

	// URL Path fields
	clientId *int64

	// URL query fields
	serviceId    *string
	direct2Cloud *bool
}

func newClientDeviceQuery(api *AxcientApi) *ClientDeviceQuery {
	return &ClientDeviceQuery{
		api: api,
	}
}

func (q *ClientDeviceQuery) SelectByClientId(clientId int64) *ClientDeviceQuery {
	q.clientId = &clientId
	return q
}

func (q *ClientDeviceQuery) SetServiceId(serviceId string) *ClientDeviceQuery {
	q.serviceId = &serviceId
	return q
}
func (q *ClientDeviceQuery) SetDirect2Cloud(state bool) *ClientDeviceQuery {
	q.direct2Cloud = &state
	return q
}

func (q *ClientDeviceQuery) Build() (*ClientDeviceQuery, error) {

	if q.clientId == nil {
		return nil, fmt.Errorf("client_id required")
	}
	endpoint := fmt.Sprintf("client/%d/device", *q.clientId)

	query := q.api.NewGetQuery(endpoint)

	if q.serviceId != nil {
		query.AddUrlQuery("service_id", *q.serviceId)
	}
	if q.direct2Cloud != nil {
		query.AddUrlQuery("d2c_only", goquadac.BooltoString(*q.direct2Cloud))
	}

	q.query = query
	return q, nil
}

func (q *ClientDeviceQuery) GetAll() ([]OrgLevelDevice, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out []OrgLevelDevice
	err := q.query.Get(&out)
	return out, err
}
