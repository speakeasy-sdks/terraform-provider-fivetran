// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteUserMembershipInAccountRequest struct {
	Accept *string `header:"style=simple,explode=false,name=Accept"`
	// The unique identifier for the user within the account.
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

type DeleteUserMembershipInAccountResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	DeleteUserMembershipInAccount200ApplicationJSONAny interface{}
}
