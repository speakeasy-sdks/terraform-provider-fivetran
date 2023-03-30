// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// MembershipResponse - Successful response
type MembershipResponse struct {
	// The date and time the membership was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The membership entity unique identifier
	ID *string `json:"id,omitempty"`
	// The role the user has within the entity
	Role *string `json:"role,omitempty"`
}
