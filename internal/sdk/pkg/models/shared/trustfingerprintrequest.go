// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TrustFingerprintRequest struct {
	// The unique identifier for the connector
	ConnectorID *string `json:"connector_id,omitempty"`
	// The unique identifier for the destination
	DestinationID *string `json:"destination_id,omitempty"`
	// Hash of the fingerprint
	Hash string `json:"hash"`
	// The SSH public key
	PublicKey string `json:"public_key"`
}
