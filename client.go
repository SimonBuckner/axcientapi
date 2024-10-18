package axcient

import (
	"fmt"
	"net/http"
)

type ClientsQuery struct {
	*axcientQuery

	clientId          *int64
	includeAppliances *bool
}

func newClientsQuery(ac *Axcient) *ClientsQuery {
	return &ClientsQuery{
		axcientQuery: newAxcientQuery(ac),
	}
}

func (q *ClientsQuery) SetIncludeAppliances(state bool) *ClientsQuery {
	q.includeAppliances = &state
	return q
}

func (q *ClientsQuery) SelectByClientId(clientId int64) *ClientsQuery {
	q.clientId = &clientId
	return q
}

func (q *ClientsQuery) Build() (*ClientsQuery, error) {

	endpoint := "client"
	if q.clientId != nil {
		endpoint = fmt.Sprintf("client/%d", *q.clientId)
	}

	if err := q.NewGetQuery(endpoint); err != nil {
		return nil, err
	}

	if q.includeAppliances != nil {
		if *q.includeAppliances {
			q.AddUrlQuery("include_appliances", "true")
		} else {
			q.AddUrlQuery("include_appliances", "false")
		}
	}

	return q, nil
}

func (q *ClientsQuery) GetAll() ([]OrgLevelClient, error) {
	if q.request == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}

	res, err := q.ac.ApiHelper.Call(q.request)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, err
	}

	var out []OrgLevelClient

	if err := q.ac.DecodeJsonBody(res.Body, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (q *ClientsQuery) GetSingle() (*OrgLevelClient, error) {
	if q.request == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}

	res, err := q.ac.ApiHelper.Call(q.request)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, err
	}

	var out OrgLevelClient

	if err := q.ac.DecodeJsonBody(res.Body, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
