// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type CreateTeamRequest struct {
	Accept      *string             `header:"style=simple,explode=false,name=Accept"`
	TeamRequest *shared.TeamRequest `request:"mediaType=application/json"`
}

type CreateTeamResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// default response
	CreateTeam201ApplicationJSONAny interface{}
}