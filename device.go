package axcient

import (
	"fmt"

	"github.com/simonbuckner/axcient/apihelper"
)

type DeviceQuery struct {
	api   *AxcientApi
	query *apihelper.ApiQuery

	// URL query fields
	deviceId *int64

	// URL query fields
	limit  *int64
	offset *int64
}

func newDeviceQuery(api *AxcientApi) *DeviceQuery {
	return &DeviceQuery{
		api: api,
	}
}

func (q *DeviceQuery) SelectByDeviceId(deviceId int64) *DeviceQuery {
	q.deviceId = &deviceId
	return q
}

func (q *DeviceQuery) SetLimit(limit int64) *DeviceQuery {
	q.limit = &limit
	return q
}
func (q *DeviceQuery) SetOffset(offset int64) *DeviceQuery {
	q.offset = &offset
	return q
}

func (q *DeviceQuery) Build() (*DeviceQuery, error) {

	endpoint := "device"
	if q.deviceId != nil {
		endpoint = fmt.Sprintf("device/%d", *q.deviceId)
	}

	query := q.api.NewGetQuery(endpoint).
		SetDumpRequest(true).
		SetDumpResponse(false).
		SetDumpResponseBody(false)

	if q.limit != nil {
		query.AddUrlQuery("limit", apihelper.I64toString(*q.limit))
	}
	if q.offset != nil {
		query.AddUrlQuery("offset", apihelper.I64toString(*q.offset))
	}

	q.query = query
	return q, nil
}

func (q *DeviceQuery) get() (*apihelper.ApiQuery, error) {
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
func (q *DeviceQuery) GetAll() ([]OrgLevelDevice, error) {

	query, err := q.get()
	if err != nil {
		return nil, err
	}
	var out []OrgLevelDevice

	if err := query.DecodeJsonBody(&out); err != nil {
		return nil, err
	}

	return out, nil
}

func (q *DeviceQuery) GetSingle() (*OrgLevelDevice, error) {

	query, err := q.get()
	if err != nil {
		return nil, err
	}
	var out OrgLevelDevice

	if err := query.DecodeJsonBody(&out); err != nil {
		return nil, err
	}

	return &out, nil
}

type DeviceAutoverifyQuery struct {
	api   *AxcientApi
	query *apihelper.ApiQuery

	// URL query fields
	deviceId *int64

	// URL query fields
}

func newDeviceAutoverifyQuery(api *AxcientApi) *DeviceAutoverifyQuery {
	return &DeviceAutoverifyQuery{
		api: api,
	}
}

func (q *DeviceAutoverifyQuery) SelectByDeviceId(deviceId int64) *DeviceAutoverifyQuery {
	q.deviceId = &deviceId
	return q
}

func (q *DeviceAutoverifyQuery) Build() (*DeviceAutoverifyQuery, error) {

	endpoint := "device"
	if q.deviceId != nil {
		endpoint = fmt.Sprintf("device/%d/autoverify", *q.deviceId)
	}

	query := q.api.NewGetQuery(endpoint).
		SetDumpRequest(true).
		SetDumpResponse(false).
		SetDumpResponseBody(false)

	q.query = query
	return q, nil
}

func (q *DeviceAutoverifyQuery) get() (*apihelper.ApiQuery, error) {
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

func (q *DeviceAutoverifyQuery) GetAll() ([]OrgLevelMachineAutoverifyDetails, error) {

	query, err := q.get()
	if err != nil {
		return nil, err
	}
	var out []OrgLevelMachineAutoverifyDetails

	if err := query.DecodeJsonBody(&out); err != nil {
		return nil, err
	}

	return out, nil
}

type DeviceRestorePointQuery struct {
	api   *AxcientApi
	query *apihelper.ApiQuery

	// URL query fields
	deviceId *int64

	// URL query fields
}

func newDeviceRestorePointQuery(api *AxcientApi) *DeviceRestorePointQuery {
	return &DeviceRestorePointQuery{
		api: api,
	}
}

func (q *DeviceRestorePointQuery) SelectByDeviceId(deviceId int64) *DeviceRestorePointQuery {
	q.deviceId = &deviceId
	return q
}

func (q *DeviceRestorePointQuery) Build() (*DeviceRestorePointQuery, error) {

	endpoint := "device"
	if q.deviceId != nil {
		endpoint = fmt.Sprintf("device/%d/restore_point", *q.deviceId)
	}

	query := q.api.NewGetQuery(endpoint).
		SetDumpRequest(true).
		SetDumpResponse(false).
		SetDumpResponseBody(false)

	q.query = query
	return q, nil
}

func (q *DeviceRestorePointQuery) get() (*apihelper.ApiQuery, error) {
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

func (q *DeviceRestorePointQuery) GetAll() ([]MachineRestorePoint, error) {

	query, err := q.get()
	if err != nil {
		return nil, err
	}
	var out []MachineRestorePoint

	if err := query.DecodeJsonBody(&out); err != nil {
		return nil, err
	}

	return out, nil
}
