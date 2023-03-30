// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type ConnectCardRequest struct {
	Accept                   *string                          `header:"style=simple,explode=false,name=Accept"`
	ConnectCardConfigRequest *shared.ConnectCardConfigRequest `request:"mediaType=application/json"`
	// The unique identifier for the connector within the Fivetran system
	ConnectorID string `pathParam:"style=simple,explode=false,name=connectorId"`
}

type ConnectCardResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	ConnectCard200ApplicationJSONAny interface{}
}