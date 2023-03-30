# {{classname}}

All URIs are relative to */hub/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizationsCookieCookieNameCookieValueGet**](DefaultApi.md#AuthorizationsCookieCookieNameCookieValueGet) | **Get** /authorizations/cookie/{cookie_name}/{cookie_value} | Identify a user from a cookie
[**AuthorizationsTokenPost**](DefaultApi.md#AuthorizationsTokenPost) | **Post** /authorizations/token | Request a new API token
[**AuthorizationsTokenTokenGet**](DefaultApi.md#AuthorizationsTokenTokenGet) | **Get** /authorizations/token/{token} | Identify a user or service from an API token
[**GroupsGet**](DefaultApi.md#GroupsGet) | **Get** /groups | List groups
[**GroupsNameDelete**](DefaultApi.md#GroupsNameDelete) | **Delete** /groups/{name} | Delete a group
[**GroupsNameGet**](DefaultApi.md#GroupsNameGet) | **Get** /groups/{name} | Get a group by name
[**GroupsNamePost**](DefaultApi.md#GroupsNamePost) | **Post** /groups/{name} | Create a group
[**GroupsNamePropertiesPut**](DefaultApi.md#GroupsNamePropertiesPut) | **Put** /groups/{name}/properties | Set the group properties.  Added in JupyterHub 3.2. 
[**GroupsNameUsersDelete**](DefaultApi.md#GroupsNameUsersDelete) | **Delete** /groups/{name}/users | Remove users from a group 
[**GroupsNameUsersPost**](DefaultApi.md#GroupsNameUsersPost) | **Post** /groups/{name}/users | Add users to a group
[**InfoGet**](DefaultApi.md#InfoGet) | **Get** /info | Get detailed info about JupyterHub
[**Oauth2AuthorizeGet**](DefaultApi.md#Oauth2AuthorizeGet) | **Get** /oauth2/authorize | OAuth 2.0 authorize endpoint
[**Oauth2TokenPost**](DefaultApi.md#Oauth2TokenPost) | **Post** /oauth2/token | Request an OAuth2 token
[**ProxyGet**](DefaultApi.md#ProxyGet) | **Get** /proxy | Get the proxy&#x27;s routing table
[**ProxyPatch**](DefaultApi.md#ProxyPatch) | **Patch** /proxy | Notify the Hub about a new proxy
[**ProxyPost**](DefaultApi.md#ProxyPost) | **Post** /proxy | Force the Hub to sync with the proxy
[**RootGet**](DefaultApi.md#RootGet) | **Get** / | Get JupyterHub version
[**ServicesGet**](DefaultApi.md#ServicesGet) | **Get** /services | List services
[**ServicesNameGet**](DefaultApi.md#ServicesNameGet) | **Get** /services/{name} | Get a service by name
[**ShutdownPost**](DefaultApi.md#ShutdownPost) | **Post** /shutdown | Shutdown the Hub
[**UserGet**](DefaultApi.md#UserGet) | **Get** /user | Return authenticated user&#x27;s model
[**UsersGet**](DefaultApi.md#UsersGet) | **Get** /users | List users
[**UsersNameActivityPost**](DefaultApi.md#UsersNameActivityPost) | **Post** /users/{name}/activity | Notify Hub of activity for a given user.
[**UsersNameDelete**](DefaultApi.md#UsersNameDelete) | **Delete** /users/{name} | Delete a user
[**UsersNameGet**](DefaultApi.md#UsersNameGet) | **Get** /users/{name} | Get a user by name
[**UsersNamePatch**](DefaultApi.md#UsersNamePatch) | **Patch** /users/{name} | Modify a user
[**UsersNamePost**](DefaultApi.md#UsersNamePost) | **Post** /users/{name} | Create a single user
[**UsersNameServerDelete**](DefaultApi.md#UsersNameServerDelete) | **Delete** /users/{name}/server | Stop a user&#x27;s server
[**UsersNameServerPost**](DefaultApi.md#UsersNameServerPost) | **Post** /users/{name}/server | Start a user&#x27;s single-user notebook server
[**UsersNameServersServerNameDelete**](DefaultApi.md#UsersNameServersServerNameDelete) | **Delete** /users/{name}/servers/{server_name} | Stop a user&#x27;s named server
[**UsersNameServersServerNamePost**](DefaultApi.md#UsersNameServersServerNamePost) | **Post** /users/{name}/servers/{server_name} | Start a user&#x27;s single-user named-server notebook server
[**UsersNameTokensGet**](DefaultApi.md#UsersNameTokensGet) | **Get** /users/{name}/tokens | List tokens for the user
[**UsersNameTokensPost**](DefaultApi.md#UsersNameTokensPost) | **Post** /users/{name}/tokens | Create a new token for the user
[**UsersNameTokensTokenIdDelete**](DefaultApi.md#UsersNameTokensTokenIdDelete) | **Delete** /users/{name}/tokens/{token_id} | Delete (revoke) a token by id
[**UsersNameTokensTokenIdGet**](DefaultApi.md#UsersNameTokensTokenIdGet) | **Get** /users/{name}/tokens/{token_id} | Get the model for a token by id
[**UsersPost**](DefaultApi.md#UsersPost) | **Post** /users | Create multiple users

# **AuthorizationsCookieCookieNameCookieValueGet**
> User AuthorizationsCookieCookieNameCookieValueGet(ctx, cookieName, cookieValue)
Identify a user from a cookie

Used by single-user notebook servers to hand off cookie authentication to the Hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cookieName** | **string**|  | 
  **cookieValue** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

[oauth2](../README.md#oauth2), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationsTokenPost**
> InlineResponse2002 AuthorizationsTokenPost(ctx, optional)
Request a new API token

Request a new API token to use with the JupyterHub REST API. If not already authenticated, username and password can be sent in the JSON request body. Logging in via this method is only available when the active Authenticator accepts passwords (e.g. not OAuth). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiAuthorizationsTokenPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiAuthorizationsTokenPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AuthorizationsTokenBody**](AuthorizationsTokenBody.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationsTokenTokenGet**
> AuthorizationsTokenTokenGet(ctx, token)
Identify a user or service from an API token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGet**
> []Group GroupsGet(ctx, optional)
List groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiGroupsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGroupsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Float64**| Return a number of groups starting at the specified offset. Can be used with limit to paginate. If unspecified, return all groups.  | 
 **limit** | **optional.Float64**| Return a finite number of groups. Can be used with offset to paginate. If unspecified, use api_page_default_limit.  | 

### Return type

[**[]Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameDelete**
> GroupsNameDelete(ctx, name)
Delete a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| group name | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameGet**
> Group GroupsNameGet(ctx, name)
Get a group by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| group name | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNamePost**
> Group GroupsNamePost(ctx, name)
Create a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| group name | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNamePropertiesPut**
> Group GroupsNamePropertiesPut(ctx, body, name)
Set the group properties.  Added in JupyterHub 3.2. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)| The new group properties, as a JSON dict. | 
  **name** | **string**| group name | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameUsersDelete**
> GroupsNameUsersDelete(ctx, name)
Remove users from a group 

Body should be a JSON dictionary where `users` is a list of usernames to remove from the groups.  ```json {   \"users\": [\"name1\", \"name2\"] } ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| group name | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameUsersPost**
> Group GroupsNameUsersPost(ctx, body, name)
Add users to a group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NameUsersBody**](NameUsersBody.md)| The users to add to the group | 
  **name** | **string**| group name | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InfoGet**
> InlineResponse2001 InfoGet(ctx, )
Get detailed info about JupyterHub

Detailed JupyterHub information, including Python version, JupyterHub's version and executable path, and which Authenticator and Spawner are active. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Oauth2AuthorizeGet**
> Oauth2AuthorizeGet(ctx, clientId, responseType, redirectUri, optional)
OAuth 2.0 authorize endpoint

Redirect users to this URL to begin the OAuth process. It is not an API endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| The client id | 
  **responseType** | **string**| The response type (always &#x27;code&#x27;) | 
  **redirectUri** | **string**| The redirect url | 
 **optional** | ***DefaultApiOauth2AuthorizeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiOauth2AuthorizeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **state** | **optional.String**| A state string | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Oauth2TokenPost**
> InlineResponse2003 Oauth2TokenPost(ctx, clientId, clientSecret, grantType, code, redirectUri)
Request an OAuth2 token

Request an OAuth2 token from an authorization code. This request completes the OAuth process. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**|  | 
  **clientSecret** | **string**|  | 
  **grantType** | **string**|  | 
  **code** | **string**|  | 
  **redirectUri** | **string**|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[oauth2](../README.md#oauth2), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProxyGet**
> interface{} ProxyGet(ctx, optional)
Get the proxy's routing table

A convenience alias for getting the routing table directly from the proxy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiProxyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiProxyGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Float64**| Return a number of routes starting at the given offset. Can be used with limit to paginate. If unspecified, return all routes.  | 
 **limit** | **optional.Float64**| Return a finite number of routes. Can be used with offset to paginate. If unspecified, use api_page_default_limit  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProxyPatch**
> ProxyPatch(ctx, body)
Notify the Hub about a new proxy

Notifies the Hub of a new proxy to use.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProxyBody**](ProxyBody.md)| Any values that have changed for the new proxy. All keys are optional. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProxyPost**
> ProxyPost(ctx, )
Force the Hub to sync with the proxy

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RootGet**
> InlineResponse200 RootGet(ctx, )
Get JupyterHub version

This endpoint is not authenticated for the purpose of clients and user to identify the JupyterHub version before setting up authentication. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[oauth2](../README.md#oauth2), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesGet**
> []Service ServicesGet(ctx, )
List services

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Service**](Service.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesNameGet**
> Service ServicesNameGet(ctx, name)
Get a service by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| service name | 

### Return type

[**Service**](Service.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShutdownPost**
> ShutdownPost(ctx, optional)
Shutdown the Hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiShutdownPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiShutdownPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ShutdownBody**](ShutdownBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGet**
> RequestIdentity UserGet(ctx, )
Return authenticated user's model

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RequestIdentity**](RequestIdentity.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGet**
> []User UsersGet(ctx, optional)
List users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiUsersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUsersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **optional.String**| Return only users who have servers in the given state. If unspecified, return all users.  active: all users with any active servers (ready OR pending) ready: all users who have any ready servers (running, not pending) inactive: all users who have *no* active servers (complement of active)  Added in JupyterHub 1.3  | 
 **offset** | **optional.Float64**| Return a number users starting at the given offset. Can be used with limit to paginate. If unspecified, return all users.  | 
 **limit** | **optional.Float64**| Return a finite number of users. Can be used with offset to paginate. If unspecified, use api_page_default_limit.  | 
 **includeStoppedServers** | **optional.Bool**| Include stopped servers in user model(s). Added in JupyterHub 3.0. Allows retrieval of information about stopped servers, such as activity and state fields.  | 

### Return type

[**[]User**](User.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameActivityPost**
> UsersNameActivityPost(ctx, name, optional)
Notify Hub of activity for a given user.

Notify the Hub of activity by the user, e.g. accessing a service or (more likely) actively using a server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| username | 
 **optional** | ***DefaultApiUsersNameActivityPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUsersNameActivityPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NameActivityBody**](NameActivityBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameDelete**
> UsersNameDelete(ctx, name)
Delete a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| username | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameGet**
> User UsersNameGet(ctx, name)
Get a user by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| username | 

### Return type

[**User**](User.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNamePatch**
> User UsersNamePatch(ctx, body, name)
Modify a user

Change a user's name or admin status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsersNameBody**](UsersNameBody.md)| Updated user info. At least one key to be updated (name or admin) is required. | 
  **name** | **string**| username | 

### Return type

[**User**](User.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNamePost**
> User UsersNamePost(ctx, name)
Create a single user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| username | 

### Return type

[**User**](User.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameServerDelete**
> UsersNameServerDelete(ctx, name)
Stop a user's server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| username | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameServerPost**
> UsersNameServerPost(ctx, name, optional)
Start a user's single-user notebook server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| username | 
 **optional** | ***DefaultApiUsersNameServerPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUsersNameServerPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)| Spawn options can be passed as a JSON body
when spawning via the API instead of spawn form.
The structure of the options
will depend on the Spawner&#x27;s configuration.
The body itself will be available as &#x60;user_options&#x60; for the
Spawner.
 | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameServersServerNameDelete**
> UsersNameServersServerNameDelete(ctx, name, serverName)
Stop a user's named server

To remove the named server in addition to deleting it, the body may be a JSON dictionary with a boolean `remove` field:  ```json {\"remove\": true} ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| username | 
  **serverName** | **string**| name given to a named-server | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameServersServerNamePost**
> UsersNameServersServerNamePost(ctx, name, serverName, optional)
Start a user's single-user named-server notebook server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| username | 
  **serverName** | **string**| name given to a named-server.  Note that depending on your JupyterHub infrastructure there are chracterter size limitation to &#x60;server_name&#x60;. Default spawner with K8s pod will not allow Jupyter Notebooks to be spawned with a name that contains more than 253 characters (keep in mind that the pod will be spawned with extra characters to identify the user and hub).  | 
 **optional** | ***DefaultApiUsersNameServersServerNamePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUsersNameServersServerNamePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of interface{}**](interface{}.md)| Spawn options can be passed as a JSON body
when spawning via the API instead of spawn form.
The structure of the options
will depend on the Spawner&#x27;s configuration.
 | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameTokensGet**
> []Token UsersNameTokensGet(ctx, name)
List tokens for the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| username | 

### Return type

[**[]Token**](Token.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameTokensPost**
> Token UsersNameTokensPost(ctx, name, optional)
Create a new token for the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| username | 
 **optional** | ***DefaultApiUsersNameTokensPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUsersNameTokensPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NameTokensBody**](NameTokensBody.md)|  | 

### Return type

[**Token**](Token.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameTokensTokenIdDelete**
> UsersNameTokensTokenIdDelete(ctx, name, tokenId)
Delete (revoke) a token by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| username | 
  **tokenId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameTokensTokenIdGet**
> Token UsersNameTokensTokenIdGet(ctx, name, tokenId)
Get the model for a token by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| username | 
  **tokenId** | **string**|  | 

### Return type

[**Token**](Token.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPost**
> []User UsersPost(ctx, body)
Create multiple users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsersBody**](UsersBody.md)|  | 

### Return type

[**[]User**](User.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

