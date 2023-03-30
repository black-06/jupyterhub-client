# User

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The user&#x27;s name | [optional] [default to null]
**Admin** | **bool** | Whether the user is an admin | [optional] [default to null]
**Roles** | **[]string** | The names of roles this user has | [optional] [default to null]
**Groups** | **[]string** | The names of groups where this user is a member | [optional] [default to null]
**Server** | **string** | The user&#x27;s notebook server&#x27;s base URL, if running; null if not. | [optional] [default to null]
**Pending** | **string** | The currently pending action, if any | [optional] [default to null]
**LastActivity** | [**time.Time**](time.Time.md) | Timestamp of last-seen activity from the user | [optional] [default to null]
**Servers** | [**[]Server**](Server.md) | The servers for this user. By default: only includes _active_ servers. Changed in 3.0: if &#x60;?include_stopped_servers&#x60; parameter is specified, stopped servers will be included as well.  | [optional] [default to null]
**AuthState** | [***interface{}**](interface{}.md) | Authentication state of the user. Only available with admin:users:auth_state scope. None otherwise.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

