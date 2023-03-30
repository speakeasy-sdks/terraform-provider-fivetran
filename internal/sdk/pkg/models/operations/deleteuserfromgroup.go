// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteUserFromGroupRequest struct {
	Accept *string `header:"style=simple,explode=false,name=Accept"`
	// The unique identifier for the group within the Fivetran system.
	GroupID string `pathParam:"style=simple,explode=false,name=groupId"`
	// The unique identifier for the user within the Fivetran system.
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

type DeleteUserFromGroupResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	DeleteUserFromGroup200ApplicationJSONAny interface{}
}