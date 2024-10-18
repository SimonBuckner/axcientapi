package axcient

type OrganisationQuery struct {
	ac *Axcient
}

func newOrganisationQuery(ac *Axcient) *OrganisationQuery {
	return &OrganisationQuery{
		ac: ac,
	}
}

// func (q *OrganisationQuery) Build() (*OrganisationQuery, error) {

// 	req, err := q.ac.NewGetQuery("organization")
// 	if err != nil {
// 		return nil, err
// 	}

// 	query := req.URL.Query()

// 	req.URL.RawQuery = query.Encode()
// 	q
// 	return req, nil
// }
