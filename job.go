package axcient

import (
	"fmt"
	"time"

	"github.com/simonbuckner/goquadac"
)

type JobQuery struct {
	api   *AxcientApi
	query *goquadac.ApiQuery

	// URL Path fields
	clientId *int64
	deviceId *int64
	jobId    *int64

	// URL query fields
}

func newJobQuery(api *AxcientApi) *JobQuery {
	return &JobQuery{
		api: api,
	}
}

func (q *JobQuery) SelectByClientId(clientId int64) *JobQuery {
	q.clientId = &clientId
	return q
}

func (q *JobQuery) SelectByDeviceId(deviceId int64) *JobQuery {
	q.deviceId = &deviceId
	return q
}

func (q *JobQuery) SelectByJobId(jobId int64) *JobQuery {
	q.jobId = &jobId
	return q
}

func (q *JobQuery) Build() (*JobQuery, error) {

	if q.clientId == nil {
		return nil, fmt.Errorf("client_id required")
	}
	if q.deviceId == nil && q.jobId != nil {
		return nil, fmt.Errorf("device_id is required in job_id is specified")
	}

	var client string = ""
	var device string = ""
	var job string = ""

	client = fmt.Sprintf("client/%d", *q.clientId)
	if q.deviceId != nil {
		device = fmt.Sprintf("/device/%d/job", *q.deviceId)
	}
	if q.jobId != nil {
		job = fmt.Sprintf("/%d", *q.jobId)
	}

	endpoint := client + device + job

	query := q.api.NewGetQuery(endpoint).
		SetDumpRequest(true).
		SetDumpResponse(true).
		SetDumpResponseBody(true)

	q.query = query
	return q, nil
}
func (q *JobQuery) GetAll() ([]OrgLevelJobResponse, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out []OrgLevelJobResponse
	err := q.query.Get(&out)
	return out, err
}

func (q *JobQuery) GetSingle() (*OrgLevelJobResponse, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out OrgLevelJobResponse
	err := q.query.Get(&out)
	return &out, err
}

// -------------------------------------------------------------------------
//   GET:/client/{client_id}/device/{device_id}/job/{job_id}/history
// -------------------------------------------------------------------------

type JobHistoryQuery struct {
	api   *AxcientApi
	query *goquadac.ApiQuery

	// URL Path fields
	clientId *int64
	deviceId *int64
	jobId    *int64

	// URL query fields
	limit          *int64
	offset         *int64
	startTimeBegin *int64
}

func newJobHistoryQuery(api *AxcientApi) *JobHistoryQuery {
	return &JobHistoryQuery{
		api: api,
	}
}

func (q *JobHistoryQuery) SelectByClientId(clientId int64) *JobHistoryQuery {
	q.clientId = &clientId
	return q
}

func (q *JobHistoryQuery) SelectByDeviceId(deviceId int64) *JobHistoryQuery {
	q.deviceId = &deviceId
	return q
}

func (q *JobHistoryQuery) SelectByJobId(jobId int64) *JobHistoryQuery {
	q.jobId = &jobId
	return q
}

func (q *JobHistoryQuery) SetLimit(limit int64) *JobHistoryQuery {
	q.limit = &limit
	return q
}
func (q *JobHistoryQuery) SetOffset(offset int64) *JobHistoryQuery {
	q.offset = &offset
	return q
}

func (q *JobHistoryQuery) SetStartTimeBegin(startTime time.Time) *JobHistoryQuery {
	unixTime := startTime.Unix()
	q.startTimeBegin = &unixTime
	return q
}

func (q *JobHistoryQuery) Build() (*JobHistoryQuery, error) {

	if q.clientId == nil {
		return nil, fmt.Errorf("client_id required")
	}
	if q.deviceId == nil {
		return nil, fmt.Errorf("device_id required")
	}
	if q.jobId == nil {
		return nil, fmt.Errorf("job_id required")
	}

	endpoint := fmt.Sprintf("client/%d/device/%d/job/%d/history", *q.clientId, *q.deviceId, *q.jobId)
	query := q.api.NewGetQuery(endpoint).
		SetDumpRequest(true).
		SetDumpResponse(true).
		SetDumpResponseBody(true)

	if q.limit != nil {
		query.AddUrlQuery("limit", goquadac.I64toString(*q.limit))
	}
	if q.offset != nil {
		query.AddUrlQuery("offset", goquadac.I64toString(*q.offset))
	}
	if q.startTimeBegin != nil {
		query.AddUrlQuery("starttime_begin", goquadac.I64toString(*q.startTimeBegin))
	}
	q.query = query
	return q, nil
}
func (q *JobHistoryQuery) GetAll() ([]JobHistory, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out []JobHistory
	err := q.query.Get(&out)
	return out, err
}
