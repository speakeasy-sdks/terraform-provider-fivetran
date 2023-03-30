// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GroupDetailsRequest struct {
	Accept *string `header:"style=simple,explode=false,name=Accept"`
	// The unique identifier for the group within the Fivetran system.
	GroupID string `pathParam:"style=simple,explode=false,name=groupId"`
}

type GroupDetailsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	GroupDetails200ApplicationJSONAny interface{}
}