// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ResyncConnectorRequest struct {
	// A map containing an array of tables to re-sync for each schema, must be non-empty. The parameter is optional
	Scope map[string][]string `json:"scope,omitempty"`
}
