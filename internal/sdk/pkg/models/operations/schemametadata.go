// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type SchemaMetadataRequest struct {
	Accept *string `header:"style=simple,explode=false,name=Accept"`
	// The unique identifier for the connector within your Fivetran account
	ConnectorID string `pathParam:"style=simple,explode=false,name=connectorId"`
	// Paging cursor, [read more about pagination](https://fivetran.com/docs/rest-api/pagination)
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of records to fetch per page. Accepts a number in the range 1..1000; the default value is 100.
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
}

type SchemaMetadataResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	SchemaMetadata200ApplicationJSONAny interface{}
}
