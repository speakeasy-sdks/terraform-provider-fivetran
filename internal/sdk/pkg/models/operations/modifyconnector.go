// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type ModifyConnectorRequest struct {
	Accept                 *string                        `header:"style=simple,explode=false,name=Accept"`
	UpdateConnectorRequest *shared.UpdateConnectorRequest `request:"mediaType=application/json"`
	// The unique identifier for the connector within the Fivetran system
	ConnectorID string `pathParam:"style=simple,explode=false,name=connectorId"`
}

type ModifyConnectorResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	ModifyConnector200ApplicationJSONAny interface{}
}