// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type CreateDestinationRequest struct {
	Accept             *string                    `header:"style=simple,explode=false,name=Accept"`
	DestinationRequest *shared.DestinationRequest `request:"mediaType=application/json"`
}

type CreateDestinationResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	CreateDestination201ApplicationJSONAny interface{}
}
