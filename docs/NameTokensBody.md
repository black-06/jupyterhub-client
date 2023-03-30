# NameTokensBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | **float64** | lifetime (in seconds) after which the requested token will expire. | [optional] [default to null]
**Note** | **string** | A note attached to the token for future bookkeeping | [optional] [default to null]
**Roles** | **[]string** | A list of role names from which to derive scopes. This is a shortcut for assigning collections of scopes; Tokens do not retain role assignment. (Changed in 3.0: roles are immediately resolved to scopes instead of stored on roles.)  | [optional] [default to null]
**Scopes** | **[]string** | A list of scopes that the token should have. (new in JupyterHub 3.0).  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

