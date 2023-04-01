package client

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type OAuth2TokenResponse struct {
	// The new API token for the user
	AccessToken string `json:"access_token,omitempty"`
	// Will always be 'Bearer'
	TokenType string `json:"token_type,omitempty"`
}

// OAuth2Authorize OAuth 2.0 authorize endpoint
// Redirect users to this URL to begin the OAuth process. It is not an API endpoint.
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param clientId The client id
// @param responseType The response type (always 'code')
// @param optional state string
// @param redirectUri The redirect url
func (c *APIClient) OAuth2Authorize(ctx context.Context, clientId, responseType, state, redirectUri string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := c.cfg.BasePath + "/oauth2/authorize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("client_id", parameterToString(clientId, ""))
	localVarQueryParams.Add("response_type", parameterToString(responseType, ""))
	if state != "" {
		localVarQueryParams.Add("state", parameterToString(state, ""))
	}
	localVarQueryParams.Add("redirect_uri", parameterToString(redirectUri, ""))
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
		return localVarHttpResponse, GenericError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
	}

	return localVarHttpResponse, nil
}

// RequestOauth2Token Request an OAuth2 token
// Request an OAuth2 token from an authorization code. This request completes the OAuth process.
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param clientId
// @param clientSecret
// @param grantType
// @param code
// @param redirectUri
// @return OAuth2TokenResponse
func (c *APIClient) RequestOauth2Token(ctx context.Context, clientId string, clientSecret string, grantType string, code string, redirectUri string) (OAuth2TokenResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue OAuth2TokenResponse
	)

	// create path and map variables
	localVarPath := c.cfg.BasePath + "/oauth2/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
	localVarFormParams.Add("client_id", parameterToString(clientId, ""))
	localVarFormParams.Add("client_secret", parameterToString(clientSecret, ""))
	localVarFormParams.Add("grant_type", parameterToString(grantType, ""))
	localVarFormParams.Add("code", parameterToString(code, ""))
	localVarFormParams.Add("redirect_uri", parameterToString(redirectUri, ""))
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
	return localVarReturnValue, localVarHttpResponse, err
}
