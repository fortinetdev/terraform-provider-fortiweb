---
subcategory: "FortiWEB Server Policy"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_server_policy_server_pool_pserver_list"
description: |-
  Configure FortiWEB server policy server pool pserver list configuration.
---

# fortiweb_server_policy_server_pool_pserver_list
Configure FortiWEB server policy server pool pserver list configuration.

## Example Usage
```hcl
resource "fortiweb_server_policy_server_pool_pserver_list" "test" {
	mkey = "test"
	sub_mkey = "1"
	status = "enable"
	server_type = "physical"
	ip = "1.1.1.1"
	port = 443
	conn_limit = 0
	proxy_protocol = "enable"
	proxy_protocol_version = "v1"
	http2 = "enable"
	ssl = "enable"
	client_certificate = "sign-p12"
	server_side_sni = "enable"
	server_certificate_verify = "enable"
	server_certificate_verify_policy = "test"
	server_certificate_verify_action = "alert_deny"
	use_ciphers_group = "disable"
	tls_v10 = "disable"
	tls_v11 = "disable"
	tls_v12 = "enable"
	tls_v13 = "enable"
	ssl_cipher = "medium"
	rfc7919_comply = "disable"
	recover = 123
	warm_up = 1234
	warm_rate = 10
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the server policy server pool.
* `sub_mkey` - The ID of the server policy server pool pserver list entry.
* `status` - Specify the status of the pool member. Valid values: enable, disable, maintain.
* `server_type` - Specify whether to specify the pool member by IP address, domain, or automatically pulled by SDN connector. Valid values: physical, domain, sdn-connector.
* `ip` - The IP address of the web server to include in the pool.
* `port` - The TCP port number where the pool member listens for connections. Valid values: 1–65535. 
* `conn_limit` - The maximum number of TCP connections that FortiWeb forwards to this pool member. Valid values: 0–1048576.
* `weight` - If the server pool uses the weighted round robin load-balancing algorithm, type the numerical weight of the pool member. Valid values: 1–99999.
* `health_check_inherit` - Use the health check specified by health in the server pool/in this pool member configuration. Valid values: enable, disable.
* `hlck_domain` - Enter the domain name of the server pool.
* `backup_server` - Enter enable to configure this pool member as a backup server. Valid values: enable, disable.
* `proxy_protocol` - If the back-end server enables proxy protocol, you need to enable the Proxy Protocol option on FortiWeb so that the TCP SSL and HTTP traffic can successfully go through. Valid values: enable, disable.
* `proxy_protocol_version` - the proxy protocol version for the back-end server. Valid values: v1, v2.
* `http2` - Enable to allow HTTP/2 communication between the FortiWeb and this back-end web server for HTTP/2 security inspections in Reverse Proxy mode; or enable HTTP/2 security inspections in True Transparent Proxy mode.
* `ssl` - For Reverse Proxy, Offline Protection, and Transparent Inspection modes, specifies whether connections between FortiWeb and the pool member use SSL/TLS. Valid values: enable, disable.
* `client_certificate_proxy` - Enable to configure seamless PKI integration. Valid values: enable, disable.
* `client_certificate_proxy_sign_ca` - Select a Sign CA FortiWeb will use to verify and resign new client certificates.
* `client_certificate` - The client certificate that FortiWeb uses to connect to this server pool member.
* `server_side_sni` - Specify whether FortiWeb supports SNI for back-end servers that it applies this policy to. Valid values: enable, disable.
* `server_certificate_verify` - Enable so that FortiWeb appliance will verify certificates presented by HTTP server. Valid values: enable, disable.
* `server_certificate_verify_policy` - Enter the certificate verity policy name.
* `server_certificate_verify_action` - Select which action the FortiWeb appliance will take when it detects a certificate violation. Valid values: alert, alert_deny, redirect.
* `tls_v10` - For Reverse Proxy mode, specifies whether secure connections between FortiWeb and the server pool member can use the TLS 1.0 cryptographic protocol. Valid values: enable, disable.
* `tls_v11` - For Reverse Proxy mode, specifies whether secure connections between FortiWeb and the server pool member can use the TLS 1.1 cryptographic protocol. Valid values: enable, disable.
* `tls_v12` - For Reverse Proxy mode, specifies whether secure connections between FortiWeb and the server pool member can use the TLS 1.2 cryptographic protocol. Valid values: enable, disable.
* `tls_v13` - For Reverse Proxy mode, specifies whether secure connections between FortiWeb and the server pool member can use the TLS 1.3 cryptographic protocol. Valid values: enable, disable.
* `ssl_cipher` - For Reverse Proxy mode, specifies whether secure connections between FortiWeb and the server pool member use a medium-security, high-security, or custom set of cipher suites. Valid values: medium, high, custom.
* `use_ciphers_group` - Enable to use SSL Ciphers Group. Valid values: enable, disable.
* `ssl_ciphers_group` - Specify the name of SSL Ciphers Group when use_ciphers_group is enable. Example: Mozilla-Modern, Mozilla-Intermediate.
* `rfc7919_comply` - Enable to apply cipher suites that comply with RFC-9719. Valid values: enable, disable.
* `supported_group` - Select the RFC-9719 ciphers to be supported. Valid values: X25519, prime256v1, secp384r1, secp521r1, brainpoolP256r1, brainpoolP384r1, brainpoolP512r1, ffdhe2048, ffdhe3072, ffdhe4096, ffdhe6144, ffdhe8192.
* `session_ticket_reuse` - Enable so that FortiWeb reuses the session ticket when establishing an SSL connection to a pserver. Valid values: enable, disable.
* `session_id_reuse` - Enable so that FortiWeb reuses the session ID when establishing an SSL connection to a pserver. Valid values: enable, disable.
* `recover` - Specify the number of seconds that FortiWeb waits before it forwards traffic to this pool member after a health check indicates that this server is available again. Valid values: 0–86400.
* `warm_up` - Specify for how long FortiWeb forwards traffic at a reduced rate after a health check indicates that this pool member is available again but it cannot yet handle a full connection load. Valid values: 0–86400.
* `warm_rate` - Specify the maximum connection rate while the pool member is starting up. Valid values: 1–86400.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Server Policy Server Pool Pserver List can be imported using any of these accepted formats:
```
$ terraform import fortiweb_server_policy_server_pool_pserver_list.labelname {mkey}
```
