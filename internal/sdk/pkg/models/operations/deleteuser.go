// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteUserRequest struct {
	Accept *string `header:"style=simple,explode=false,name=Accept"`
	// The unique identifier for the user within the account.
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

type DeleteUserResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	DeleteUser200ApplicationJSONAny interface{}
}
