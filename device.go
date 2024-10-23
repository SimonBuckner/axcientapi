package axcientapi

import (
	"fmt"

	"github.com/simonbuckner/goquadac"
)

type DeviceQuery struct {
	api   *AxcientApi
	query *goquadac.ApiQuery

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

	query := q.api.NewGetQuery(endpoint)

	if q.limit != nil {
		query.AddUrlQuery("limit", goquadac.I64toString(*q.limit))
	}
	if q.offset != nil {
		query.AddUrlQuery("offset", goquadac.I64toString(*q.offset))
	}

	q.query = query
	return q, nil
}
func (q *DeviceQuery) GetAll() ([]OrgLevelDevice, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out []OrgLevelDevice
	err := q.query.Get(&out)
	return out, err
}
func (q *DeviceQuery) GetSingle() (*OrgLevelDevice, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out OrgLevelDevice
	err := q.query.Get(&out)
	return &out, err
}

type DeviceAutoverifyQuery struct {
	api   *AxcientApi
	query *goquadac.ApiQuery

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

	query := q.api.NewGetQuery(endpoint)
	q.query = query
	return q, nil
}
func (q *DeviceAutoverifyQuery) GetAll() ([]OrgLevelMachineAutoverifyDetails, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out []OrgLevelMachineAutoverifyDetails
	err := q.query.Get(&out)
	return out, err
}

type DeviceRestorePointQuery struct {
	api   *AxcientApi
	query *goquadac.ApiQuery

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

	query := q.api.NewGetQuery(endpoint)
	q.query = query
	return q, nil
}
func (q *DeviceRestorePointQuery) GetAll() ([]MachineRestorePoint, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out []MachineRestorePoint
	err := q.query.Get(&out)
	return out, err
}
