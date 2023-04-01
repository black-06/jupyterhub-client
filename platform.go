package client

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type VersionResponse struct {
	// The version of JupyterHub itself
	Version string `json:"version,omitempty"`
}

type DetailedInfo struct {
	// The version of JupyterHub itself
	Version string `json:"version,omitempty"`
	// The Python version, as returned by sys.version
	Python string `json:"python,omitempty"`
	// The path to sys.executable running JupyterHub
	SysExecutable string         `json:"sys_executable,omitempty"`
	Authenticator *Authenticator `json:"authenticator,omitempty"`
	Spawner       *Spawner       `json:"spawner,omitempty"`
}

type Authenticator struct {
	// The Python class currently active for JupyterHub Authentication
	Class string `json:"class,omitempty"`
	// The version of the currently active Authenticator
	Version string `json:"version,omitempty"`
}

type Spawner struct {
	// The Python class currently active for spawning single-user notebook servers
	Class string `json:"class,omitempty"`
	// The version of the currently active Spawner
	Version string `json:"version,omitempty"`
}

// GetVersion Get JupyterHub version.
// This endpoint is not authenticated for the purpose of clients and user to identify the JupyterHub version before setting up authentication.
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return VersionResponse
func (c *APIClient) GetVersion(ctx context.Context) (VersionResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue VersionResponse
	)

	// create path and map variables
	localVarPath := c.cfg.BasePath + "/"

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
	return localVarReturnValue, localVarHttpResponse, err
}

// GetDetailedInfo Get detailed info about JupyterHub.
// Detailed JupyterHub information, including Python version, JupyterHub's version and executable path, and which Authenticator and Spawner are active.
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return DetailedInfo
func (c *APIClient) GetDetailedInfo(ctx context.Context) (DetailedInfo, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue DetailedInfo
	)

	// create path and map variables
	localVarPath := c.cfg.BasePath + "/info"

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

type ShutdownRequest struct {
	// Whether the proxy should be shutdown as well (default from Hub config)
	Proxy bool `json:"proxy,omitempty"`
	// Whether users' notebook servers should be shutdown as well (default from Hub config)
	Servers bool `json:"servers,omitempty"`
}

// Shutdown the Hub.
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param optional nil or *ShutdownOpts - Optional Parameters.
func (c *APIClient) Shutdown(ctx context.Context, request *ShutdownRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := c.cfg.BasePath + "/shutdown"

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
	var localVarHttpHeaderAccepts []string

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = request
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
