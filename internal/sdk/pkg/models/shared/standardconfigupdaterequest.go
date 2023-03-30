// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// StandardConfigUpdateRequestSchemaChangeHandlingEnum - The possible values for the schema_change_handling parameter are as follows: <br /> ALLOW_ALL - all new schemas, tables, and columns which appear in the source after the initial setup are included in syncs <br /> ALLOW_COLUMNS - all new schemas and tables which appear in the source after the initial setup are excluded from syncs, but new columns are included <br /> BLOCK_ALL - all new schemas, tables, and columns which appear in the source after the initial setup are excluded from syncs
type StandardConfigUpdateRequestSchemaChangeHandlingEnum string

const (
	StandardConfigUpdateRequestSchemaChangeHandlingEnumAllowAll     StandardConfigUpdateRequestSchemaChangeHandlingEnum = "ALLOW_ALL"
	StandardConfigUpdateRequestSchemaChangeHandlingEnumAllowColumns StandardConfigUpdateRequestSchemaChangeHandlingEnum = "ALLOW_COLUMNS"
	StandardConfigUpdateRequestSchemaChangeHandlingEnumBlockAll     StandardConfigUpdateRequestSchemaChangeHandlingEnum = "BLOCK_ALL"
)

func (e *StandardConfigUpdateRequestSchemaChangeHandlingEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "ALLOW_ALL":
		fallthrough
	case "ALLOW_COLUMNS":
		fallthrough
	case "BLOCK_ALL":
		*e = StandardConfigUpdateRequestSchemaChangeHandlingEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for StandardConfigUpdateRequestSchemaChangeHandlingEnum: %s", s)
	}
}

type StandardConfigUpdateRequest struct {
	// The possible values for the schema_change_handling parameter are as follows: <br /> ALLOW_ALL - all new schemas, tables, and columns which appear in the source after the initial setup are included in syncs <br /> ALLOW_COLUMNS - all new schemas and tables which appear in the source after the initial setup are excluded from syncs, but new columns are included <br /> BLOCK_ALL - all new schemas, tables, and columns which appear in the source after the initial setup are excluded from syncs
	SchemaChangeHandling *StandardConfigUpdateRequestSchemaChangeHandlingEnum `json:"schema_change_handling,omitempty"`
	// The set of schemas within your connector schema config that are synced into the destination
	Schemas map[string]SchemaUpdateRequest `json:"schemas,omitempty"`
}
