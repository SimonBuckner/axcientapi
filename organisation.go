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

func (q *OrganisationQuery) GetSingle() (*OrgLevelOrganisation, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out OrgLevelOrganisation
	err := q.query.Get(&out)
	return &out, err
}
