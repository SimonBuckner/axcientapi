package axcient

import (
	"fmt"

	"github.com/simonbuckner/axcient/apihelper"
)

type ApplianceQuery struct {
	api   *AxcientApi
	query *apihelper.ApiQuery

	// URL Path fields
	applianceId *int64
	clientId    *int64

	// URL query fields
	serviceId      *string
	includeDevices *bool
}

func newApplianceQuery(api *AxcientApi) *ApplianceQuery {
	return &ApplianceQuery{
		api: api,
	}
}
func (q *ApplianceQuery) SelectByClientId(clientId int64) *ApplianceQuery {
	q.clientId = &clientId
	return q
}

func (q *ApplianceQuery) SetIncludeDevices(state bool) *ApplianceQuery {
	q.includeDevices = &state
	return q
}

func (q *ApplianceQuery) SetServiceId(serviceId string) *ApplianceQuery {
	q.serviceId = &serviceId
	return q
}

func (q *ApplianceQuery) Build() (*ApplianceQuery, error) {

	endpoint := "appliance"
	if q.clientId != nil && q.applianceId != nil {
		return nil, fmt.Errorf("supply applicance_id or client_id but not both")
	}
	if q.applianceId != nil {
		endpoint = fmt.Sprintf("appliance/%d", *q.applianceId)
	}
	if q.clientId != nil {
		endpoint = fmt.Sprintf("client/%d/appliance", *q.clientId)
	}

	query := q.api.NewGetQuery(endpoint).
		SetDumpRequest(true).
		SetDumpResponse(true).
		SetDumpResponseBody(true)

	if q.includeDevices != nil {
		query.AddUrlQuery("include_devices", apihelper.BooltoString(*q.includeDevices))
	}
	if q.serviceId != nil {
		query.AddUrlQuery("service_id", *q.serviceId)
	}

	q.query = query
	return q, nil
}

func (q *ApplianceQuery) get() (*apihelper.ApiQuery, error) {
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
func (q *ApplianceQuery) GetAll() ([]OrgLevelAppliance, error) {
	query, err := q.get()
	if err != nil {
		return nil, err
	}

	var out []OrgLevelAppliance

	if err := query.DecodeJsonBody(&out); err != nil {
		return nil, err
	}

	return out, nil
}

func (q *ApplianceQuery) GetSingle() (*OrgLevelAppliance, error) {
	query, err := q.get()
	if err != nil {
		return nil, err
	}

	var out OrgLevelAppliance

	if err := query.DecodeJsonBody(&out); err != nil {
		return nil, err
	}

	return &out, nil
}
