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

type Server struct {
	// The server's name. The user's default server has an empty name ('')
	Name string `json:"name,omitempty"`
	// Whether the server is ready for traffic. Will always be false when any transition is pending.
	Ready bool `json:"ready,omitempty"`
	// Whether the server is stopped. Added in JupyterHub 3.0, and only useful when using the `?include_stopped_servers` request parameter. Now that stopped servers may be included (since JupyterHub 3.0), this is the simplest way to select stopped servers. Always equivalent to `not (ready or pending)`.
	Stopped bool `json:"stopped,omitempty"`
	// The currently pending action, if any. A server is not ready if an action is pending.
	Pending string `json:"pending,omitempty"`
	// The URL where the server can be accessed (typically /user/:name/:server.name/).
	Url string `json:"url,omitempty"`
	// The URL for an event-stream to retrieve events during a spawn.
	ProgressUrl string `json:"progress_url,omitempty"`
	// UTC timestamp when the server was last started.
	Started time.Time `json:"started,omitempty"`
	// UTC timestamp last-seen activity on this server.
	LastActivity time.Time `json:"last_activity,omitempty"`
	// Arbitrary internal state from this server's spawner. Only available on the hub's users list or get-user-by-name method, and only with admin:users:server_state scope. None otherwise.
	State *interface{} `json:"state,omitempty"`
	// User specified options for the user's spawned instance of a single-user server.
	UserOptions *interface{} `json:"user_options,omitempty"`
}

// StartUserServer Start a user's single-user notebook server
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param name username
// @param spawnOpts - Optional Spawn options can be passed as a JSON body
//
// when spawning via the API instead of spawn form. The structure of the options will depend on the Spawner's configuration.
// The body itself will be available as `user_options` for the Spawner.
func (c *APIClient) StartUserServer(ctx context.Context, name string, spawnOpts interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := c.cfg.BasePath + "/users/{name}/server"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

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
	localVarPostBody = spawnOpts
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

// StopUserServer Stop a user's server
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param name username
func (c *APIClient) StopUserServer(ctx context.Context, name string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := c.cfg.BasePath + "/users/{name}/server"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

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
		return localVarHttpResponse, GenericError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
	}

	return localVarHttpResponse, nil
}

// StartUserNamedServer Start a user's single-user named-server notebook server
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param name username
// @param serverName name given to a named-server.  Note that depending on your JupyterHub infrastructure there are chracterter size limitation to &#x60;server_name&#x60;. Default spawner with K8s pod will not allow Jupyter Notebooks to be spawned with a name that contains more than 253 characters (keep in mind that the pod will be spawned with extra characters to identify the user and hub).
// @param spawnOpts - Optional Spawn options can be passed as a JSON body
//
// when spawning via the API instead of spawn form. The structure of the options will depend on the Spawner's configuration.
func (c *APIClient) StartUserNamedServer(ctx context.Context, name string, serverName string, spawnOpts interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := c.cfg.BasePath + "/users/{name}/servers/{server_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"server_name"+"}", fmt.Sprintf("%v", serverName), -1)

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
	localVarPostBody = spawnOpts
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

// StopUserNamedServer Stop a user's named server
// To remove the named server in addition to deleting it, the body may be a JSON dictionary with a boolean &#x60;remove&#x60; field:  &#x60;&#x60;&#x60;json {\&quot;remove\&quot;: true} &#x60;&#x60;&#x60;
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param name username
// @param serverName name given to a named-server
func (c *APIClient) StopUserNamedServer(ctx context.Context, name string, serverName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := c.cfg.BasePath + "/users/{name}/servers/{server_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"server_name"+"}", fmt.Sprintf("%v", serverName), -1)

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
		return localVarHttpResponse, GenericError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
	}

	return localVarHttpResponse, nil
}
