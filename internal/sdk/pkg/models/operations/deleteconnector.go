// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteConnectorRequest struct {
	Accept *string `header:"style=simple,explode=false,name=Accept"`
	// The unique identifier for the connector within the Fivetran system
	ConnectorID string `pathParam:"style=simple,explode=false,name=connectorId"`
}

type DeleteConnectorResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	DeleteConnector200ApplicationJSONAny interface{}
}
