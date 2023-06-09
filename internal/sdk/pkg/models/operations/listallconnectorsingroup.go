// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ListAllConnectorsInGroupRequest struct {
	Accept *string `header:"style=simple,explode=false,name=Accept"`
	// Paging cursor, [read more about pagination](https://fivetran.com/docs/rest-api/pagination)
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// The unique identifier for the group within the Fivetran system.
	GroupID string `pathParam:"style=simple,explode=false,name=groupId"`
	// Number of records to fetch per page. Accepts a number in the range 1..1000; the default value is 100.
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	// The name used both as the connector's name within the Fivetran system and as the source schema's name within your destination.
	Schema *string `queryParam:"style=form,explode=true,name=schema"`
}

type ListAllConnectorsInGroupResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	ListAllConnectorsInGroup200ApplicationJSONAny interface{}
}
