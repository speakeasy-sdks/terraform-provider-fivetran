// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type ModifyConnectorDatabaseSchemaConfigRequest struct {
	Accept              *string                     `header:"style=simple,explode=false,name=Accept"`
	SchemaUpdateRequest *shared.SchemaUpdateRequest `request:"mediaType=application/json"`
	// The unique identifier for the connector within the Fivetran system
	ConnectorID string `pathParam:"style=simple,explode=false,name=connectorId"`
	// The database schema name within your destination (different from the connector schema)
	SchemaName string `pathParam:"style=simple,explode=false,name=schemaName"`
}

type ModifyConnectorDatabaseSchemaConfigResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	ModifyConnectorDatabaseSchemaConfig200ApplicationJSONAny interface{}
}
