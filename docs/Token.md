# Token

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | The token itself. Only present in responses to requests for a new token. | [optional] [default to null]
**Id** | **string** | The id of the API token. Used for modifying or deleting the token. | [optional] [default to null]
**User** | **string** | The user that owns a token (undefined if owned by a service) | [optional] [default to null]
**Service** | **string** | The service that owns the token (undefined of owned by a user) | [optional] [default to null]
**Roles** | **[]string** | Deprecated in JupyterHub 3, always an empty list. Tokens have &#x27;scopes&#x27; starting from JupyterHub 3. | [optional] [default to null]
**Scopes** | **[]string** | List of scopes this token has been assigned. New in JupyterHub 3. In JupyterHub 2.x, tokens were assigned &#x27;roles&#x27; insead of scopes. | [optional] [default to null]
**Note** | **string** | A note about the token, typically describing what it was created for. | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) | Timestamp when this token was created | [optional] [default to null]
**ExpiresAt** | [**time.Time**](time.Time.md) | Timestamp when this token expires. Null if there is no expiry. | [optional] [default to null]
**LastActivity** | [**time.Time**](time.Time.md) | Timestamp of last-seen activity using this token. Can be null if token has never been used.  | [optional] [default to null]
**SessionId** | **string** | The session id associated with the token, if any. Only used for tokens set during oauth flows.  Added in 2.0.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

