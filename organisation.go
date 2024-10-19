package axcient

import "github.com/simonbuckner/axcient/apihelper"

type OrganisationQuery struct {
	api   *AxcientApi
	query *apihelper.ApiQuery

	// URL query fields

	// URL query fields
}

func newOrganisationQuery(api *AxcientApi) *OrganisationQuery {
	return &OrganisationQuery{
		api: api,
	}
}

func (q *OrganisationQuery) Build() (*OrganisationQuery, error) {

	query := q.api.NewGetQuery("organization").
		SetDumpRequest(true).
		SetDumpResponse(false).
		SetDumpResponseBody(false)

	q.query = query
	return q, nil
}

func (q *OrganisationQuery) get() (*apihelper.ApiQuery, error) {
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
func (q *OrganisationQuery) GetSingle() (*OrgLevelOrganisation, error) {

	query, err := q.get()
	if err != nil {
		return nil, err
	}
	var out OrgLevelOrganisation

	if err := query.DecodeJsonBody(&out); err != nil {
		return nil, err
	}

	return &out, nil
}
