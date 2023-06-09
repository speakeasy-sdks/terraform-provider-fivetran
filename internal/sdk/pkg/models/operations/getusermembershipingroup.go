// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetUserMembershipInGroupRequest struct {
	Accept *string `header:"style=simple,explode=false,name=Accept"`
	// The unique identifier for the group within the account.
	GroupID string `pathParam:"style=simple,explode=false,name=groupId"`
	// The unique identifier for the user within the account.
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

type GetUserMembershipInGroupResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	GetUserMembershipInGroup200ApplicationJSONAny interface{}
}
