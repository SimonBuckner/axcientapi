package axcient

import "net/http"

type axcientQuery struct {
	ac      *Axcient
	request *http.Request
}

func newAxcientQuery(ac *Axcient) *axcientQuery {
	return &axcientQuery{
		ac:      ac,
		request: nil,
	}
}

func (aq *axcientQuery) NewGetQuery(endpoint string) error {
	req, err := aq.ac.NewGetQuery(endpoint)
	if err != nil {
		return err
	}
	aq.request = req
	return nil
}

func (aq *axcientQuery) AddUrlQuery(key, value string) {
	query := aq.request.URL.Query()
	query.Add(key, value)
	aq.request.URL.RawQuery = query.Encode()
}
