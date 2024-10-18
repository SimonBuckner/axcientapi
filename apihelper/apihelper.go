package apihelper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/gookit/goutil/dump"
)

type ApiHelper struct {
	baseUrl      string
	headers      http.Header
	dumpRequests bool
}

func NewApiHelper(baseUrl string, dumpRequests bool) *ApiHelper {
	return &ApiHelper{
		baseUrl:      baseUrl,
		headers:      make(http.Header),
		dumpRequests: dumpRequests,
	}
}

func (api *ApiHelper) SetAuthHeader(key, value string) {
	api.headers.Set(key, value)
}

func (api *ApiHelper) SetDefaultHeader(key, value string) {
	api.headers.Set(key, value)
}

func (api *ApiHelper) NewGetQuery(endpoint string) (*http.Request, error) {

	url := api.baseUrl + endpoint
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = api.headers
	return req, nil
}

func (api *ApiHelper) Call(req *http.Request) (*http.Response, error) {
	if api.dumpRequests {
		dump, err := httputil.DumpRequestOut(req, false)
		if err != nil {
			return nil, err
		}
		fmt.Printf("\nDump request:\n %q\n", dump)
	}
	return http.DefaultClient.Do(req)
}

func (api *ApiHelper) DecodeJsonBody(body io.Reader, out any) error {

	content, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, out)
}

func (api *ApiHelper) DumpRequest(req *http.Request) {
	out, err := httputil.DumpRequestOut(req, false)
	if err != nil {
		panic(err)
	}
	dump.Println(out)
	fmt.Println()
}

func (api *ApiHelper) DumpRespone(res *http.Response) {
	out, err := httputil.DumpResponse(res, false)
	if err != nil {
		panic(err)
	}
	dump.Println(out)
	fmt.Println()
}
