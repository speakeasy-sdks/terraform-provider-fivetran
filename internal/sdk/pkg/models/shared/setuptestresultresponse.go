// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SetupTestResultResponse struct {
	// Setup test details.
	Details map[string]interface{} `json:"details,omitempty"`
	// Setup test message.
	Message *string `json:"message,omitempty"`
	// The current state of the connector.
	Status *string `json:"status,omitempty"`
	// Setup test title.
	Title *string `json:"title,omitempty"`
}
