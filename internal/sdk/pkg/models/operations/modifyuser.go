// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type ModifyUserRequest struct {
	Accept            *string                   `header:"style=simple,explode=false,name=Accept"`
	UpdateUserRequest *shared.UpdateUserRequest `request:"mediaType=application/json"`
	// The unique identifier for the user within the account.
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

type ModifyUserResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	ModifyUser200ApplicationJSONAny interface{}
}
