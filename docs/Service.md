# Service

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The service&#x27;s name | [optional] [default to null]
**Admin** | **bool** | Whether the service is an admin | [optional] [default to null]
**Roles** | **[]string** | The names of roles this service has | [optional] [default to null]
**Url** | **string** | The internal url where the service is running | [optional] [default to null]
**Prefix** | **string** | The proxied URL prefix to the service&#x27;s url | [optional] [default to null]
**Pid** | **float64** | The PID of the service process (if managed) | [optional] [default to null]
**Command** | **[]string** | The command used to start the service (if managed) | [optional] [default to null]
**Info** | [***interface{}**](interface{}.md) | Additional information a deployment can attach to a service. JupyterHub does not use this field.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

