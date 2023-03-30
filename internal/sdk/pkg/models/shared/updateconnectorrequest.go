// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UpdateConnectorRequestScheduleTypeEnum - The connector schedule configuration type. Supported values: auto, manual
type UpdateConnectorRequestScheduleTypeEnum string

const (
	UpdateConnectorRequestScheduleTypeEnumAuto   UpdateConnectorRequestScheduleTypeEnum = "auto"
	UpdateConnectorRequestScheduleTypeEnumManual UpdateConnectorRequestScheduleTypeEnum = "manual"
)

func (e *UpdateConnectorRequestScheduleTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "auto":
		fallthrough
	case "manual":
		*e = UpdateConnectorRequestScheduleTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateConnectorRequestScheduleTypeEnum: %s", s)
	}
}

// UpdateConnectorRequestSyncFrequencyEnum - The connector sync frequency in minutes
type UpdateConnectorRequestSyncFrequencyEnum string

const (
	UpdateConnectorRequestSyncFrequencyEnumFive                           UpdateConnectorRequestSyncFrequencyEnum = "5"
	UpdateConnectorRequestSyncFrequencyEnumFifteen                        UpdateConnectorRequestSyncFrequencyEnum = "15"
	UpdateConnectorRequestSyncFrequencyEnumThirty                         UpdateConnectorRequestSyncFrequencyEnum = "30"
	UpdateConnectorRequestSyncFrequencyEnumSixty                          UpdateConnectorRequestSyncFrequencyEnum = "60"
	UpdateConnectorRequestSyncFrequencyEnumOneHundredAndTwenty            UpdateConnectorRequestSyncFrequencyEnum = "120"
	UpdateConnectorRequestSyncFrequencyEnumOneHundredAndEighty            UpdateConnectorRequestSyncFrequencyEnum = "180"
	UpdateConnectorRequestSyncFrequencyEnumThreeHundredAndSixty           UpdateConnectorRequestSyncFrequencyEnum = "360"
	UpdateConnectorRequestSyncFrequencyEnumFourHundredAndEighty           UpdateConnectorRequestSyncFrequencyEnum = "480"
	UpdateConnectorRequestSyncFrequencyEnumSevenHundredAndTwenty          UpdateConnectorRequestSyncFrequencyEnum = "720"
	UpdateConnectorRequestSyncFrequencyEnumOneThousandFourHundredAndForty UpdateConnectorRequestSyncFrequencyEnum = "1440"
)

func (e *UpdateConnectorRequestSyncFrequencyEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "5":
		fallthrough
	case "15":
		fallthrough
	case "30":
		fallthrough
	case "60":
		fallthrough
	case "120":
		fallthrough
	case "180":
		fallthrough
	case "360":
		fallthrough
	case "480":
		fallthrough
	case "720":
		fallthrough
	case "1440":
		*e = UpdateConnectorRequestSyncFrequencyEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateConnectorRequestSyncFrequencyEnum: %s", s)
	}
}

type UpdateConnectorRequest struct {
	// The connector authorization settings. Check possible config formats in [create method](/openapi/reference/v1/operation/create_connector/)
	Auth interface{} `json:"auth,omitempty"`
	// The connector setup configuration. Check possible config formats in [create method](/openapi/reference/v1/operation/create_connector/)
	Config interface{} `json:"config,omitempty"`
	// The connector daily sync start time that we return only when the sync frequency is set to 1440 (which means 24 hours) and the daily_sync_time parameter was set using the Create a Connector or Modify a Connector request
	DailySyncTime *string `json:"daily_sync_time,omitempty"`
	// The boolean specifying whether the connector should be triggered to re-sync all historical data. If you set this parameter to TRUE, the next scheduled sync will be historical. If the value is FALSE or not specified, the connector will not re-sync historical data. NOTE: When the value is TRUE, only the next scheduled sync will be historical, all subsequent ones will be incremental. This parameter is set to FALSE once the historical sync is completed.
	IsHistoricalSync *bool `json:"is_historical_sync,omitempty"`
	// Specifies whether the connector should be paused after the free trial period has ended
	PauseAfterTrial *bool `json:"pause_after_trial,omitempty"`
	// Specifies whether the connector is paused
	Paused *bool `json:"paused,omitempty"`
	// Specifies whether the connector should be paused after the free trial period has ended
	PausedAfterTrial *bool `json:"paused_after_trial,omitempty"`
	// Specifies whether the setup tests should be run automatically. The default value is TRUE.
	RunSetupTests *bool `json:"run_setup_tests,omitempty"`
	// The connector schedule configuration type. Supported values: auto, manual
	ScheduleType *UpdateConnectorRequestScheduleTypeEnum `json:"schedule_type,omitempty"`
	// Schema status
	SchemaStatus *string `json:"schema_status,omitempty"`
	// The connector sync frequency in minutes
	SyncFrequency *UpdateConnectorRequestSyncFrequencyEnum `json:"sync_frequency,omitempty"`
	// Specifies whether we should trust the certificate automatically. The default value is FALSE. If a certificate is not trusted automatically, it has to be approved with [Certificates Management API Approve a destination certificate](https://fivetran.com/docs/rest-api/certificates#approveadestinationcertificate).
	TrustCertificates *bool `json:"trust_certificates,omitempty"`
	// Specifies whether we should trust the SSH fingerprint automatically. The default value is FALSE. If a fingerprint is not trusted automatically, it has to be approved with [Certificates Management API Approve a destination fingerprint](https://fivetran.com/docs/rest-api/certificates#approveadestinationfingerprint).
	TrustFingerprints *bool `json:"trust_fingerprints,omitempty"`
}
