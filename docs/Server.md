# Server

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The server&#x27;s name. The user&#x27;s default server has an empty name (&#x27;&#x27;) | [optional] [default to null]
**Ready** | **bool** | Whether the server is ready for traffic. Will always be false when any transition is pending.  | [optional] [default to null]
**Stopped** | **bool** | Whether the server is stopped. Added in JupyterHub 3.0, and only useful when using the &#x60;?include_stopped_servers&#x60; request parameter. Now that stopped servers may be included (since JupyterHub 3.0), this is the simplest way to select stopped servers. Always equivalent to &#x60;not (ready or pending)&#x60;.  | [optional] [default to null]
**Pending** | **string** | The currently pending action, if any. A server is not ready if an action is pending.  | [optional] [default to null]
**Url** | **string** | The URL where the server can be accessed (typically /user/:name/:server.name/).  | [optional] [default to null]
**ProgressUrl** | **string** | The URL for an event-stream to retrieve events during a spawn.  | [optional] [default to null]
**Started** | [**time.Time**](time.Time.md) | UTC timestamp when the server was last started. | [optional] [default to null]
**LastActivity** | [**time.Time**](time.Time.md) | UTC timestamp last-seen activity on this server. | [optional] [default to null]
**State** | [***interface{}**](interface{}.md) | Arbitrary internal state from this server&#x27;s spawner. Only available on the hub&#x27;s users list or get-user-by-name method, and only with admin:users:server_state scope. None otherwise. | [optional] [default to null]
**UserOptions** | [***interface{}**](interface{}.md) | User specified options for the user&#x27;s spawned instance of a single-user server. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

