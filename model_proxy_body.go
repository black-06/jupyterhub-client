/*
 * JupyterHub
 *
 * The REST API for JupyterHub
 *
 * API version: 4.0.0b2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type ProxyBody struct {
	// IP address of the new proxy
	Ip string `json:"ip,omitempty"`
	// Port of the new proxy
	Port string `json:"port,omitempty"`
	// Protocol of new proxy, if changed
	Protocol string `json:"protocol,omitempty"`
	// CONFIGPROXY_AUTH_TOKEN for the new proxy
	AuthToken string `json:"auth_token,omitempty"`
}
