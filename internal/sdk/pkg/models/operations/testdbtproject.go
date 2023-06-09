// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type TestDbtProjectRequest struct {
	Accept *string `header:"style=simple,explode=false,name=Accept"`
	// The unique identifier for the DBT Project within the Fivetran system.
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
}

type TestDbtProjectResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	TestDbtProject200ApplicationJSONAny interface{}
}
