package axcientapi

import (
	"strings"
	"time"
)

// Source schemas can be found here :https://developer.axcient.com/x360recover/

type OrgLevelAppliance struct {
	Id                 int64                          `json:"id"`
	ServiceId          string                         `json:"service_id"`
	ClientId           int64                          `json:"client_id"`
	Alias              string                         `json:"alias"`
	IpAddress          string                         `json:"ip_address"`
	ServerId           string                         `json:"server_id"`
	Active             bool                           `json:"active"`
	Product            string                         `json:"product"`
	Model              *OrgLevelServerModel           `json:"model"`
	Package            *OrgLevelPackage               `json:"package"`
	ServiceType        string                         `json:"service_type"`
	LastTunnelUp       time.Time                      `json:"last_tunnel_up"`
	TunnelStatus       string                         `json:"tunnel_status"`
	StorageDetails     *StorageDetail                 `json:"storage_details"`
	Devices            []OrgLevelShortApplianceDevice `json:"devices"`
	SoftwareVersion    *OrgLevelVersion               `json:"software_version"`
	HealthStatus       string                         `json:"health_status"`
	HealthStatusReason string                         `json:"health_status_reason"`
}

type OrgLevelShortAppliance struct {
	Id          int64  `json:"id"`
	ServiceId   string `json:"service_id"`
	Alias       string `json:"alias"`
	IpAddress   string `json:"ip_address"`
	ServerId    string `json:"server_id"`
	Active      bool   `json:"active"`
	Product     string `json:"product"`
	ServiceType string `json:"service_type"`
}

type OrgLevelVault struct {
	Id                 int64                 `json:"id"`
	ServiceId          string                `json:"service_id"`
	Name               string                `json:"name"`
	IpAddress          string                `json:"ip_address"`
	ServerId           string                `json:"server_id"`
	Active             bool                  `json:"active"`
	Type               string                `json:"type"`
	Model              *OrgLevelServerModel  `json:"model"`
	ServiceType        string                `json:"service_type"`
	LastTunnelUp       time.Time             `json:"last_tunnel_up"`
	TunnelStatus       string                `json:"tunnel_status"`
	StorageDetails     *StorageDetail        `json:"storage_details"`
	Devices            []OrgLevelShortDevice `json:"devices"`
	SoftwareVersion    *OrgLevelVersion      `json:"software_version"`
	ReplicationData    int64                 `json:"replication_data"`
	VaultThresholds    *VaultThresholds      `json:"vault_thresholds"`
	HealthStatus       string                `json:"health_status"`
	HealthStatusReason string                `json:"health_status_reason"`
	CreationTimestamp  time.Time             `json:"creation_timestamp"`
}

type OrgLevelShortVault struct {
	Id            int64     `json:"id"`
	Type          string    `json:"type"`
	LatestVaultRp time.Time `json:"latest_vault_rp"`
	DeviceUsage   int64     `json:"device_usage"`
}

type OrgLevelDevice struct {
	Id                      int64                            `json:"id"`
	Name                    string                           `json:"name"`
	ClientId                int64                            `json:"client_id"`
	Type                    string                           `json:"type"`
	Os                      *DeviceOs                        `json:"os"`
	ServiceId               string                           `json:"service_id"`
	Product                 string                           `json:"product"`
	LocalPsId               string                           `json:"local_ps_id"`
	Vaults                  []OrgLevelVault                  `json:"vaults"`
	IpAddress               string                           `json:"ip_address"`
	CurrentHealthStatus     *HealthStatusObj                 `json:"current_health_status"`
	PreviousHealthStatus    *HealthStatusObj                 `json:"previous_health_status"`
	Thresholds              *DeviceThresholds                `json:"thresholds"`
	LocalUsage              int64                            `json:"local_usage"`
	LocalTotal              int64                            `json:"local_total"`
	CloudUsage              int64                            `json:"cloud_usage"`
	VaultUsage              int64                            `json:"vault_usage"`
	Jobs                    []OrgLevelShortJob               `json:"jobs"`
	LatestCloudRp           time.Time                        `json:"latest_cloud_rp"`
	LatestLocalRp           time.Time                        `json:"latest_local_rp"`
	LatestVaultRp           time.Time                        `json:"latest_vault_rp"`
	Direct2Cloud            bool                             `json:"d2c"`
	AsioEndpointId          string                           `json:"asio_endpoint _id"`
	AgentVersion            string                           `json:"agent_version"`
	Volumes                 []string                         `json:"volumes"`
	ExcludedVolumes         []string                         `json:"excluded_volumes"`
	DeviceDetailsPageUrl    string                           `json:"device_details_page_url"`
	LocalCacheDetails       *LocalCacheDetails               `json:"local_cache_details"`
	LatestAutoverifyDetails *OrgLevelLatestAutoverifyDetails `json:"latest_autoverify_details"`
}

type OrgLevelShortDevice struct {
	ApplianceId int64                `json:"appliance_id"`
	IpAddress   string               `json:"ip_address"`
	Os          string               `json:"os"`
	Id          int64                `json:"id"`
	Name        string               `json:"name"`
	Client      *OrgLevelShortClient `json:"client"`
}

type OrgLevelShortApplianceDevice struct {
	IpAddress     string    `json:"ip_address"`
	Os            string    `json:"os"`
	Id            int64     `json:"id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	LocalUsage    int64     `json:"local_usage"`
	LatestLocalRp time.Time `json:"latest_local_rp"`
}

type OrgLevelShortJob struct {
	// Common fields
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	HealthStatus string `json:"health_status"`
	//
	// org_level_brc_job fields
	LatestLocalRp *time.Time `json:"latest_local_rp"`
	LatestCloudRp *time.Time `json:"latest_cloud_rp"`
	//
	// org_level_replibit_job fields
	VaultId  int64      `json:"vault_id"`
	LatestRp *time.Time `json:"latest_rp"`
}
type OrgLevelJob struct {
	// Common fields
	Id           int64          `json:"id"`
	Name         string         `json:"name"`
	Offsite      bool           `json:"offsite"`
	Enabled      bool           `json:"enabled"`
	Thresholds   *JobThresholds `json:"thresholds"`
	HealthStatus string         `json:"health_status"`
	//
	// org_level_brc_job fields
	JobType       *string    `json:"job_type"`
	LatestLocalRp *time.Time `json:"latest_local_rp"`
	LatestCloudRp *time.Time `json:"latest_cloud_rp"`
	//
	// org_level_replibit_job fields
	Vault    *OrgLevelVault `json:"vault"`
	Schedule *string        `json:"schedule"`
	LatestRp *time.Time     `json:"latest_rp"`
}

type JobHistory struct {
	Status    string    `json:"status"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Rp        time.Time `json:"rp"`
	ErrorMsg  string    `json:"error_msg"`
}

type OrgLevelServerModel struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type OrgLevelPackage struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type OrgLevelVersion struct {
	Id      int64  `json:"id"`
	Version string `json:"version"`
}

type OrgLevelShortClient struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type OrgLevelClient struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	ClientCode string `json:"client_code"`
	Active     int64  `json:"active"`
	// Active          bool                            `json:"active"`
	HealthStatus    string                          `json:"health_status"`
	DevicesCounters *ClientProtectedSystemsCounters `json:"devices_counters"`
}
type ClientProtectedSystemsCounters struct {
	ApplianceBased []ProtectedSystemCounters `json:"appliance_based"`
	Direct2Cloud   []ProtectedSystemCounters `json:"d2c"`
	CloudArchive   []ProtectedSystemCounters `json:"cloud_archive"`
}

type ProtectedSystemCounters struct {
	Type  string `json:"type"`
	Count int64  `json:"count"`
}

type StorageDetail struct {
	DriveSize int64 `json:"drive_size"`
	UsedSize  int64 `json:"used_size"`
}

type DeviceOs struct {
	OsType string `json:"os_type"`
	Bits   int64  `json:"bits"`
	OsName string `json:"os_name"`
}

type LocalCacheDetails struct {
	Enabled                             bool      `json:"enabled"`
	LastSuccessfulVerificationTimestamp time.Time `json:"last_successful_verification_timestamp"`
	LastVerificationTimestamp           time.Time `json:"last_verification_timestamp"`
	Path                                string    `json:"path"`
	VerificationStatus                  string    `json:"verification_status"`
}

type OrgLevelLatestAutoverifyDetails struct {
	Id                     string    `json:"id"`
	Timestamp              time.Time `json:"timestamp"`
	StartTimestamp         time.Time `json:"start_timestamp"`
	EndTimestamp           time.Time `json:"end_timestamp"`
	Rp                     time.Time `json:"rp"`
	Status                 string    `json:"status"`
	ScreenshotUrl          string    `json:"screenshot_url"`
	ScreenshotThumbnailUrl string    `json:"screenshot_thumbnail_url"`
	IsHealthy              bool      `json:"is_healthy"`
}

type VaultThresholds struct {
	VaultId               int64 `json:"vault_id"`
	ConnectivityThreshold int64 `json:"connectivity_threshold"`
}

type VaultThresholdBody struct {
	Threshold int64 `json:"threshold"`
}

type DeviceThresholds struct {
	VaultRpThreshold    *ThresholdObj `json:"vault_rp_threshold"`
	CloudRpThreshold    *ThresholdObj `json:"cloud_rp_threshold"`
	LocalRpThreshold    *ThresholdObj `json:"local_rp_threshold"`
	ProtectionThreshold *ThresholdObj `json:"protection_threshold"`
}

type JobThresholds struct {
	VaultRpThreshold    *ThresholdObj `json:"vault_rp_threshold"`
	CloudRpThreshold    *ThresholdObj `json:"cloud_rp_threshold"`
	LocalRpThreshold    *ThresholdObj `json:"local_rp_threshold"`
	ProtectionThreshold *ThresholdObj `json:"protection_threshold"`
}

type ThresholdObj struct {
	Value      int64 `json:"value"`
	Enabled    bool  `json:"enabled"`
	Overridden bool  `json:"oberridden"`
}

type OrgLevelJobResponse OrgLevelJob

type HealthStatusObj struct {
	Status    string    `json:"status"`
	Reason    string    `json:"reason"`
	Timestamp time.Time `json:"timestamp"`
}

type OrgLevelOrganisation struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Active       bool   `json:"active"`
	BrandId      string `json:"brand_id"`
	SalesforceId string `json:"salesforce_id"`
}

type MachineRestorePoint struct {
	VaultId       int64          `json:"vault_id"`
	Status        string         `json:"status"`
	ErrorMsg      string         `json:"error_msg"`
	RestorePoints []RestorePoint `json:"restore_point"`
}

type RestorePoint struct {
	Timestamp      RestorePointTimestamp `json:"timestamp"`
	InUse          bool                  `json:"in_use"`
	UsageInitiator string                `json:"usage_initiator"`
}

// org_level_machine_autoverify_details
type OrgLevelMachineAutoverifyDetails struct {
	VaultId           int64                    `json:"vault_id"`
	ApplianceId       int64                    `json:"appliance_id"`
	AutoverifyDetails []OrgLevelAutoverifyInfo `json:"autoverify_details"`
}

type OrgLevelAutoverifyInfo struct {
	Id                     string    `json:"id"`
	Timestamp              time.Time `json:"timestamp"`
	StartTimestamp         time.Time `json:"start_timestamp"`
	EndTimestamp           time.Time `json:"end_timestamp"`
	Rp                     time.Time `json:"rp"`
	Status                 string    `json:"status"`
	ScreenshotUrl          string    `json:"screenshot_url"`
	ScreenshotThumbnailUrl string    `json:"screenshot_thumbnail_url"`
	IsHealthy              bool      `json:"is_healthy"`
}
type ClientLevelDirect2CloudAgentToken struct {
	TokenId string `json:"token_id"`
}

/*
 *  Custom fields to plug holes in their API
 *
 */

// This particular timestamp has no timezone so cannot be parsed by the default RFC
type RestorePointTimestamp time.Time

func (t *RestorePointTimestamp) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	out, err := time.Parse("2006-01-02T15:04:05", value) //parse time
	if err != nil {
		return err
	}
	*t = RestorePointTimestamp(out) //set result using the pointer
	return nil
}
