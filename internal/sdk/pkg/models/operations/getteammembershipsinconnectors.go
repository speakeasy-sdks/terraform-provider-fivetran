// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetTeamMembershipsInConnectorsRequest struct {
	Accept *string `header:"style=simple,explode=false,name=Accept"`
	// Paging cursor, [read more about pagination](https://fivetran.com/docs/rest-api/pagination)
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of records to fetch per page. Accepts a number in the range 1..1000; the default value is 100
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	// The unique identifier for the team within the account
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
}

type GetTeamMembershipsInConnectorsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	GetTeamMembershipsInConnectors200ApplicationJSONAny interface{}
}
