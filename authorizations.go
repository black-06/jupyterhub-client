package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Token struct {
	// The token itself. Only present in responses to requests for a new token.
	Token string `json:"token,omitempty"`
	// The id of the API token. Used for modifying or deleting the token.
	Id string `json:"id,omitempty"`
	// The user that owns a token (undefined if owned by a service)
	User string `json:"user,omitempty"`
	// The service that owns the token (undefined of owned by a user)
	Service string `json:"service,omitempty"`
	// Deprecated in JupyterHub 3, always an empty list. Tokens have 'scopes' starting from JupyterHub 3.
	Roles []string `json:"roles,omitempty"`
	// List of scopes this token has been assigned. New in JupyterHub 3. In JupyterHub 2.x, tokens were assigned 'roles' insead of scopes.
	Scopes []string `json:"scopes,omitempty"`
	// A note about the token, typically describing what it was created for.
	Note string `json:"note,omitempty"`
	// Timestamp when this token was created
	Created time.Time `json:"created,omitempty"`
	// Timestamp when this token expires. Null if there is no expiry.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Timestamp of last-seen activity using this token. Can be null if token has never been used.
	LastActivity time.Time `json:"last_activity,omitempty"`
	// The session id associated with the token, if any. Only used for tokens set during oauth flows.  Added in 2.0.
	SessionId string `json:"session_id,omitempty"`
}

// RequestNewAPIToken Request a new API token.
// Request a new API token to use with the JupyterHub REST API. If not already authenticated, username and password can be sent in the JSON request body. Logging in via this method is only available when the active Authenticator accepts passwords (e.g. not OAuth).
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param optional BasicAuth - optional.
// @return token value.
func (c *APIClient) RequestNewAPIToken(ctx context.Context, auth *BasicAuth) (string, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue struct {
			// The new API token.
			Token string `json:"token,omitempty"`
		}
	)

	// create path and map variables
	localVarPath := c.cfg.BasePath + "/authorizations/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = auth

	r, err := c.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return "", nil, err
	}

	localVarHttpResponse, err := c.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return "", localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	_ = localVarHttpResponse.Body.Close()
	if err != nil {
		return "", localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return "", localVarHttpResponse, GenericError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
	}

	// If we succeed, return the data, otherwise pass on to decode error.
	err = c.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err == nil {
		return "", localVarHttpResponse, err
	}
	return localVarReturnValue.Token, localVarHttpResponse, nil
}

// IdentifyUserOrServiceFromAPIToken Identify a user or service from an API token.
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param token
func (c *APIClient) IdentifyUserOrServiceFromAPIToken(ctx context.Context, token string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := c.cfg.BasePath + "/authorizations/token/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", fmt.Sprintf("%v", token), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	var localVarHttpContentTypes []string

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	var localVarHttpHeaderAccepts []string

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := c.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := c.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	_ = localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

// IdentifyUserFromCookie identify a user from a cookie.
// Used by single-user notebook servers to hand off cookie authentication to the Hub
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param cookieName
// @param cookieValue
// @return User
func (c *APIClient) IdentifyUserFromCookie(ctx context.Context, cookieName string, cookieValue string) (User, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue User
	)

	// create path and map variables
	localVarPath := c.cfg.BasePath + "/authorizations/cookie/{cookie_name}/{cookie_value}"
	localVarPath = strings.Replace(localVarPath, "{"+"cookie_name"+"}", fmt.Sprintf("%v", cookieName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cookie_value"+"}", fmt.Sprintf("%v", cookieValue), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	var localVarHttpContentTypes []string

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := c.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := c.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	_ = localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHttpResponse, GenericError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
	}
	// If we succeed, return the data, otherwise pass on to decode error.
	err = c.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	return localVarReturnValue, localVarHttpResponse, nil
}
