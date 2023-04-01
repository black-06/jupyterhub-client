# Go API client for JupyterHub

The REST API for JupyterHub

## Overview

- API version: 4.0.0b2
- Package version: 1.0.0

## Installation

```shell
go get -u "github.com/black-06/jupyterhub-client"
```

## Usage

In jupyterhub_config.py

```python
c.JupyterHub.services = [
    {
        "name": "my-service",
        "api_token": "tokentokentoekntokentokentoekn",
        "admin": True,
    }
]
```

In golang client

```go
ctx := context.WithValue(context.Background(), ContextAccessToken, "tokentokentoekntokentokentoekn")

apiClient := NewAPIClient(&Configuration{
    BasePath:      "/hub/api",
    Host:          "127.0.0.1:8000",
    Scheme:        "http",
    DefaultHeader: make(map[string]string),
})

users, _, err := apiClient.UsersGet(ctx, &DefaultApiUsersGetOpts{})
```

## Documentation for API Endpoints

All URIs are relative to */hub/api*

### authorizations.go

| Method                            | HTTP request                                                | Description                                  |
|-----------------------------------|-------------------------------------------------------------|----------------------------------------------|
| RequestNewAPIToken                | **Post** /authorizations/token                              | Request a new API token                      |
| IdentifyUserOrServiceFromAPIToken | **Get** /authorizations/token/{token}                       | Identify a user or service from an API token |
| IdentifyUserFromCookie            | **Get** /authorizations/cookie/{cookie_name}/{cookie_value} | Identify a user from a cookie                |

### groups.go

| Method               | HTTP request                      | Description                                        |
|----------------------|-----------------------------------|----------------------------------------------------|
| ListGroups           | **Get** /groups                   | List groups                                        |
| DeleteGroup          | **Delete** /groups/{name}         | Delete a group                                     |
| GetGroupByName       | **Get** /groups/{name}            | Get a group by name                                |
| CreateGroup          | **Post** /groups/{name}           | Create a group                                     |
| SetGroupProperties   | **Put** /groups/{name}/properties | Set the group properties. Added in JupyterHub 3.2. |
| RemoveUsersFromGroup | **Delete** /groups/{name}/users   | Remove users from a group                          |
| AddUsersToGroup      | **Post** /groups/{name}/users     | Add users to a group                               |

### oauth2.go

| Method             | HTTP request              | Description                  |
|--------------------|---------------------------|------------------------------|
| OAuth2Authorize    | **Get** /oauth2/authorize | OAuth 2.0 authorize endpoint |
| RequestOauth2Token | **Post** /oauth2/token    | Request an OAuth2 token      |

### platform.go

| Method          | HTTP request       | Description                        |
|-----------------|--------------------|------------------------------------|
| GetVersion      | **Get** /          | Get JupyterHub version             |
| GetDetailedInfo | **Get** /info      | Get detailed info about JupyterHub |
| Shutdown        | **Post** /shutdown | Shutdown the Hub                   |

### proxy.go

| Method      | HTTP request     | Description                          |
|-------------|------------------|--------------------------------------|
| GetProxy    | **Get** /proxy   | Get the proxy's routing table        |
| SyncProxy   | **Post** /proxy  | Force the Hub to sync with the proxy |
| NotifyProxy | **Patch** /proxy | Notify the Hub about a new proxy     |

### services.go

| Method           | HTTP request             | Description           |
|------------------|--------------------------|-----------------------|
| ListServices     | **Get** /services        | List services         |
| GetServiceByName | **Get** /services/{name} | Get a service by name |

### user_server.go

| Method               | HTTP request                                   | Description                                             |
|----------------------|------------------------------------------------|---------------------------------------------------------|
| StartUserServer      | **Post** /users/{name}/server                  | Start a user's single-user notebook server              |
| StopUserServer       | **Delete** /users/{name}/server                | Stop a user's server                                    |
| StartUserNamedServer | **Post** /users/{name}/servers/{server_name}   | Start a user's single-user named-server notebook server |
| StopUserNamedServer  | **Delete** /users/{name}/servers/{server_name} | Stop a user's named server                              |

### user_token.go

| Method           | HTTP request                                  | Description                     |
|------------------|-----------------------------------------------|---------------------------------|
| ListUserTokens   | **Get** /users/{name}/tokens                  | List tokens for the user        |
| CreateUserToken  | **Post** /users/{name}/tokens                 | Create a new token for the user |
| DeleteUserToken  | **Deltete** "/users/{name}/tokens/{token_id}" | Delete (revoke) a token by id   |
| GetUserTokenByID | **Get** /users/{name}/tokens/{token_id}       | Get the model for a token by id |

### users.go

| Method               | HTTP request                    | Description                              |
|----------------------|---------------------------------|------------------------------------------|
| GetAuthenticatedUser | **Get** /user                   | Return authenticated user's model        |
| ListUsers            | **Get** /users                  | List users                               |
| CreateUsers          | **Post** /users                 | Create multiple users                    |
| GetUserByName        | **Get** /users/{name}           | Get a user by name                       |
| CreateUser           | **Post** /users/{name}          | Create a single user                     |
| DeleteUser           | **Delete** /users/{name}        | Delete a user                            |
| UpdateUser           | **Patch** /users/{name}         | Modify a user                            |
| UpdateUserActivity   | **Post** /users/{name}/activity | Notify Hub of activity for a given user. |




