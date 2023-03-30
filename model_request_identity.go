/*
 * JupyterHub
 *
 * The REST API for JupyterHub
 *
 * API version: 4.0.0b2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

// The model for the entity making the request. Extends User or Service model to add information about the specific credentials (e.g. session or token-authorised scopes).
type RequestIdentity struct {
	// The session id associated with the request's OAuth token, if any. null, if the request token not associated with a session id.  Added in 2.0.
	SessionId string `json:"session_id,omitempty"`
	// The list of all expanded scopes the request credentials have access to.  Added in 2.0.
	Scopes []string `json:"scopes,omitempty"`
}