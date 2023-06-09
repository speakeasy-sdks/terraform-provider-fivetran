// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TransformationScheduleDaysOfWeekEnum string

const (
	TransformationScheduleDaysOfWeekEnumMonday    TransformationScheduleDaysOfWeekEnum = "MONDAY"
	TransformationScheduleDaysOfWeekEnumTuesday   TransformationScheduleDaysOfWeekEnum = "TUESDAY"
	TransformationScheduleDaysOfWeekEnumWednesday TransformationScheduleDaysOfWeekEnum = "WEDNESDAY"
	TransformationScheduleDaysOfWeekEnumThursday  TransformationScheduleDaysOfWeekEnum = "THURSDAY"
	TransformationScheduleDaysOfWeekEnumFriday    TransformationScheduleDaysOfWeekEnum = "FRIDAY"
	TransformationScheduleDaysOfWeekEnumSaturday  TransformationScheduleDaysOfWeekEnum = "SATURDAY"
	TransformationScheduleDaysOfWeekEnumSunday    TransformationScheduleDaysOfWeekEnum = "SUNDAY"
)

func (e *TransformationScheduleDaysOfWeekEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "MONDAY":
		fallthrough
	case "TUESDAY":
		fallthrough
	case "WEDNESDAY":
		fallthrough
	case "THURSDAY":
		fallthrough
	case "FRIDAY":
		fallthrough
	case "SATURDAY":
		fallthrough
	case "SUNDAY":
		*e = TransformationScheduleDaysOfWeekEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for TransformationScheduleDaysOfWeekEnum: %s", s)
	}
}

// TransformationScheduleScheduleTypeEnum - Schedule type
type TransformationScheduleScheduleTypeEnum string

const (
	TransformationScheduleScheduleTypeEnumIntegrated TransformationScheduleScheduleTypeEnum = "INTEGRATED"
	TransformationScheduleScheduleTypeEnumTimeOfDay  TransformationScheduleScheduleTypeEnum = "TIME_OF_DAY"
	TransformationScheduleScheduleTypeEnumInterval   TransformationScheduleScheduleTypeEnum = "INTERVAL"
)

func (e *TransformationScheduleScheduleTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "INTEGRATED":
		fallthrough
	case "TIME_OF_DAY":
		fallthrough
	case "INTERVAL":
		*e = TransformationScheduleScheduleTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for TransformationScheduleScheduleTypeEnum: %s", s)
	}
}

type TransformationSchedule struct {
	// Days of week
	DaysOfWeek []TransformationScheduleDaysOfWeekEnum `json:"days_of_week,omitempty"`
	// Interval.
	Interval *int `json:"interval,omitempty"`
	// Schedule type
	ScheduleType *TransformationScheduleScheduleTypeEnum `json:"schedule_type,omitempty"`
	// Time of day
	TimeOfDay *string `json:"time_of_day,omitempty"`
}
