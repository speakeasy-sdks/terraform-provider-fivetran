// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type AddUserToGroupRequest struct {
	Accept                *string                       `header:"style=simple,explode=false,name=Accept"`
	AddUserToGroupRequest *shared.AddUserToGroupRequest `request:"mediaType=application/json"`
	// The unique identifier for the group within the Fivetran system.
	GroupID string `pathParam:"style=simple,explode=false,name=groupId"`
}

type AddUserToGroupResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	AddUserToGroup200ApplicationJSONAny interface{}
}