// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteTeamMembershipInGroupRequest struct {
	Accept *string `header:"style=simple,explode=false,name=Accept"`
	// The unique identifier for the group within the account
	GroupID string `pathParam:"style=simple,explode=false,name=groupId"`
	// The unique identifier for the team within the account
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
}

type DeleteTeamMembershipInGroupResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	DeleteTeamMembershipInGroup200ApplicationJSONAny interface{}
}
