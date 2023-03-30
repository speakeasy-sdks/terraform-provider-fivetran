// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConnectCardConfig struct {
	// An optional parameter that lets you hide the embedded setup guide in the Connect Card.
	HideSetupGuide *bool `json:"hide_setup_guide,omitempty"`
	// The URI on your site we redirect the end user to after successful setup. The URI must start with the `https` or `http` prefix.
	RedirectURI *string `json:"redirect_uri,omitempty"`
}