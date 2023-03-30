// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DbtModelDetailsRequest struct {
	Accept *string `header:"style=simple,explode=false,name=Accept"`
	// The unique identifier for the DBT Model within the Fivetran system.
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

type DbtModelDetailsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	DbtModelDetails200ApplicationJSONAny interface{}
}