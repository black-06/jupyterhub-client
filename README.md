# Go API client for client

The REST API for JupyterHub

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 4.0.0b2
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "github.com/black-06/jupyterhub-client"
```

## Documentation for API Endpoints

All URIs are relative to */hub/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**AuthorizationsCookieCookieNameCookieValueGet**](docs/DefaultApi.md#authorizationscookiecookienamecookievalueget) | **Get** /authorizations/cookie/{cookie_name}/{cookie_value} | Identify a user from a cookie
*DefaultApi* | [**AuthorizationsTokenPost**](docs/DefaultApi.md#authorizationstokenpost) | **Post** /authorizations/token | Request a new API token
*DefaultApi* | [**AuthorizationsTokenTokenGet**](docs/DefaultApi.md#authorizationstokentokenget) | **Get** /authorizations/token/{token} | Identify a user or service from an API token
*DefaultApi* | [**GroupsGet**](docs/DefaultApi.md#groupsget) | **Get** /groups | List groups
*DefaultApi* | [**GroupsNameDelete**](docs/DefaultApi.md#groupsnamedelete) | **Delete** /groups/{name} | Delete a group
*DefaultApi* | [**GroupsNameGet**](docs/DefaultApi.md#groupsnameget) | **Get** /groups/{name} | Get a group by name
*DefaultApi* | [**GroupsNamePost**](docs/DefaultApi.md#groupsnamepost) | **Post** /groups/{name} | Create a group
*DefaultApi* | [**GroupsNamePropertiesPut**](docs/DefaultApi.md#groupsnamepropertiesput) | **Put** /groups/{name}/properties | Set the group properties.  Added in JupyterHub 3.2. 
*DefaultApi* | [**GroupsNameUsersDelete**](docs/DefaultApi.md#groupsnameusersdelete) | **Delete** /groups/{name}/users | Remove users from a group 
*DefaultApi* | [**GroupsNameUsersPost**](docs/DefaultApi.md#groupsnameuserspost) | **Post** /groups/{name}/users | Add users to a group
*DefaultApi* | [**InfoGet**](docs/DefaultApi.md#infoget) | **Get** /info | Get detailed info about JupyterHub
*DefaultApi* | [**Oauth2AuthorizeGet**](docs/DefaultApi.md#oauth2authorizeget) | **Get** /oauth2/authorize | OAuth 2.0 authorize endpoint
*DefaultApi* | [**Oauth2TokenPost**](docs/DefaultApi.md#oauth2tokenpost) | **Post** /oauth2/token | Request an OAuth2 token
*DefaultApi* | [**ProxyGet**](docs/DefaultApi.md#proxyget) | **Get** /proxy | Get the proxy&#x27;s routing table
*DefaultApi* | [**ProxyPatch**](docs/DefaultApi.md#proxypatch) | **Patch** /proxy | Notify the Hub about a new proxy
*DefaultApi* | [**ProxyPost**](docs/DefaultApi.md#proxypost) | **Post** /proxy | Force the Hub to sync with the proxy
*DefaultApi* | [**RootGet**](docs/DefaultApi.md#rootget) | **Get** / | Get JupyterHub version
*DefaultApi* | [**ServicesGet**](docs/DefaultApi.md#servicesget) | **Get** /services | List services
*DefaultApi* | [**ServicesNameGet**](docs/DefaultApi.md#servicesnameget) | **Get** /services/{name} | Get a service by name
*DefaultApi* | [**ShutdownPost**](docs/DefaultApi.md#shutdownpost) | **Post** /shutdown | Shutdown the Hub
*DefaultApi* | [**UserGet**](docs/DefaultApi.md#userget) | **Get** /user | Return authenticated user&#x27;s model
*DefaultApi* | [**UsersGet**](docs/DefaultApi.md#usersget) | **Get** /users | List users
*DefaultApi* | [**UsersNameActivityPost**](docs/DefaultApi.md#usersnameactivitypost) | **Post** /users/{name}/activity | Notify Hub of activity for a given user.
*DefaultApi* | [**UsersNameDelete**](docs/DefaultApi.md#usersnamedelete) | **Delete** /users/{name} | Delete a user
*DefaultApi* | [**UsersNameGet**](docs/DefaultApi.md#usersnameget) | **Get** /users/{name} | Get a user by name
*DefaultApi* | [**UsersNamePatch**](docs/DefaultApi.md#usersnamepatch) | **Patch** /users/{name} | Modify a user
*DefaultApi* | [**UsersNamePost**](docs/DefaultApi.md#usersnamepost) | **Post** /users/{name} | Create a single user
*DefaultApi* | [**UsersNameServerDelete**](docs/DefaultApi.md#usersnameserverdelete) | **Delete** /users/{name}/server | Stop a user&#x27;s server
*DefaultApi* | [**UsersNameServerPost**](docs/DefaultApi.md#usersnameserverpost) | **Post** /users/{name}/server | Start a user&#x27;s single-user notebook server
*DefaultApi* | [**UsersNameServersServerNameDelete**](docs/DefaultApi.md#usersnameserversservernamedelete) | **Delete** /users/{name}/servers/{server_name} | Stop a user&#x27;s named server
*DefaultApi* | [**UsersNameServersServerNamePost**](docs/DefaultApi.md#usersnameserversservernamepost) | **Post** /users/{name}/servers/{server_name} | Start a user&#x27;s single-user named-server notebook server
*DefaultApi* | [**UsersNameTokensGet**](docs/DefaultApi.md#usersnametokensget) | **Get** /users/{name}/tokens | List tokens for the user
*DefaultApi* | [**UsersNameTokensPost**](docs/DefaultApi.md#usersnametokenspost) | **Post** /users/{name}/tokens | Create a new token for the user
*DefaultApi* | [**UsersNameTokensTokenIdDelete**](docs/DefaultApi.md#usersnametokenstokeniddelete) | **Delete** /users/{name}/tokens/{token_id} | Delete (revoke) a token by id
*DefaultApi* | [**UsersNameTokensTokenIdGet**](docs/DefaultApi.md#usersnametokenstokenidget) | **Get** /users/{name}/tokens/{token_id} | Get the model for a token by id
*DefaultApi* | [**UsersPost**](docs/DefaultApi.md#userspost) | **Post** /users | Create multiple users

## Documentation For Models

 - [AuthorizationsTokenBody](docs/AuthorizationsTokenBody.md)
 - [Group](docs/Group.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse2001Authenticator](docs/InlineResponse2001Authenticator.md)
 - [InlineResponse2001Spawner](docs/InlineResponse2001Spawner.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [NameActivityBody](docs/NameActivityBody.md)
 - [NameTokensBody](docs/NameTokensBody.md)
 - [NameUsersBody](docs/NameUsersBody.md)
 - [Oauth2TokenBody](docs/Oauth2TokenBody.md)
 - [ProxyBody](docs/ProxyBody.md)
 - [RequestIdentity](docs/RequestIdentity.md)
 - [Server](docs/Server.md)
 - [Service](docs/Service.md)
 - [ShutdownBody](docs/ShutdownBody.md)
 - [Token](docs/Token.md)
 - [User](docs/User.md)
 - [UsersBody](docs/UsersBody.md)
 - [UsersNameBody](docs/UsersNameBody.md)
 - [UsersnameactivityServers](docs/UsersnameactivityServers.md)
 - [UsersnameactivityServersServerName](docs/UsersnameactivityServersServerName.md)

## Documentation For Authorization

## oauth2
- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://hub.example/hub/api/oauth2/authorize
- **Scopes**: 
 - **(no_scope)**: Identify the owner of the requesting entity.
 - **self**: The user’s own resources _(metascope for users, resolves to (no_scope) for services)_
 - **inherit**: Everything that the token-owning entity can access _(metascope for tokens)_
 - **admin-ui**: Access the admin page. Permission to take actions via the admin page granted separately.
 - **admin:users**: Read, write, create and delete users and their authentication state, not including their servers or tokens.
 - **admin:auth_state**: Read a user’s authentication state.
 - **users**: Read and write permissions to user models (excluding servers, tokens and authentication state).
 - **delete:users**: Delete users.
 - **list:users**: List users, including at least their names.
 - **read:users**: Read user models (excluding including servers, tokens and authentication state).
 - **read:users:name**: Read names of users.
 - **read:users:groups**: Read users’ group membership.
 - **read:users:activity**: Read time of last user activity.
 - **read:roles**: Read role assignments.
 - **read:roles:users**: Read user role assignments.
 - **read:roles:services**: Read service role assignments.
 - **read:roles:groups**: Read group role assignments.
 - **users:activity**: Update time of last user activity.
 - **admin:servers**: Read, start, stop, create and delete user servers and their state.
 - **admin:server_state**: Read and write users’ server state.
 - **servers**: Start and stop user servers.
 - **read:servers**: Read users’ names and their server models (excluding the server state).
 - **delete:servers**: Stop and delete users&#x27; servers.
 - **tokens**: Read, write, create and delete user tokens.
 - **read:tokens**: Read user tokens.
 - **admin:groups**: Read and write group information, create and delete groups.
 - **groups**: Read and write group information, including adding/removing users to/from groups.
 - **list:groups**: List groups, including at least their names.
 - **read:groups**: Read group models.
 - **read:groups:name**: Read group names.
 - **delete:groups**: Delete groups.
 - **list:services**: List services, including at least their names.
 - **read:services**: Read service models.
 - **read:services:name**: Read service names.
 - **read:hub**: Read detailed information about the Hub.
 - **access:servers**: Access user servers via API or browser.
 - **access:services**: Access services via API or browser.
 - **proxy**: Read information about the proxy’s routing table, sync the Hub with the proxy and notify the Hub about a new proxy.
 - **shutdown**: Shutdown the hub.
 - **read:metrics**: Read prometheus metrics.

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.
```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```
## token
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

